package inventory

import (
	"wms.com/core/database"
)

type inventoryService struct {
}

func NewInventoryService() InventoryService {
	return &inventoryService{}
}

// InventoryService represents a service for managing inventorys.
type InventoryService interface {
	//Item Management
	GetItemByID(int64) (Item, error)
	GetItemList(ItemFilter) (int, []Item, error)
}

func (s *inventoryService) GetItemByID(id int64) (Item, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	item, err := repo.GetItemByID(id)
	return item, err
}

func (s *inventoryService) GetItemList(filter ItemFilter) (int, []Item, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	count, err := repo.GetItemCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetItemList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}
