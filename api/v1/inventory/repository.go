package inventory

import (
	"errors"
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
	CheckSOExist(ids []string) error
	//PickingOrder Management
	GetPickingOrderByID(id int64) (*PickingOrder, error)
	GetPickingOrderCount(filter PickingOrderFilter) (int, error)
	GetPickingOrderList(filter PickingOrderFilter) ([]PickingOrder, error)
	FilterPickingOrderItem(filter FilterPickingOrderItem) (*[]PickingOrderItem, error)
	CreatePickingOrder([]string, string) (int64, error)
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

func (r *inventoryRepository) CheckSOExist(ids []string) error {
	var count int
	idstring := strings.Join(ids[:], ",")
	err := r.conn.Get(&count, "SELECT count(1) as count FROM i_sales_orders WHERE id in  ("+idstring+") AND status != 'PICKING'")
	if err != nil {
		return err
	}
	if count != len(ids) {
		return errors.New("COUNT ERROR")
	}
	return nil
}

func (r *inventoryRepository) GetPickingOrderByID(id int64) (*PickingOrder, error) {
	var pickingOrder PickingOrder
	err := r.conn.Get(&pickingOrder, "SELECT * FROM i_picking_orders WHERE id = ? ", id)
	if err != nil {
		return nil, err
	}
	return &pickingOrder, nil
}

func (r *inventoryRepository) GetPickingOrderCount(filter PickingOrderFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.OrderNumber; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	if v := filter.UserName; v != "" {
		where, args = append(where, "created_by like ?"), append(args, "%"+v+"%")
	}
	if v := filter.OrderDate; v != "" {
		where, args = append(where, "so_date = ?"), append(args, v)
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM i_picking_orders 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *inventoryRepository) GetPickingOrderList(filter PickingOrderFilter) ([]PickingOrder, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.OrderNumber; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	if v := filter.UserName; v != "" {
		where, args = append(where, "created_by like ?"), append(args, "%"+v+"%")
	}
	if v := filter.OrderDate; v != "" {
		where, args = append(where, "so_date = ?"), append(args, v)
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var pickingOrders []PickingOrder
	err := r.conn.Select(&pickingOrders, `
		SELECT * 
		FROM i_picking_orders 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return pickingOrders, nil
}

func (r *inventoryRepository) FilterPickingOrderItem(filter FilterPickingOrderItem) (*[]PickingOrderItem, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.POID; v != 0 {
		where, args = append(where, "po_id = ?"), append(args, v)
	}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku = ?"), append(args, v)
	}
	var items []PickingOrderItem
	err := r.conn.Select(&items, `
		SELECT * 
		FROM i_picking_order_items 
		WHERE `+strings.Join(where, " AND ")+`
	`, args...)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (r *inventoryRepository) CreatePickingOrder(ids []string, user string) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO i_picking_orders
		(
			name,
			sales_orders,
			picking_date,
			status,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, time.Now().Format("2006-01-02 15:04:05")+"_by_"+user, strings.Join(ids[:], ","), time.Now().Format("2006-01-02"), "TOPICK", 1, time.Now(), user, time.Now(), user)
	if err != nil {
		return 0, err
	}
	pickingID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`		
		INSERT INTO i_picking_order_items 
		(
			picking_order_id,
			item_id,
			sku,
			zoho_item_id,
			name,
			quantity,
			quantity_picked,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		SELECT 
		?,
		item_id,
		sku,
		zoho_item_id,
		name,
		sum(quantity)-sum(quantity_picked),
		0,
		1,
    	?,
		?,
		?,
		?
		FROM i_sales_order_items 
		WHERE so_id in (`+strings.Join(ids[:], ",")+`) group by item_id,sku,zoho_item_id,name
	`, pickingID, time.Now(), user, time.Now(), user)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update i_sales_orders SET 
		status = "PICKING",
		updated = ?,
		updated_by = ? 
		WHERE id in (`+strings.Join(ids[:], ",")+`)
	`, time.Now(), user)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return pickingID, nil
}
