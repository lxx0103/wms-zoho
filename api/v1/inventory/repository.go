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
	CancelReceive(int64, string) (int64, error)
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
	CancelPicking(int64, string) (int64, error)
	CancelPacking(int64, string) (int64, error)
	//Adjustment
	CreateAdjustment(AdjustmentInfo) (int64, error)
	GetAdjustmentCount(AdjustmentFilter) (int, error)
	GetAdjustmentList(AdjustmentFilter) ([]Adjustment, error)
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
	where, args := []string{"1 = 1  AND po_number != \"adjustment\""}, []interface{}{}
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
	where, args := []string{"1 = 1 AND po_number != \"adjustment\""}, []interface{}{}
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
	_, err = tx.Exec(`		
		INSERT INTO i_picking_order_logs 
		(
			picking_order_id,
			transaction_id,
			quantity,
			created,
			created_by,
			updated,
			updated_by
		)
		Values (?, ?, ?, ?, ?, ?, ?)
	`, info.POID, info.TransactionID, info.Quantity, time.Now(), info.UserName, time.Now(), info.UserName)
	if err != nil {
		return 0, err
	}
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
	row := tx.QueryRow(`SELECT id FROM i_sales_order_items WHERE so_id = ? AND quantity > quantity_packed LIMIT 1`, info.SOID)
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

func (r *inventoryRepository) CancelReceive(poID int64, user string) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	var canCancel int
	row := tx.QueryRow(`SELECT count(1) FROM i_transactions WHERE po_id = ? AND quantity > balance AND po_number != "adjustment"`, poID)
	err = row.Scan(&canCancel)
	if err != nil {
		return 0, err
	}
	if canCancel != 0 {
		return 0, errors.New("YOU CAN NOT CANCEL THE RECEIVING WHEN IT IS ALREADY PICKED")
	}

	_, err = tx.Exec(`
		UPDATE i_purchase_order_items poi
		LEFT JOIN i_items i
		ON poi.item_id = i.id
		SET i.stock_available = i.stock_available - poi.quantity_received,
		i.stock = i.stock - poi.quantity_received,
		i.updated = ?,
		i.updated_by = ?
		WHERE poi.po_id = ?
	`, time.Now(), user, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		UPDATE i_transactions t
		LEFT JOIN s_locations l
		ON t.location_code = l.code
		SET l.quantity = l.quantity - t.quantity,
		l.can_pick = l.can_pick - t.quantity,
		l.available = l.available + t.quantity,
		l.updated = ?,
		l.updated_by = ?
		WHERE t.po_id = ?
	`, time.Now(), user, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update i_purchase_order_items SET 
		quantity_received = 0,
		updated = ?,
		updated_by = ?
		WHERE po_id = ?
	`, time.Now(), user, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update i_purchase_orders SET 
		status = "CONFIRM",
		updated = ?,
		updated_by = ?
		WHERE id = ?
	`, time.Now(), user, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		DELETE FROM  i_transactions
		WHERE po_id = ?
	`, poID)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return 1, nil
}

func (r *inventoryRepository) CancelPicking(poID int64, user string) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	var ids string
	row := tx.QueryRow(`SELECT sales_orders FROM i_picking_orders WHERE id = ? LIMIT 1`, poID)
	err = row.Scan(&ids)
	if err != nil {
		return 0, err
	}
	var canCancel int
	row1 := tx.QueryRow(`SELECT count(1) FROM i_sales_order_items WHERE so_id in (` + ids + `) AND quantity_packed > 0`)
	err = row1.Scan(&canCancel)
	if err != nil {
		return 0, err
	}
	if canCancel != 0 {
		return 0, errors.New("YOU CAN NOT CANCEL THE PICKING WHEN SALESORDER IS ALREADY PACKED")
	}
	_, err = tx.Exec(`
		UPDATE i_picking_order_items poi
		LEFT JOIN i_items i
		ON poi.item_id = i.id
		SET i.stock_picking = i.stock_picking + poi.quantity_picked,
		i.stock_packing = i.stock_packing - poi.quantity_picked,
		i.updated = ?,
		i.updated_by = ?
		where poi.picking_order_id = ?
	`, time.Now(), user, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		UPDATE i_picking_order_items poi
		LEFT JOIN i_items i
		ON poi.item_id = i.id
		SET i.stock_available = i.stock_available + poi.quantity,
		i.stock_picking = i.stock_picking - poi.quantity,
		i.updated = ?,
		i.updated_by = ?
		where poi.picking_order_id = ?
	`, time.Now(), user, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		UPDATE i_picking_order_details pod
		LEFT JOIN s_locations l
		ON pod.location_code = l.code
		SET l.quantity = l.quantity + pod.quantity,
		l.can_pick = l.can_pick + pod.quantity,
		l.available = l.available - pod.quantity,
		l.updated = ?,
		l.updated_by = ?
		WHERE pod.picking_order_id = ?
	`, time.Now(), user, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		UPDATE i_picking_order_logs pol
		LEFT JOIN i_transactions t
		ON pol.transaction_id = t.id
		SET t.balance = t.balance + pol.quantity,
		t.updated = ?,
		t.updated_by = ?
		WHERE pol.picking_order_id = ?
	`, time.Now(), user, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update i_sales_orders SET
		status = "CONFIRMED",
		updated = ?,
		updated_by = ?
		WHERE id in (`+ids+`)
	`, time.Now(), user)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update i_sales_order_items SET
		quantity_picked = 0,
		updated = ?,
		updated_by = ?
		WHERE so_id in (`+ids+`)
	`, time.Now(), user)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		DELETE FROM  i_picking_order_details
		WHERE picking_order_id = ?
	`, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		DELETE FROM  i_picking_order_items
		WHERE picking_order_id = ?
	`, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		DELETE FROM  i_picking_order_logs
		WHERE picking_order_id = ?
	`, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		DELETE FROM  i_picking_transactions
		WHERE po_id = ?
	`, poID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		DELETE FROM  i_picking_orders
		WHERE id = ?
	`, poID)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return 1, nil
}

func (r *inventoryRepository) CancelPacking(soID int64, user string) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		UPDATE i_sales_order_items soi
		LEFT JOIN i_items i
		ON soi.item_id = i.id
		SET i.stock_packing = i.stock_packing + soi.quantity_packed,
		i.stock = i.stock + soi.quantity_packed,
		i.updated = ?,
		i.updated_by = ?
		where soi.so_id = ?
	`, time.Now(), user, soID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update i_sales_orders SET
		status = "PICKED",
		updated = ?,
		updated_by = ?
		WHERE id =?
	`, time.Now(), user, soID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update i_sales_order_items SET 
		quantity_packed = 0,
		updated = ?,
		updated_by = ? 
		WHERE so_id = ?
	`, time.Now(), user, soID)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		DELETE FROM  i_packing_transactions 
		WHERE so_id = ?
	`, soID)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return 1, nil
}

func (r *inventoryRepository) CreateAdjustment(info AdjustmentInfo) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO i_adjustments
		(
			location_code,
			quantity,
			remark,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, info.Code, info.Quantity, info.Remark, time.Now(), info.UserName, time.Now(), info.UserName)
	if err != nil {
		return 0, err
	}
	adjustmentID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		UPDATE i_adjustments a
		LEFT join s_locations l
		ON a.location_code = l.code 
		LEFT join i_items i
		ON l.sku = i.sku
		LEFT JOIN s_shelves s
		ON l.shelf_id = s.id
		SET a.item_name = i.name,
		a.sku = l.sku,
		a.shelf_code = s.code,
		a.shelf_location = s.location,
		a.location_level = l.level,
		a.updated = ?,
		a.updated_by = ?
		where a.id = ?
	`, time.Now(), info.UserName, adjustmentID)
	if err != nil {
		return 0, err
	}
	var nowQuantity int64
	var capacity int64
	var canPick int64
	row := tx.QueryRow(`SELECT capacity,quantity,can_pick FROM s_locations WHERE code = ?`, info.Code)
	err = row.Scan(&capacity, &nowQuantity, &canPick)
	if err != nil {
		return 0, err
	}
	if info.Quantity < 0 {
		adjustQuantity := -info.Quantity
		if canPick-adjustQuantity < 0 {
			return 0, errors.New("QUANTITY NOT ENOUGH")
		}
		for adjustQuantity > 0 {
			var balance int64
			var id int64
			transactionRow := tx.QueryRow(`SELECT id, balance FROM i_transactions WHERE location_code = ? AND balance > 0  ORDER BY id ASC LIMIT 1`, info.Code)
			err = transactionRow.Scan(&id, &balance)
			if err != nil {
				return 0, err
			}
			if balance > adjustQuantity {
				_, err = tx.Exec(`
					UPDATE i_transactions
					SET balance = balance - ?,
					updated = ?,
					updated_by = ?
					where id = ?
				`, adjustQuantity, time.Now(), info.UserName, id)
				if err != nil {
					return 0, err
				}
				_, err := tx.Exec(`
					INSERT INTO i_adjustment_logs
					(
						adjustment_id,
						transaction_id,
						quantity,
						created,
						created_by,
						updated,
						updated_by
					)
					VALUES (?, ?, ?, ?, ?, ?, ?)
				`, adjustmentID, id, adjustQuantity, time.Now(), info.UserName, time.Now(), info.UserName)
				if err != nil {
					return 0, err
				}
				adjustQuantity = 0
			} else {
				_, err = tx.Exec(`
					UPDATE i_transactions
					SET balance = 0,
					updated = ?,
					updated_by = ?
					where id = ?
				`, time.Now(), info.UserName, id)
				if err != nil {
					return 0, err
				}
				_, err := tx.Exec(`
					INSERT INTO i_adjustment_logs
					(
						adjustment_id,
						transaction_id,
						quantity,
						created,
						created_by,
						updated,
						updated_by
					)
					VALUES (?, ?, ?, ?, ?, ?, ?)
				`, adjustmentID, id, balance, time.Now(), info.UserName, time.Now(), info.UserName)
				if err != nil {
					return 0, err
				}
				adjustQuantity = adjustQuantity - balance
			}
		}
	} else {
		if capacity < info.Quantity+nowQuantity {
			return 0, errors.New("CAPACITY NOT ENOUGH")
		}
		trans, err := tx.Exec(`
			INSERT INTO i_transactions
			(
				po_id,
				po_number,
				quantity,
				balance,
				location_code,
				enabled,
				created,
				created_by,
				updated,
				updated_by
			)
			VALUES (?, "adjustment", ?, ?, ?, ?, ?, ?, ?, ?)
		`, adjustmentID, info.Quantity, info.Quantity, info.Code, 1, time.Now(), info.UserName, time.Now(), info.UserName)
		if err != nil {
			return 0, err
		}

		transID, err := trans.LastInsertId()
		if err != nil {
			return 0, err
		}
		_, err = tx.Exec(`
			UPDATE i_transactions t
			LEFT join s_locations l
			ON t.location_code = l.code 
			LEFT join i_items i
			ON l.sku = i.sku
			LEFT JOIN s_shelves s
			ON l.shelf_id = s.id
			SET t.item_name = i.name,
			t.sku = l.sku,
			t.shelf_code = s.code,
			t.shelf_location = s.location,
			t.location_level = l.level,
			t.updated = ?,
			t.updated_by = ?
			where t.id = ?
		`, time.Now(), info.UserName, transID)
		if err != nil {
			return 0, err
		}
	}
	_, err = tx.Exec(`
		Update s_locations SET 
		quantity = quantity + ?,
		can_pick = can_pick + ?,
		available = available - ?,
		updated = ?,
		updated_by = ? 
		WHERE code = ?
	`, info.Quantity, info.Quantity, info.Quantity, time.Now(), info.UserName, info.Code)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update s_locations l
		LEFT JOIN i_items i
		ON l.sku = i.sku
		SET i.stock = i.stock + ?,
		i.stock_available = i.stock_available + ?,
		i.updated = ?,
		i.updated_by = ? 
		WHERE l.code = ?
	`, info.Quantity, info.Quantity, time.Now(), info.UserName, info.Code)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return 1, nil
}

func (r *inventoryRepository) GetAdjustmentCount(filter AdjustmentFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku like ?"), append(args, "%"+v+"%")
	}
	if v := filter.LocationCode; v != "" {
		where, args = append(where, "location_code like ?"), append(args, "%"+v+"%")
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM i_adjustments 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *inventoryRepository) GetAdjustmentList(filter AdjustmentFilter) ([]Adjustment, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku like ?"), append(args, "%"+v+"%")
	}
	if v := filter.LocationCode; v != "" {
		where, args = append(where, "location_code like ?"), append(args, "%"+v+"%")
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var adjustments []Adjustment
	err := r.conn.Select(&adjustments, `
		SELECT * 
		FROM i_adjustments 
		WHERE `+strings.Join(where, " AND ")+`
		ORDER BY id desc
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return adjustments, nil
}
