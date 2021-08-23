package setting

import (
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type settingRepository struct {
	conn *sqlx.DB
}

func NewSettingRepository(connection *sqlx.DB) SettingRepository {
	return &settingRepository{
		conn: connection,
	}
}

type SettingRepository interface {
	GetShelfByID(id int64) (Shelf, error)
	CreateShelf(info ShelfNew) (int64, error)
	GetShelfCount(filter ShelfFilter) (int, error)
	GetShelfList(filter ShelfFilter) ([]Shelf, error)
	UpdateShelf(shelfID int64, info ShelfNew) (int64, error)
}

func (r *settingRepository) GetShelfByID(id int64) (Shelf, error) {
	var setting Shelf
	err := r.conn.Get(&setting, "SELECT * FROM s_shelves WHERE id = ? ", id)
	if err != nil {
		return Shelf{}, err
	}
	return setting, nil
}
func (r *settingRepository) CreateShelf(info ShelfNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO s_shelves
		(
			code,
			level,
			location,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, info.Code, info.Level, info.Location, info.Enabled, time.Now(), info.User, time.Now(), info.User)
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

func (r *settingRepository) GetShelfCount(filter ShelfFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Code; v != "" {
		where, args = append(where, "code = ?"), append(args, v)
	}
	if v := filter.Location; v != "" {
		where, args = append(where, "location = ?"), append(args, v)
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM s_shelves 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *settingRepository) GetShelfList(filter ShelfFilter) ([]Shelf, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Code; v != "" {
		where, args = append(where, "code = ?"), append(args, v)
	}
	if v := filter.Location; v != "" {
		where, args = append(where, "location = ?"), append(args, v)
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var shelves []Shelf
	err := r.conn.Select(&shelves, `
		SELECT * 
		FROM s_shelves 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return shelves, nil
}

func (r *settingRepository) UpdateShelf(shelfID int64, info ShelfNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		Update s_shelves SET 
		code = ?,
		level = ?,
		location = ?,
		enabled = ?,
		updated = ?,
		updated_by = ? 
		WHERE id = ?
	`, info.Code, info.Level, info.Location, info.Enabled, time.Now(), info.User, shelfID)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return id, nil
}
