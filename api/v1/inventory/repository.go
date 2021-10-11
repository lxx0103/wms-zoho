package inventory

import (
	"strings"
	"time"

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
	GetItemBySKU(sku string) (Item, error)
	GetItemCount(filter ItemFilter) (int, error)
	GetItemList(filter ItemFilter) ([]Item, error)
	//PurchaseOrder Management
	GetPurchaseOrderByID(id int64) (*PurchaseOrder, error)
	GetPurchaseOrderCount(filter PurchaseOrderFilter) (int, error)
	GetPurchaseOrderList(filter PurchaseOrderFilter) ([]PurchaseOrder, error)
	FilterPOItem(filter FilterPOItem) (*[]PurchaseOrderItem, error)
	UpdatePOItem(info POItemUpdate) (int64, error)
	//Transaction
	CreateTransaction(t TransactionNew) error
	GetTransactionCount(filter ReceiveFilter) (int, error)
	GetTransactionList(filter ReceiveFilter) ([]Transaction, error)
	//SalesOrder Management
	GetSalesOrderByID(id int64) (*SalesOrder, error)
	GetSalesOrderCount(filter SalesOrderFilter) (int, error)
	GetSalesOrderList(filter SalesOrderFilter) ([]SalesOrder, error)
	FilterSOItem(filter FilterSOItem) (*[]SalesOrderItem, error)
	UpdateSOItem(info SOItemUpdate) (int64, error)
}

func (r *inventoryRepository) GetItemByID(id int64) (Item, error) {
	var item Item
	err := r.conn.Get(&item, "SELECT * FROM i_items WHERE id = ? ", id)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func (r *inventoryRepository) GetItemBySKU(sku string) (Item, error) {
	var item Item
	err := r.conn.Get(&item, "SELECT * FROM i_items WHERE sku = ? ", sku)
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
	if v := filter.ReceiveDate; v != "" {
		where, args = append(where, "expected_delivery_date = ?"), append(args, v)
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
	if v := filter.ReceiveDate; v != "" {
		where, args = append(where, "expected_delivery_date = ?"), append(args, v)
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

func (r *inventoryRepository) CreateTransaction(t TransactionNew) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO i_transactions
		(
			po_id,
			po_number,
			item_name,
			sku,
			quantity,
			shelf_code,
			shelf_location,
			location_code,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, t.POID, t.PONumber, t.ItemName, t.SKU, t.Quantity, t.ShelfCode, t.ShelfLocation, t.LocationCode, 1, time.Now(), t.User, time.Now(), t.User)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
func (r *inventoryRepository) UpdatePOItem(info POItemUpdate) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		Update i_purchase_order_items SET 
		quantity_received = quantity_received + ?,
		updated = ?,
		updated_by = ? 
		WHERE po_id = ?
		AND sku = ?
	`, info.Quantity, time.Now(), info.User, info.POID, info.SKU)
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

func (r *inventoryRepository) GetTransactionCount(filter ReceiveFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.PONumber; v != "" {
		where, args = append(where, "po_number = ?"), append(args, v)
	}
	if v := filter.POID; v != 0 {
		where, args = append(where, "po_id = ?"), append(args, v)
	}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku = ?"), append(args, v)
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM i_transactions 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *inventoryRepository) GetTransactionList(filter ReceiveFilter) ([]Transaction, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.PONumber; v != "" {
		where, args = append(where, "po_number = ?"), append(args, v)
	}
	if v := filter.POID; v != 0 {
		where, args = append(where, "po_id = ?"), append(args, v)
	}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku = ?"), append(args, v)
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var transactions []Transaction
	err := r.conn.Select(&transactions, `
		SELECT * 
		FROM i_transactions 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *inventoryRepository) GetSalesOrderByID(id int64) (*SalesOrder, error) {
	var salesOrder SalesOrder
	err := r.conn.Get(&salesOrder, "SELECT * FROM i_sales_orders WHERE id = ? ", id)
	if err != nil {
		return nil, err
	}
	return &salesOrder, nil
}

func (r *inventoryRepository) GetSalesOrderCount(filter SalesOrderFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.SONumber; v != "" {
		where, args = append(where, "so_number like ?"), append(args, "%"+v+"%")
	}
	if v := filter.CustomerName; v != "" {
		where, args = append(where, "customer_name like ?"), append(args, "%"+v+"%")
	}
	if v := filter.SalesName; v != "" {
		where, args = append(where, "sales_name like ?"), append(args, "%"+v+"%")
	}
	if v := filter.OrderDate; v != "" {
		where, args = append(where, "so_date = ?"), append(args, v)
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM i_sales_orders 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *inventoryRepository) GetSalesOrderList(filter SalesOrderFilter) ([]SalesOrder, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.SONumber; v != "" {
		where, args = append(where, "so_number like ?"), append(args, "%"+v+"%")
	}
	if v := filter.CustomerName; v != "" {
		where, args = append(where, "customer_name like ?"), append(args, "%"+v+"%")
	}
	if v := filter.SalesName; v != "" {
		where, args = append(where, "sales_name like ?"), append(args, "%"+v+"%")
	}
	if v := filter.OrderDate; v != "" {
		where, args = append(where, "so_date = ?"), append(args, v)
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var salesOrders []SalesOrder
	err := r.conn.Select(&salesOrders, `
		SELECT * 
		FROM i_sales_orders 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return salesOrders, nil
}

func (r *inventoryRepository) FilterSOItem(filter FilterSOItem) (*[]SalesOrderItem, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.SOID; v != 0 {
		where, args = append(where, "so_id = ?"), append(args, v)
	}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku = ?"), append(args, v)
	}
	var items []SalesOrderItem
	err := r.conn.Select(&items, `
		SELECT * 
		FROM i_sales_order_items 
		WHERE `+strings.Join(where, " AND ")+`
	`, args...)
	if err != nil {
		return nil, err
	}
	return &items, nil
}
func (r *inventoryRepository) UpdateSOItem(info SOItemUpdate) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		Update i_sales_order_items SET 
		quantity_received = quantity_received + ?,
		updated = ?,
		updated_by = ? 
		WHERE so_id = ?
		AND sku = ?
	`, info.Quantity, time.Now(), info.User, info.SOID, info.SKU)
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
