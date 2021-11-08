package auth

import (
	"fmt"
	"strings"
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

	//Role Management
	GetRoleByID(id int64) (UserRole, error)
	CreateRole(info RoleNew) (int64, error)
	GetRoleCount(filter RoleFilter) (int, error)
	GetRoleList(filter RoleFilter) ([]UserRole, error)
	UpdateRole(id int64, info RoleNew) (int64, error)
	//API Management
	GetAPIByID(id int64) (UserAPI, error)
	CreateAPI(info APINew) (int64, error)
	GetAPICount(filter APIFilter) (int, error)
	GetAPIList(filter APIFilter) ([]UserAPI, error)
	UpdateAPI(id int64, info APINew) (int64, error)
	//Menu Management
	GetMenuByID(id int64) (UserMenu, error)
	CreateMenu(info MenuNew) (int64, error)
	GetMenuCount(filter MenuFilter) (int, error)
	GetMenuList(filter MenuFilter) ([]UserMenu, error)
	UpdateMenu(id int64, info MenuNew) (int64, error)
	//Privilege Management
	GetRoleMenuByID(int64) ([]int64, error)
	NewRoleMenu(int64, RoleMenuNew) (int64, error)
	GetMenuAPIByID(int64) ([]int64, error)
	NewMenuAPI(int64, MenuAPINew) (int64, error)
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

func (r *authRepository) GetRoleByID(id int64) (UserRole, error) {
	var role UserRole
	err := r.conn.Get(&role, "SELECT * FROM user_roles WHERE id = ? ", id)
	if err != nil {
		return UserRole{}, err
	}
	return role, nil
}
func (r *authRepository) CreateRole(info RoleNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO user_roles
		(
			name,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`, info.Name, info.Enabled, time.Now(), info.User, time.Now(), info.User)
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

func (r *authRepository) GetRoleCount(filter RoleFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Name; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM user_roles 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *authRepository) GetRoleList(filter RoleFilter) ([]UserRole, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Name; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var roles []UserRole
	err := r.conn.Select(&roles, `
		SELECT * 
		FROM user_roles 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *authRepository) UpdateRole(id int64, info RoleNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		Update user_roles SET 
		name = ?,
		enabled = ?,
		updated = ?,
		updated_by = ? 
		WHERE id = ?
	`, info.Name, info.Enabled, time.Now(), info.User, id)
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

func (r *authRepository) GetAPIByID(id int64) (UserAPI, error) {
	var api UserAPI
	err := r.conn.Get(&api, "SELECT * FROM user_apis WHERE id = ? ", id)
	if err != nil {
		return UserAPI{}, err
	}
	return api, nil
}
func (r *authRepository) CreateAPI(info APINew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO user_apis
		(
			name,
			route,
			method,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, info.Name, info.Route, info.Method, info.Enabled, time.Now(), info.User, time.Now(), info.User)
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

func (r *authRepository) GetAPICount(filter APIFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Name; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	if v := filter.Route; v != "" {
		where, args = append(where, "route like ?"), append(args, "%"+v+"%")
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM user_apis 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *authRepository) GetAPIList(filter APIFilter) ([]UserAPI, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Name; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	if v := filter.Route; v != "" {
		where, args = append(where, "route like ?"), append(args, "%"+v+"%")
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var apis []UserAPI
	err := r.conn.Select(&apis, `
		SELECT * 
		FROM user_apis 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return apis, nil
}

func (r *authRepository) UpdateAPI(id int64, info APINew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		Update user_apis SET 
		name = ?,
		route = ?,
		method = ?,
		enabled = ?,
		updated = ?,
		updated_by = ? 
		WHERE id = ?
	`, info.Name, info.Route, info.Method, info.Enabled, time.Now(), info.User, id)
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

func (r *authRepository) GetMenuByID(id int64) (UserMenu, error) {
	var menu UserMenu
	err := r.conn.Get(&menu, "SELECT * FROM user_menus WHERE id = ? ", id)
	if err != nil {
		return UserMenu{}, err
	}
	return menu, nil
}
func (r *authRepository) CreateMenu(info MenuNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO user_menus
		(
			name,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`, info.Name, info.Enabled, time.Now(), info.User, time.Now(), info.User)
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

func (r *authRepository) GetMenuCount(filter MenuFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Name; v != "" {
		where, args = append(where, "code like ?"), append(args, "%"+v+"%")
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM user_menus 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *authRepository) GetMenuList(filter MenuFilter) ([]UserMenu, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Name; v != "" {
		where, args = append(where, "code like ?"), append(args, "%"+v+"%")
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var menus []UserMenu
	err := r.conn.Select(&menus, `
		SELECT * 
		FROM user_menus 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *authRepository) UpdateMenu(id int64, info MenuNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		Update user_menus SET 
		name = ?,
		enabled = ?,
		updated = ?,
		updated_by = ? 
		WHERE id = ?
	`, info.Name, info.Enabled, time.Now(), info.User, id)
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

func (r *authRepository) GetRoleMenuByID(id int64) ([]int64, error) {
	var menu []int64
	err := r.conn.Select(&menu, "SELECT menu_id FROM user_role_menus WHERE role_id = ? and enabled = 1", id)
	if err != nil {
		return nil, err
	}
	return menu, nil
}
func (r *authRepository) NewRoleMenu(role_id int64, info RoleMenuNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	_, err = tx.Exec(`
		Update user_role_menus SET 
		enabled = 2,
		updated = ?,
		updated_by = ? 
		WHERE role_id = ?
		AND enabled = 1
	`, time.Now(), info.User, role_id)
	if err != nil {
		return 0, err
	}
	sql := `
	INSERT INTO user_role_menus
	(
		role_id,
		menu_id,
		enabled,
		created,
		created_by,
		updated,
		updated_by
	)
	VALUES
	`
	for i := 0; i < len(info.IDS); i++ {
		sql += "(" + fmt.Sprint(role_id) + "," + fmt.Sprint(info.IDS[i]) + ",1,\"" + time.Now().Format("2006-01-02 15:01:01") + "\",\"" + info.User + "\",\"" + time.Now().Format("2006-01-02 15:01:01") + "\",\"" + info.User + "\"),"
	}
	sql = sql[:len(sql)-1]
	result, err := tx.Exec(sql)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return rows, nil
}

func (r *authRepository) GetMenuAPIByID(id int64) ([]int64, error) {
	var apis []int64
	err := r.conn.Select(&apis, "SELECT api_id FROM user_menu_apis WHERE menu_id = ? and enabled = 1", id)
	if err != nil {
		return nil, err
	}
	return apis, nil
}
func (r *authRepository) NewMenuAPI(menu_id int64, info MenuAPINew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	_, err = tx.Exec(`
		Update user_menu_apis SET 
		enabled = 2,
		updated = ?,
		updated_by = ? 
		WHERE menu_id = ?
		AND enabled = 1
	`, time.Now(), info.User, menu_id)
	if err != nil {
		return 0, err
	}
	sql := `
	INSERT INTO user_menu_apis
	(
		menu_id,
		api_id,
		enabled,
		created,
		created_by,
		updated,
		updated_by
	)
	VALUES
	`
	for i := 0; i < len(info.IDS); i++ {
		sql += "(" + fmt.Sprint(menu_id) + "," + fmt.Sprint(info.IDS[i]) + ",1,\"" + time.Now().Format("2006-01-02 15:01:01") + "\",\"" + info.User + "\",\"" + time.Now().Format("2006-01-02 15:01:01") + "\",\"" + info.User + "\"),"
	}
	sql = sql[:len(sql)-1]
	result, err := tx.Exec(sql)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return rows, nil
}
