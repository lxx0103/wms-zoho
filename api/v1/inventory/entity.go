package inventory

import "time"

type Item struct {
	ID             int64     `db:"id" json:"id"`
	SKU            string    `db:"sku" json:"sku"`
	Name           string    `db:"name" json:"name"`
	ZohoItemID     string    `db:"zoho_item_id" json:"zoho_item_id"`
	Unit           string    `db:"unit" json:"unit"`
	Stock          int       `db:"stock" json:"stock"`
	StockAvailable int64     `db:"stock_available" json:"stock_available"`
	StockPicking   int64     `db:"stock_picking" json:"stock_picking"`
	StockPacking   int64     `db:"stock_packing" json:"stock_packing"`
	Enabled        string    `db:"enabled" json:"enabled"`
	Created        time.Time `db:"created" json:"created"`
	CreatedBy      string    `db:"created_by" json:"created_by"`
	Updated        time.Time `db:"updated" json:"updated"`
	UpdatedBy      string    `db:"updated_by" json:"updated_by"`
}

type PurchaseOrder struct {
	ID              int64     `db:"id" json:"id"`
	ZohoPOID        string    `db:"zoho_po_id" json:"zoho_po_id"`
	PONumber        string    `db:"po_number" json:"po_number"`
	PODate          time.Time `db:"po_date" json:"po_date"`
	ExpectedDate    time.Time `db:"expected_delivery_date" json:"expected_delivery_date"`
	ReferenceNumber string    `db:"reference_number" json:"reference_number"`
	Status          string    `db:"status" json:"status"`
	VendorID        string    `db:"vendor_id" json:"vendor_id"`
	VendorName      string    `db:"vendor_name" json:"vendor_name"`
	Enabled         string    `db:"enabled" json:"enabled"`
	Created         time.Time `db:"created" json:"created"`
	CreatedBy       string    `db:"created_by" json:"created_by"`
	Updated         time.Time `db:"updated" json:"updated"`
	UpdatedBy       string    `db:"updated_by" json:"updated_by"`
}

type PurchaseOrderItem struct {
	ID               int64     `db:"id" json:"id"`
	POID             int64     `db:"po_id" json:"po_id"`
	ItemID           int64     `db:"item_id" json:"item_id"`
	SKU              string    `db:"sku" json:"sku"`
	ZohoItemID       string    `db:"zoho_item_id" json:"zoho_item_id"`
	Name             string    `db:"name" json:"name"`
	Quantity         int64     `db:"quantity" json:"quantity"`
	QuantityReceived int64     `db:"quantity_received" json:"quantity_received"`
	Enabled          string    `db:"enabled" json:"enabled"`
	Created          time.Time `db:"created" json:"created"`
	CreatedBy        string    `db:"created_by" json:"created_by"`
	Updated          time.Time `db:"updated" json:"updated"`
	UpdatedBy        string    `db:"updated_by" json:"updated_by"`
}

type Transaction struct {
	ID            int64     `db:"id" json:"id"`
	POID          int64     `db:"po_id" json:"po_id"`
	PONumber      string    `db:"po_number" json:"po_number"`
	ItemName      string    `db:"item_name" json:"item_name"`
	SKU           string    `db:"sku" json:"sku"`
	Quantity      int64     `db:"quantity" json:"quantity"`
	Balance       int64     `db:"balance" json:"balance"`
	ShelfCode     string    `db:"shelf_code" json:"shelf_code"`
	ShelfLocation string    `db:"shelf_location" json:"shelf_location"`
	LocationCode  string    `db:"location_code" json:"location_code"`
	LocationLevel string    `db:"location_level" json:"location_level"`
	Enabled       string    `db:"enabled" json:"enabled"`
	Created       time.Time `db:"created" json:"created"`
	CreatedBy     string    `db:"created_by" json:"created_by"`
	Updated       time.Time `db:"updated" json:"updated"`
	UpdatedBy     string    `db:"updated_by" json:"updated_by"`
}

type SalesOrder struct {
	ID           int64     `db:"id" json:"id"`
	ZohoSOID     string    `db:"zoho_so_id" json:"zoho_so_id"`
	SONumber     string    `db:"so_number" json:"so_number"`
	SODate       time.Time `db:"so_date" json:"so_date"`
	Status       string    `db:"status" json:"status"`
	CustomerID   string    `db:"customer_id" json:"customer_id"`
	CustomerName string    `db:"customer_name" json:"customer_name"`
	SalesName    string    `db:"sales_name" json:"sales_name"`
	Enabled      string    `db:"enabled" json:"enabled"`
	HasPallet    bool      `db:"has_pallet" json:"has_pallet"`
	ExpectedDate time.Time `db:"expected_shipment_date" json:"expected_shipment_date"`
	Created      time.Time `db:"created" json:"created"`
	CreatedBy    string    `db:"created_by" json:"created_by"`
	Updated      time.Time `db:"updated" json:"updated"`
	UpdatedBy    string    `db:"updated_by" json:"updated_by"`
}

type SalesOrderItem struct {
	ID              int64     `db:"id" json:"id"`
	SOID            int64     `db:"so_id" json:"so_id"`
	ItemID          int64     `db:"item_id" json:"item_id"`
	SKU             string    `db:"sku" json:"sku"`
	ZohoItemID      string    `db:"zoho_item_id" json:"zoho_item_id"`
	Name            string    `db:"name" json:"name"`
	Quantity        int64     `db:"quantity" json:"quantity"`
	QuantityPicked  int64     `db:"quantity_picked" json:"quantity_picked"`
	QuantityPacked  int64     `db:"quantity_packed" json:"quantity_packed"`
	QuantityShipped int64     `db:"quantity_shipped" json:"quantity_shipped"`
	Enabled         string    `db:"enabled" json:"enabled"`
	Created         time.Time `db:"created" json:"created"`
	CreatedBy       string    `db:"created_by" json:"created_by"`
	Updated         time.Time `db:"updated" json:"updated"`
	UpdatedBy       string    `db:"updated_by" json:"updated_by"`
}

type SalesOrderPallet struct {
	ID        int64     `db:"id" json:"id"`
	SOID      int64     `db:"so_id" json:"so_id"`
	Name      string    `db:"name" json:"name"`
	Status    int       `db:"status" json:"status"`
	Created   time.Time `db:"created" json:"created"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	Updated   time.Time `db:"updated" json:"updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}
type PickingOrder struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	SalesOrders string    `db:"sales_orders" json:"sales_orders"`
	PickingDate time.Time `db:"picking_date" json:"picking_date"`
	Status      string    `db:"status" json:"status"`
	Enabled     string    `db:"enabled" json:"enabled"`
	Created     time.Time `db:"created" json:"created"`
	CreatedBy   string    `db:"created_by" json:"created_by"`
	Updated     time.Time `db:"updated" json:"updated"`
	UpdatedBy   string    `db:"updated_by" json:"updated_by"`
}

type PickingOrderItem struct {
	ID             int64     `db:"id" json:"id"`
	POID           int64     `db:"picking_order_id" json:"picking_order_id"`
	ItemID         int64     `db:"item_id" json:"item_id"`
	SKU            string    `db:"sku" json:"sku"`
	ZohoItemID     string    `db:"zoho_item_id" json:"zoho_item_id"`
	Name           string    `db:"name" json:"name"`
	Quantity       int64     `db:"quantity" json:"quantity"`
	QuantityPicked int64     `db:"quantity_picked" json:"quantity_picked"`
	Enabled        string    `db:"enabled" json:"enabled"`
	Created        time.Time `db:"created" json:"created"`
	CreatedBy      string    `db:"created_by" json:"created_by"`
	Updated        time.Time `db:"updated" json:"updated"`
	UpdatedBy      string    `db:"updated_by" json:"updated_by"`
}

type PickingOrderDetail struct {
	ID             int64     `db:"id" json:"id"`
	POID           int64     `db:"picking_order_id" json:"picking_order_id"`
	ShelfLocation  string    `db:"shelf_location" json:"shelf_location"`
	ShelfCode      string    `db:"shelf_code" json:"shelf_code"`
	LocationLevel  string    `db:"location_level" json:"location_level"`
	LocationCode   string    `db:"location_code" json:"location_code"`
	ItemID         int64     `db:"item_id" json:"item_id"`
	SKU            string    `db:"sku" json:"sku"`
	ZohoItemID     string    `db:"zoho_item_id" json:"zoho_item_id"`
	Name           string    `db:"name" json:"name"`
	Quantity       int64     `db:"quantity" json:"quantity"`
	QuantityPicked int64     `db:"quantity_picked" json:"quantity_picked"`
	Enabled        string    `db:"enabled" json:"enabled"`
	Created        time.Time `db:"created" json:"created"`
	CreatedBy      string    `db:"created_by" json:"created_by"`
	Updated        time.Time `db:"updated" json:"updated"`
	UpdatedBy      string    `db:"updated_by" json:"updated_by"`
}
type PickingTransaction struct {
	ID            int64     `db:"id" json:"id"`
	POID          int64     `db:"po_id" json:"po_id"`
	PONumber      string    `db:"po_number" json:"po_number"`
	ItemName      string    `db:"item_name" json:"item_name"`
	SKU           string    `db:"sku" json:"sku"`
	Quantity      int64     `db:"quantity" json:"quantity"`
	ShelfCode     string    `db:"shelf_code" json:"shelf_code"`
	ShelfLocation string    `db:"shelf_location" json:"shelf_location"`
	LocationCode  string    `db:"location_code" json:"location_code"`
	LocationLevel string    `db:"location_level" json:"location_level"`
	Enabled       string    `db:"enabled" json:"enabled"`
	Created       time.Time `db:"created" json:"created"`
	CreatedBy     string    `db:"created_by" json:"created_by"`
	Updated       time.Time `db:"updated" json:"updated"`
	UpdatedBy     string    `db:"updated_by" json:"updated_by"`
}

type Adjustment struct {
	ID            int64     `db:"id" json:"id"`
	SKU           string    `db:"sku" json:"sku"`
	ItemName      string    `db:"item_name" json:"item_name"`
	LocationCode  string    `db:"location_code" json:"location_code"`
	LocationLevel string    `db:"location_level" json:"location_level"`
	ShelfCode     string    `db:"shelf_code" json:"shelf_code"`
	ShelfLocation string    `db:"shelf_location" json:"shelf_location"`
	Quantity      int64     `db:"quantity" json:"quantity"`
	Remark        string    `db:"remark" json:"remark"`
	Created       time.Time `db:"created" json:"created"`
	CreatedBy     string    `db:"created_by" json:"created_by"`
	Updated       time.Time `db:"updated" json:"updated"`
	UpdatedBy     string    `db:"updated_by" json:"updated_by"`
}
