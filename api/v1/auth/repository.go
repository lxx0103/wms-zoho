package auth

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type authRepository struct {
	conn *sqlx.DB
}

func NewAuthRepository(connection *sqlx.DB) AuthRepository {
	return &authRepository{
		conn: connection,
	}
}

type AuthRepository interface {
	GetCredential(SigninRequest) (UserAuth, error)
	CreateAuth(UserAuth) (int64, error)
	CheckConfict(int, string) (bool, error)
	UpdateUserID(UserAuth) error
	// GetAuthCount(filter AuthFilter) (int, error)
	// GetAuthList(filter AuthFilter) ([]Auth, error)
}

func (r *authRepository) GetCredential(signInfo SigninRequest) (UserAuth, error) {
	var authInfo UserAuth
	err := r.conn.Get(&authInfo, "SELECT user_id, credential FROM user_auths WHERE auth_type = ? AND identifier = ?", signInfo.AuthType, signInfo.Identifier)
	if err != nil {
		return UserAuth{}, err
	}
	return authInfo, nil
}

func (r *authRepository) CreateAuth(signupInfo UserAuth) (int64, error) {
	res, err := r.conn.Exec(`
		INSERT INTO user_auths
		(
			auth_type,
			identifier,
			credential,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, signupInfo.AuthType, signupInfo.Identifier, signupInfo.Credential, 1, time.Now(), "system", time.Now(), "system")
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *authRepository) UpdateUserID(authInfo UserAuth) error {
	_, err := r.conn.Exec(`
		UPDATE user_auths
		SET user_id = ?
		WHERE id = ?
	`, authInfo.UserID, authInfo.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *authRepository) CheckConfict(authType int, identifier string) (bool, error) {
	existed := 0
	err := r.conn.Get(&existed, "SELECT count(1) FROM user_auths WHERE auth_type = ? AND identifier = ?", authType, identifier)
	if err != nil {
		return true, err
	}
	return existed != 0, nil
}

// func (r *authRepository) GetAuthCount(filter AuthFilter) (int, error) {
// 	where, args := []string{"1 = 1"}, []interface{}{}
// 	if v := filter.Name; v != "" {
// 		where, args = append(where, "name = ?"), append(args, v)
// 	}
// 	if v := filter.Email; v != "" {
// 		where, args = append(where, "email = ?"), append(args, v)
// 	}
// 	var count int
// 	err := r.conn.Get(&count, `
// 		SELECT count(1) as count
// 		FROM auths
// 		WHERE `+strings.Join(where, " AND "), args...)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return count, nil
// }

// func (r *authRepository) GetAuthList(filter AuthFilter) ([]Auth, error) {
// 	where, args := []string{"1 = 1"}, []interface{}{}
// 	if v := filter.Name; v != "" {
// 		where, args = append(where, "name = ?"), append(args, v)
// 	}
// 	if v := filter.Email; v != "" {
// 		where, args = append(where, "email = ?"), append(args, v)
// 	}
// 	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
// 	args = append(args, filter.PageSize)
// 	var auths []Auth
// 	err := r.conn.Select(&auths, `
// 		SELECT *
// 		FROM auths
// 		WHERE `+strings.Join(where, " AND ")+`
// 		LIMIT ?, ?
// 	`, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return auths, nil
// }
