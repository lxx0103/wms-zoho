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
	GetPurchaseOrderByID(int64) (*PurchaseOrder, error)
	GetPurchaseOrderList(PurchaseOrderFilter) (int, []PurchaseOrder, error)
	FilterPOItem(FilterPOItem) (*[]PurchaseOrderItem, error)
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

func (s *inventoryService) GetPurchaseOrderByID(id int64) (*PurchaseOrder, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	purchaseOrder, err := repo.GetPurchaseOrderByID(id)
	return purchaseOrder, err
}

func (s *inventoryService) GetPurchaseOrderList(filter PurchaseOrderFilter) (int, []PurchaseOrder, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	count, err := repo.GetPurchaseOrderCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetPurchaseOrderList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *inventoryService) FilterPOItem(filter FilterPOItem) (*[]PurchaseOrderItem, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	items, err := repo.FilterPOItem(filter)
	return items, err
}
