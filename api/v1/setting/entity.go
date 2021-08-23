package setting

import "time"

type Shelf struct {
	ID        int64     `db:"id" json:"id"`
	Code      string    `db:"code" json:"code"`
	Level     int       `db:"level" json:"level"`
	Location  string    `db:"location" json:"location"`
	Enabled   string    `db:"enabled" json:"enabled"`
	Created   time.Time `db:"created" json:"created"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	Updated   time.Time `db:"updated" json:"updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}
