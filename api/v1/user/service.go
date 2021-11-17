package user

import (
	"wms.com/core/database"
)

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

// UserService represents a service for managing users.
type UserService interface {
	GetUserByID(int64) (UserProfile, error)
	CreateNewUser(UserProfile) (UserProfile, error)
	GetUserList(UserFilter) (int, []UserProfile, error)
	UpdateUser(int64, UserUpdate, int64) (UserProfile, error)
}

func (s *userService) GetUserByID(id int64) (UserProfile, error) {
	db := database.InitMySQL()
	repo := NewUserRepository(db)
	user, err := repo.GetUserByID(id)
	return user, err
}

func (s *userService) CreateNewUser(info UserProfile) (UserProfile, error) {
	db := database.InitMySQL()
	repo := NewUserRepository(db)
	userID, err := repo.CreateUser(info)
	if err != nil {
		return UserProfile{}, err
	}
	user, err := repo.GetUserByID(userID)
	return user, err
}

func (s *userService) GetUserList(filter UserFilter) (int, []UserProfile, error) {
	db := database.InitMySQL()
	repo := NewUserRepository(db)
	count, err := repo.GetUserCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetUserList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *userService) UpdateUser(userID int64, info UserUpdate, roleID int64) (UserProfile, error) {
	db := database.InitMySQL()
	repo := NewUserRepository(db)
	_, err := repo.UpdateUser(userID, info, roleID)
	if err != nil {
		return UserProfile{}, err
	}
	user, err := repo.GetUserByID(userID)
	return user, err
}
