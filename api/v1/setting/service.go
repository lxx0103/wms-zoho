package setting

import (
	"wms.com/core/database"
)

type settingService struct {
}

func NewSettingService() SettingService {
	return &settingService{}
}

// SettingService represents a service for managing settings.
type SettingService interface {
	GetShelfByID(int64) (Shelf, error)
	NewShelf(ShelfNew) (Shelf, error)
	GetShelfList(ShelfFilter) (int, []Shelf, error)
	UpdateShelf(int64, ShelfNew) (Shelf, error)
}

func (s *settingService) GetShelfByID(id int64) (Shelf, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	shelf, err := repo.GetShelfByID(id)
	return shelf, err
}

func (s *settingService) NewShelf(info ShelfNew) (Shelf, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	shelfID, err := repo.CreateShelf(info)
	if err != nil {
		return Shelf{}, err
	}
	shelf, err := repo.GetShelfByID(shelfID)
	return shelf, err
}

func (s *settingService) GetShelfList(filter ShelfFilter) (int, []Shelf, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	count, err := repo.GetShelfCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetShelfList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *settingService) UpdateShelf(shelfID int64, info ShelfNew) (Shelf, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	_, err := repo.UpdateShelf(shelfID, info)
	if err != nil {
		return Shelf{}, err
	}
	shelf, err := repo.GetShelfByID(shelfID)
	return shelf, err
}
