package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"wms.com/api/v1/user"
	"wms.com/core/database"
)

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

type AuthService interface {
	VerifyCredential(SigninRequest) (user.UserProfile, error)
	CreateAuth(SignupRequest) (int64, error)
	//Role Management
	GetRoleByID(int64) (UserRole, error)
	NewRole(RoleNew) (UserRole, error)
	GetRoleList(RoleFilter) (int, []UserRole, error)
	UpdateRole(int64, RoleNew) (UserRole, error)
	//API Management
	GetAPIByID(int64) (UserAPI, error)
	NewAPI(APINew) (UserAPI, error)
	GetAPIList(APIFilter) (int, []UserAPI, error)
	UpdateAPI(int64, APINew) (UserAPI, error)
	//Menu Management
	GetMenuByID(int64) (UserMenu, error)
	NewMenu(MenuNew) (UserMenu, error)
	GetMenuList(MenuFilter) (int, []UserMenu, error)
	UpdateMenu(int64, MenuNew) (UserMenu, error)
	//Privilege Management
	GetRoleMenuByID(int64) ([]int64, error)
	NewRoleMenu(int64, RoleMenuNew) ([]int64, error)
	GetMenuAPIByID(int64) ([]int64, error)
	NewMenuAPI(int64, MenuAPINew) ([]int64, error)
	GetMyMenu(int64) ([]UserMenu, error)
}

func (s *authService) VerifyCredential(signinInfo SigninRequest) (user.UserProfile, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	authInfo, err := repo.GetCredential(signinInfo)
	if err != nil {
		return user.UserProfile{}, err
	}
	if !checkPasswordHash(signinInfo.Credential, authInfo.Credential) {
		return user.UserProfile{}, errors.New("CREDENTIAL ERROR")
	}
	userRepo := user.NewUserRepository(db)
	userInfo, err := userRepo.GetUserByID(authInfo.UserID)
	return userInfo, err
}

func (s authService) CreateAuth(signupInfo SignupRequest) (int64, error) {
	hashed, err := hashPassword(signupInfo.Credential)
	if err != nil {
		return 0, err
	}
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	var newAuth UserAuth
	newAuth.Credential = hashed
	isConflict, err := repo.CheckConfict(signupInfo.AuthType, signupInfo.Identifier)
	if err != nil {
		return 0, err
	}
	if isConflict {
		return 0, errors.New("CONFLICT")
	}
	newAuth.Identifier = signupInfo.Identifier
	newAuth.AuthType = signupInfo.AuthType
	authID, err := repo.CreateAuth(newAuth)
	if err != nil {
		return 0, err
	}
	return authID, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *authService) GetRoleByID(id int64) (UserRole, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	role, err := repo.GetRoleByID(id)
	return role, err
}

func (s *authService) NewRole(info RoleNew) (UserRole, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	roleID, err := repo.CreateRole(info)
	if err != nil {
		return UserRole{}, err
	}
	role, err := repo.GetRoleByID(roleID)
	return role, err
}

func (s *authService) GetRoleList(filter RoleFilter) (int, []UserRole, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	count, err := repo.GetRoleCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetRoleList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *authService) UpdateRole(roleID int64, info RoleNew) (UserRole, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	_, err := repo.UpdateRole(roleID, info)
	if err != nil {
		return UserRole{}, err
	}
	role, err := repo.GetRoleByID(roleID)
	return role, err
}

func (s *authService) GetAPIByID(id int64) (UserAPI, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	api, err := repo.GetAPIByID(id)
	return api, err
}

func (s *authService) NewAPI(info APINew) (UserAPI, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	apiID, err := repo.CreateAPI(info)
	if err != nil {
		return UserAPI{}, err
	}
	api, err := repo.GetAPIByID(apiID)
	return api, err
}

func (s *authService) GetAPIList(filter APIFilter) (int, []UserAPI, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	count, err := repo.GetAPICount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetAPIList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *authService) UpdateAPI(apiID int64, info APINew) (UserAPI, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	_, err := repo.UpdateAPI(apiID, info)
	if err != nil {
		return UserAPI{}, err
	}
	api, err := repo.GetAPIByID(apiID)
	return api, err
}

func (s *authService) GetMenuByID(id int64) (UserMenu, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	menu, err := repo.GetMenuByID(id)
	return menu, err
}

func (s *authService) NewMenu(info MenuNew) (UserMenu, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	menuID, err := repo.CreateMenu(info)
	if err != nil {
		return UserMenu{}, err
	}
	menu, err := repo.GetMenuByID(menuID)
	return menu, err
}

func (s *authService) GetMenuList(filter MenuFilter) (int, []UserMenu, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	count, err := repo.GetMenuCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetMenuList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *authService) UpdateMenu(menuID int64, info MenuNew) (UserMenu, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	_, err := repo.UpdateMenu(menuID, info)
	if err != nil {
		return UserMenu{}, err
	}
	menu, err := repo.GetMenuByID(menuID)
	return menu, err
}

func (s *authService) GetRoleMenuByID(id int64) ([]int64, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	menu, err := repo.GetRoleMenuByID(id)
	return menu, err
}

func (s *authService) NewRoleMenu(id int64, info RoleMenuNew) ([]int64, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	_, err := repo.NewRoleMenu(id, info)
	if err != nil {
		return nil, err
	}
	menu, err := repo.GetRoleMenuByID(id)
	return menu, err
}

func (s *authService) GetMenuAPIByID(id int64) ([]int64, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	menu, err := repo.GetMenuAPIByID(id)
	return menu, err
}

func (s *authService) NewMenuAPI(id int64, info MenuAPINew) ([]int64, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	_, err := repo.NewMenuAPI(id, info)
	if err != nil {
		return nil, err
	}
	menu, err := repo.GetMenuAPIByID(id)
	return menu, err
}

func (s *authService) GetMyMenu(roleID int64) ([]UserMenu, error) {
	db := database.InitMySQL()
	repo := NewAuthRepository(db)
	menu, err := repo.GetMyMenu(roleID)
	return menu, err
}
