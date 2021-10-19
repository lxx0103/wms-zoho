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
	GetItemBySKU(string) (Item, error)
	GetItemList(ItemFilter) (int, []Item, error)
	//PO Management
	GetPurchaseOrderByID(int64) (*PurchaseOrder, error)
	GetPurchaseOrderList(PurchaseOrderFilter) (int, []PurchaseOrder, error)
	FilterPOItem(FilterPOItem) (*[]PurchaseOrderItem, error)
	UpdatePOItem(POItemUpdate) (int64, error)
	//Receive Management
	GetReceiveList(ReceiveFilter) (int, []Transaction, error)
	//SO Management
	GetSalesOrderByID(int64) (*SalesOrder, error)
	GetSalesOrderList(SalesOrderFilter) (int, []SalesOrder, error)
	FilterSOItem(FilterSOItem) (*[]SalesOrderItem, error)
	UpdateSOItem(SOItemUpdate) (int64, error)
	//PickingOrder Management
	GetPickingOrderByID(int64) (*PickingOrder, error)
	GetPickingOrderList(PickingOrderFilter) (int, []PickingOrder, error)
	FilterPickingOrderItem(FilterPickingOrderItem) (*[]PickingOrderItem, error)
	FilterPickingOrderDetail(FilterPickingOrderDetail) (*[]PickingOrderDetail, error)
	CreatePickingOrder([]string, string) (int64, error)
	CreatePickingTransaction(info PickingTransactionNew) (int64, bool, error)
	CreatePackingTransaction(info PackingTransactionNew) (int64, error)
}

func (s *inventoryService) GetItemByID(id int64) (Item, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	item, err := repo.GetItemByID(id)
	return item, err
}

func (s *inventoryService) GetItemBySKU(sku string) (Item, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	item, err := repo.GetItemBySKU(sku)
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

func (s *inventoryService) UpdatePOItem(info POItemUpdate) (int64, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	affected, err := repo.UpdatePOItem(info)
	return affected, err
}

func (s *inventoryService) GetReceiveList(filter ReceiveFilter) (int, []Transaction, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	count, err := repo.GetTransactionCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetTransactionList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *inventoryService) GetSalesOrderByID(id int64) (*SalesOrder, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	salesOrder, err := repo.GetSalesOrderByID(id)
	return salesOrder, err
}

func (s *inventoryService) GetSalesOrderList(filter SalesOrderFilter) (int, []SalesOrder, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	count, err := repo.GetSalesOrderCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetSalesOrderList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *inventoryService) FilterSOItem(filter FilterSOItem) (*[]SalesOrderItem, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	items, err := repo.FilterSOItem(filter)
	return items, err
}

func (s *inventoryService) UpdateSOItem(info SOItemUpdate) (int64, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	affected, err := repo.UpdateSOItem(info)
	return affected, err
}

func (s *inventoryService) GetPickingOrderByID(id int64) (*PickingOrder, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	salesOrder, err := repo.GetPickingOrderByID(id)
	return salesOrder, err
}

func (s *inventoryService) GetPickingOrderList(filter PickingOrderFilter) (int, []PickingOrder, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	count, err := repo.GetPickingOrderCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetPickingOrderList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *inventoryService) FilterPickingOrderItem(filter FilterPickingOrderItem) (*[]PickingOrderItem, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	items, err := repo.FilterPickingOrderItem(filter)
	return items, err
}

func (s *inventoryService) FilterPickingOrderDetail(filter FilterPickingOrderDetail) (*[]PickingOrderDetail, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	details, err := repo.FilterPickingOrderDetail(filter)
	return details, err
}
func (s *inventoryService) CreatePickingOrder(soIDs []string, user string) (int64, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	err := repo.CheckSOExist(soIDs)
	if err != nil {
		return 0, err
	}
	err = repo.CheckSOStock(soIDs)
	if err != nil {
		return 0, err
	}
	created, err := repo.CreatePickingOrder(soIDs, user)
	return created, err
}
func (s *inventoryService) CreatePickingTransaction(info PickingTransactionNew) (int64, bool, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	created, ifFullPicked, err := repo.CreatePickingTransaction(info)
	return created, ifFullPicked, err
}
func (s *inventoryService) CreatePackingTransaction(info PackingTransactionNew) (int64, error) {
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	created, err := repo.CreatePackingTransaction(info)
	return created, err
}
