package inventory

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

type inventoryRepository struct {
	conn *sqlx.DB
}

func NewInventoryRepository(connection *sqlx.DB) InventoryRepository {
	return &inventoryRepository{
		conn: connection,
	}
}

type InventoryRepository interface {
	//Item Management
	GetItemByID(id int64) (Item, error)
	GetItemCount(filter ItemFilter) (int, error)
	GetItemList(filter ItemFilter) ([]Item, error)
	//PurchaseOrder Management
	GetPurchaseOrderByID(id int64) (*PurchaseOrder, error)
	GetPurchaseOrderCount(filter PurchaseOrderFilter) (int, error)
	GetPurchaseOrderList(filter PurchaseOrderFilter) ([]PurchaseOrder, error)
	FilterPOItem(filter FilterPOItem) (*[]PurchaseOrderItem, error)
}

func (r *inventoryRepository) GetItemByID(id int64) (Item, error) {
	var item Item
	err := r.conn.Get(&item, "SELECT * FROM i_items WHERE id = ? ", id)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func (r *inventoryRepository) GetItemCount(filter ItemFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku like ?"), append(args, "%"+v+"%")
	}
	if v := filter.Name; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM i_items 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *inventoryRepository) GetItemList(filter ItemFilter) ([]Item, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku like ?"), append(args, "%"+v+"%")
	}
	if v := filter.Name; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var items []Item
	err := r.conn.Select(&items, `
		SELECT * 
		FROM i_items 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *inventoryRepository) GetPurchaseOrderByID(id int64) (*PurchaseOrder, error) {
	var purchaseOrder PurchaseOrder
	err := r.conn.Get(&purchaseOrder, "SELECT * FROM i_purchase_orders WHERE id = ? ", id)
	if err != nil {
		return nil, err
	}
	return &purchaseOrder, nil
}

func (r *inventoryRepository) GetPurchaseOrderCount(filter PurchaseOrderFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.PONumber; v != "" {
		where, args = append(where, "po_number like ?"), append(args, "%"+v+"%")
	}
	if v := filter.VendorName; v != "" {
		where, args = append(where, "vendor_name like ?"), append(args, "%"+v+"%")
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM i_purchase_orders 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *inventoryRepository) GetPurchaseOrderList(filter PurchaseOrderFilter) ([]PurchaseOrder, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.PONumber; v != "" {
		where, args = append(where, "po_number like ?"), append(args, "%"+v+"%")
	}
	if v := filter.VendorName; v != "" {
		where, args = append(where, "vendor_name like ?"), append(args, "%"+v+"%")
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var purchaseOrders []PurchaseOrder
	err := r.conn.Select(&purchaseOrders, `
		SELECT * 
		FROM i_purchase_orders 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func (r *inventoryRepository) FilterPOItem(filter FilterPOItem) (*[]PurchaseOrderItem, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.POID; v != 0 {
		where, args = append(where, "po_id = ?"), append(args, v)
	}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku = ?"), append(args, v)
	}
	var items []PurchaseOrderItem
	err := r.conn.Select(&items, `
		SELECT * 
		FROM i_purchase_order_items 
		WHERE `+strings.Join(where, " AND ")+`
	`, args...)
	if err != nil {
		return nil, err
	}
	return &items, nil
}
