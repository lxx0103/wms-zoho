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
	GetUserByID(id int64) (User, error)
	CreateUser(info User) (int64, error)
	GetUserCount(filter UserFilter) (int, error)
	GetUserList(filter UserFilter) ([]User, error)
}

func (r *userRepository) GetUserByID(id int64) (User, error) {
	var user User
	err := r.conn.Get(&user, "SELECT * FROM users WHERE id = ? ", id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (r *userRepository) CreateUser(info User) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO users
		(
			gender,
			name,
			email,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?)
	`, info.Gender, info.Name, info.Email, time.Now(), time.Now())
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
		FROM users 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *userRepository) GetUserList(filter UserFilter) ([]User, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Name; v != "" {
		where, args = append(where, "name = ?"), append(args, v)
	}
	if v := filter.Email; v != "" {
		where, args = append(where, "email = ?"), append(args, v)
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var users []User
	err := r.conn.Select(&users, `
		SELECT * 
		FROM users 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}
