package auth

import "time"

type UserAuth struct {
	ID         int64     `db:"id" json:"id"`
	UserID     int64     `db:"user_id" json:"user_id"`
	AuthType   int       `db:"auth_type" json:"auth_type"` //1.Email, 2.Singpass, ...
	Identifier string    `db:"identifier" json:"identifier"`
	Credential string    `db:"credential" json:"credential"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}
type UserRole struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Enabled   string    `db:"enabled" json:"enabled"`
	Created   time.Time `db:"created" json:"created"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	Updated   time.Time `db:"updated" json:"updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}
type UserAPI struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Method    string    `db:"method" json:"method"`
	Route     string    `db:"route" json:"route"`
	Enabled   string    `db:"enabled" json:"enabled"`
	Created   time.Time `db:"created" json:"created"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	Updated   time.Time `db:"updated" json:"updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}
type UserMenu struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Enabled   string    `db:"enabled" json:"enabled"`
	Created   time.Time `db:"created" json:"created"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	Updated   time.Time `db:"updated" json:"updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}

type UserMenuAPI struct {
	ID        int64     `db:"id" json:"id"`
	MenuID    int64     `db:"menu_id" json:"menu_id"`
	APIID     int64     `db:"api_id" json:"api_id"`
	Enabled   string    `db:"enabled" json:"enabled"`
	Created   time.Time `db:"created" json:"created"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	Updated   time.Time `db:"updated" json:"updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}

type UserRoleMenu struct {
	ID        int64     `db:"id" json:"id"`
	MenuID    int64     `db:"menu_id" json:"menu_id"`
	APIID     int64     `db:"api_id" json:"api_id"`
	Enabled   string    `db:"enabled" json:"enabled"`
	Created   time.Time `db:"created" json:"created"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	Updated   time.Time `db:"updated" json:"updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}
