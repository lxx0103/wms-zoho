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
