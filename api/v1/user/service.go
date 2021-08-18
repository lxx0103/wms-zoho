package user

import (
	"vandacare.com/core/database"
)

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

// UserService represents a service for managing users.
type UserService interface {
	GetUserByID(int64) (User, error)
	CreateNewUser(User) (User, error)
	GetUserList(UserFilter) (int, []User, error)
}

func (s *userService) GetUserByID(id int64) (User, error) {
	db := database.InitMySQL()
	repo := NewUserRepository(db)
	user, err := repo.GetUserByID(id)
	return user, err
}

func (s *userService) CreateNewUser(info User) (User, error) {
	db := database.InitMySQL()
	repo := NewUserRepository(db)
	userID, err := repo.CreateUser(info)
	if err != nil {
		return User{}, err
	}
	user, err := repo.GetUserByID(userID)
	return user, err
}

func (s *userService) GetUserList(filter UserFilter) (int, []User, error) {
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
