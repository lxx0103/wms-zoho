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
