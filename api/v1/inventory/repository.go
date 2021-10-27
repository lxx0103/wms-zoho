package inventory

import (
	"errors"
	"fmt"
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
	UpdatePOItem(info POItemUpdate) (bool, error)
	//Transaction
	CreateTransaction(t TransactionNew) error
	GetTransactionCount(filter ReceiveFilter) (int, error)
	GetTransactionList(filter ReceiveFilter) ([]Transaction, error)
	GetTransactionToPick(sku string) (*Transaction, error)
	//SalesOrder Management
	GetSalesOrderByID(id int64) (*SalesOrder, error)
	GetSalesOrderCount(filter SalesOrderFilter) (int, error)
	GetSalesOrderList(filter SalesOrderFilter) ([]SalesOrder, error)
	FilterSOItem(filter FilterSOItem) (*[]SalesOrderItem, error)
	UpdateSOItem(info SOItemUpdate) (int64, error)
	CheckSOExist(ids []string) error
	CheckSOStock(ids []string) error
	UpdateSOStatus(ids string, user string) error
	//PickingOrder Management
	GetPickingOrderByID(id int64) (*PickingOrder, error)
	GetPickingOrderCount(filter PickingOrderFilter) (int, error)
	GetPickingOrderList(filter PickingOrderFilter) ([]PickingOrder, error)
	FilterPickingOrderItem(filter FilterPickingOrderItem) (*[]PickingOrderItem, error)
	FilterPickingOrderDetail(filter FilterPickingOrderDetail) (*[]PickingOrderDetail, error)
	CreatePickingOrder([]string, string) (int64, error)
	CreatePickingOrderDetail(PickingOrderDetailNew) (int64, error)
	CreatePickingTransaction(PickingTransactionNew) (int64, bool, error)
	CreatePackingTransaction(PackingTransactionNew) (int64, error)
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
			balance,
			shelf_code,
			shelf_location,
			location_code,
			location_level,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, t.POID, t.PONumber, t.ItemName, t.SKU, t.Quantity, t.Quantity, t.ShelfCode, t.ShelfLocation, t.LocationCode, t.LocationLevel, 1, time.Now(), t.User, time.Now(), t.User)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
		Update i_items SET 
		stock_available = stock_available + ?,
		stock = stock + ?,
		updated = ?,
		updated_by = ? 
		WHERE sku = ?
	`, t.Quantity, t.Quantity, time.Now(), t.User, t.SKU)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
func (r *inventoryRepository) UpdatePOItem(info POItemUpdate) (bool, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()
	_, err = tx.Exec(`
		Update i_purchase_order_items SET 
		quantity_received = quantity_received + ?,
		updated = ?,
		updated_by = ? 
		WHERE po_id = ?
		AND sku = ?
	`, info.Quantity, time.Now(), info.User, info.POID, info.SKU)
	if err != nil {
		return false, err
	}
	isCompleted := false
	var orderFullPicked int64
	row := tx.QueryRow(`SELECT id FROM i_purchase_order_items WHERE po_id = ? AND quantity > quantity_received LIMIT 1`, info.POID)
	err = row.Scan(&orderFullPicked)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			_, err = tx.Exec(`
				Update i_purchase_orders SET 
				status = "COMPLETED",
				updated = ?,
				updated_by = ? 
				WHERE id = ?
			`, time.Now(), info.User, info.POID)
			if err != nil {
				return false, err
			}
			isCompleted = true
		} else {
			return false, err
		}
	}
	tx.Commit()
	return isCompleted, nil
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

func (r *inventoryRepository) GetTransactionToPick(sku string) (*Transaction, error) {
	var transaction Transaction
	err := r.conn.Get(&transaction, `
		SELECT * 
		FROM i_transactions 
		WHERE sku = ?
		AND balance > 0
		ORDER BY created ASC
		LIMIT 1
	`, sku)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
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
	if v := filter.Status; v != "" {
		where, args = append(where, "status = ?"), append(args, v)
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
	if v := filter.Status; v != "" {
		where, args = append(where, "status = ?"), append(args, v)
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

func (r *inventoryRepository) CheckSOStock(ids []string) error {
	var count int
	idstring := strings.Join(ids[:], ",")
	fmt.Println(idstring)
	err := r.conn.Get(&count, `
		SELECT count(1) AS count FROM (
			SELECT item_id,sum(quantity)-sum(quantity_picked) AS topick, stock_available 
			FROM i_sales_order_items isoi
			LEFT JOIN i_items ii
			ON ii.id = isoi.item_id
			WHERE so_id in (`+idstring+`)
			GROUP BY item_id
			HAVING topick > stock_available
		) as table1
	`)
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("STOCK ERROR")
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
		where, args = append(where, "picking_date = ?"), append(args, v)
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
		where, args = append(where, "picking_date = ?"), append(args, v)
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var pickingOrders []PickingOrder
	err := r.conn.Select(&pickingOrders, `
		SELECT * 
		FROM i_picking_orders 
		WHERE `+strings.Join(where, " AND ")+`
		ORDER BY ID DESC
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
		where, args = append(where, "picking_order_id = ?"), append(args, v)
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

func (r *inventoryRepository) FilterPickingOrderDetail(filter FilterPickingOrderDetail) (*[]PickingOrderDetail, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.POID; v != 0 {
		where, args = append(where, "picking_order_id = ?"), append(args, v)
	}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku = ?"), append(args, v)
	}
	if v := filter.LocationCode; v != "" {
		where, args = append(where, "location_code = ?"), append(args, v)
	}
	var details []PickingOrderDetail
	err := r.conn.Select(&details, `
		SELECT * 
		FROM i_picking_order_details 
		WHERE `+strings.Join(where, " AND ")+`
	`, args...)
	if err != nil {
		return nil, err
	}
	return &details, nil
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

func (r *inventoryRepository) CreatePickingOrderDetail(info PickingOrderDetailNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	var exists int64
	var pickingDetailID int64
	err = r.conn.Get(&exists, `SELECT id FROM i_picking_order_details WHERE picking_order_id = ? AND location_code = ? LIMIT 1`, info.POID, info.LocationCode)
	if err == nil {
		result, err := tx.Exec(`
			UPDATE i_picking_order_details
			SET 
				quantity = quantity + ?,
				updated = ?,
				updated_by = ?
			WHERE id = ?
		`, info.Quantity, time.Now(), info.UserName, exists)
		if err != nil {
			return 0, err
		}
		_, err = result.RowsAffected()
		if err != nil {
			return 0, err
		}
		pickingDetailID = exists
	} else {
		result, err := tx.Exec(`
			INSERT INTO i_picking_order_details
			(
				picking_order_id,
				shelf_location,
				shelf_code,
				location_level,
				location_code,
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
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, info.POID, info.ShelfLocation, info.ShelfCode, info.LocationLevel, info.LocationCode, info.ItemID, info.SKU, info.ZohoItemID, info.Name, info.Quantity, info.QuantityPicked, 1, time.Now(), info.UserName, time.Now(), info.UserName)
		if err != nil {
			return 0, err
		}
		pickingDetailID, err = result.LastInsertId()
		if err != nil {
			return 0, err
		}
	}
	_, err = tx.Exec(`		
		Update i_transactions 
		SET balance = (balance - ?),
		updated = ?,
		updated_by = ?
		WHERE id = ?
	`, info.Quantity, time.Now(), info.UserName, info.TransactionID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update s_locations SET 
		can_pick = can_pick - ?,
		quantity = quantity - ?,
		available = available + ?,
		updated = ?,
		updated_by = ? 
		WHERE code = ?
	`, info.Quantity, info.Quantity, info.Quantity, time.Now(), info.UserName, info.LocationCode)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update i_items SET 
		stock_available = stock_available - ?,
		stock_picking = stock_picking + ?,
		updated = ?,
		updated_by = ? 
		WHERE id = ?
	`, info.Quantity, info.Quantity, time.Now(), info.UserName, info.ItemID)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return pickingDetailID, nil
}

func (r *inventoryRepository) CreatePickingTransaction(info PickingTransactionNew) (int64, bool, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, false, err
	}
	defer tx.Rollback()

	result, err := tx.Exec(`
		INSERT INTO i_picking_transactions
		(
			po_id,
			item_name,
			sku,
			quantity,
			shelf_code,
			shelf_location,
			location_code,
			location_level,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, info.POID, info.ItemName, info.SKU, info.Quantity, info.ShelfCode, info.ShelfLocation, info.LocationCode, info.LocationLevel, 1, time.Now(), info.UserName, time.Now(), info.UserName)
	if err != nil {
		return 0, false, err
	}
	transactionID, err := result.LastInsertId()
	if err != nil {
		return 0, false, err
	}
	_, err = tx.Exec(`		
		Update i_picking_order_details 
		SET quantity_picked = quantity_picked + ?,
		updated = ?,
		updated_by = ?
		WHERE picking_order_id = ?
		AND location_code = ?
	`, info.Quantity, time.Now(), info.UserName, info.POID, info.LocationCode)
	if err != nil {
		return 0, false, err
	}
	_, err = tx.Exec(`		
		Update i_picking_order_items 
		SET quantity_picked = quantity_picked + ?,
		updated = ?,
		updated_by = ?
		WHERE picking_order_id = ?
		AND sku = ?
	`, info.Quantity, time.Now(), info.UserName, info.POID, info.SKU)
	if err != nil {
		return 0, false, err
	}
	_, err = tx.Exec(`
		Update i_items SET 
		stock_picking = stock_picking - ?,
		stock_packing = stock_packing + ?,
		updated = ?,
		updated_by = ? 
		WHERE sku = ?
	`, info.Quantity, info.Quantity, time.Now(), info.UserName, info.SKU)
	if err != nil {
		return 0, false, err
	}
	var orderFullPicked int64
	isFullPicked := false
	row := tx.QueryRow(`SELECT id FROM i_picking_order_items WHERE picking_order_id = ? AND quantity > quantity_picked LIMIT 1`, info.POID)
	err = row.Scan(&orderFullPicked)
	if err != nil {
		fmt.Println(err)
		if err.Error() == "sql: no rows in result set" {
			_, err = tx.Exec(`
				Update i_picking_orders SET 
				status = "COMPLETED",
				updated = ?,
				updated_by = ? 
				WHERE id = ?
			`, time.Now(), info.UserName, info.POID)
			if err != nil {
				return 0, false, err
			}
			isFullPicked = true
		} else {
			return 0, false, err
		}
	}
	tx.Commit()
	return transactionID, isFullPicked, nil
}

func (r *inventoryRepository) CreatePackingTransaction(info PackingTransactionNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	result, err := tx.Exec(`
		INSERT INTO i_packing_transactions
		(
			so_id,
			item_name,
			sku,
			quantity,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, info.SOID, info.ItemName, info.SKU, info.Quantity, 1, time.Now(), info.UserName, time.Now(), info.UserName)
	if err != nil {
		return 0, err
	}
	transactionID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update i_sales_order_items SET 
		quantity_packed = quantity_packed + ?,
		updated = ?,
		updated_by = ? 
		WHERE so_id = ?
		AND sku = ?
	`, info.Quantity, time.Now(), info.UserName, info.SOID, info.SKU)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update i_items SET 
		stock = stock - ?,
		stock_packing = stock_packing - ?,
		updated = ?,
		updated_by = ? 
		WHERE sku = ?
	`, info.Quantity, info.Quantity, time.Now(), info.UserName, info.SKU)
	if err != nil {
		return 0, err
	}
	var orderFullPacked int64
	row := tx.QueryRow(`SELECT id FROM i_sales_order_items WHERE so_id = ? AND quantity > quantity_picked LIMIT 1`, info.SOID)
	err = row.Scan(&orderFullPacked)
	if err != nil {
		fmt.Println(err)
		if err.Error() == "sql: no rows in result set" {
			_, err = tx.Exec(`
				Update i_sales_orders SET 
				status = "PACKED",
				updated = ?,
				updated_by = ? 
				WHERE id = ?
			`, time.Now(), info.UserName, info.SOID)
			if err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}
	tx.Commit()
	return transactionID, nil
}

func (r *inventoryRepository) UpdateSOStatus(ids string, user string) error {
	_, err := r.conn.Exec(`
		Update i_sales_orders SET
		status = "PICKED",
		updated = ?,
		updated_by = ?
		WHERE id in (`+ids+`)
	`, time.Now(), user)
	if err != nil {
		return err
	}
	_, err = r.conn.Exec(`
		Update i_sales_order_items SET
		quantity_picked = quantity,
		updated = ?,
		updated_by = ?
		WHERE so_id in (`+ids+`)
	`, time.Now(), user)
	if err != nil {
		return err
	}
	return nil
}
