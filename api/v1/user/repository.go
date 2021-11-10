package user

import (
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(connection *sqlx.DB) UserRepository {
	return &userRepository{
		conn: connection,
	}
}

type UserRepository interface {
	GetUserByID(id int64) (UserProfile, error)
	CreateUser(info UserProfile) (int64, error)
	GetUserCount(filter UserFilter) (int, error)
	GetUserList(filter UserFilter) ([]UserProfile, error)
	UpdateUser(int64, UserUpdate) (int64, error)
}

func (r *userRepository) GetUserByID(id int64) (UserProfile, error) {
	var user UserProfile
	err := r.conn.Get(&user, "SELECT * FROM user_profiles WHERE id = ? ", id)
	if err != nil {
		return UserProfile{}, err
	}
	return user, nil
}
func (r *userRepository) CreateUser(info UserProfile) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO user_profiles
		(
			name,
			email,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, info.Name, info.Email, 1, time.Now(), "system", time.Now(), "system")
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return id, nil
}

func (r *userRepository) GetUserCount(filter UserFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Name; v != "" {
		where, args = append(where, "name = ?"), append(args, v)
	}
	if v := filter.Email; v != "" {
		where, args = append(where, "email = ?"), append(args, v)
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM user_profiles 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *userRepository) GetUserList(filter UserFilter) ([]UserProfile, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Name; v != "" {
		where, args = append(where, "name = ?"), append(args, v)
	}
	if v := filter.Email; v != "" {
		where, args = append(where, "email = ?"), append(args, v)
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var users []UserProfile
	err := r.conn.Select(&users, `
		SELECT * 
		FROM user_profiles 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateUser(id int64, info UserUpdate) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		Update user_profiles SET 
		name = ?,
		email = ?,
		role_id = ?,
		updated = ?,
		updated_by = ? 
		WHERE id = ?
	`, info.Name, info.Email, info.RoleID, time.Now(), info.User, id)
	if err != nil {
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return affected, nil
}
