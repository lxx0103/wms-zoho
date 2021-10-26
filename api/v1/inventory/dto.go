package inventory

type ItemFilter struct {
	SKU      string `form:"sku" binding:"omitempty,max=64,min=1"`
	Name     string `form:"name" binding:"omitempty,max=64,min=1"`
	PageId   int    `form:"page_id" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=5,max=200"`
}

type ItemID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type PurchaseOrderFilter struct {
	PONumber    string `form:"po_number" binding:"omitempty,max=64,min=1"`
	VendorName  string `form:"vendor_name" binding:"omitempty,max=64,min=1"`
	ReceiveDate string `form:"receive_date" binding:"omitempty,datetime=2006-01-02"`
	PageId      int    `form:"page_id" binding:"required,min=1"`
	PageSize    int    `form:"page_size" binding:"required,min=5,max=200"`
}

type PurchaseOrderID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type ReceiveNew struct {
	POID     int64  `json:"po_id" binding:"required,min=1"`
	Barcode  string `json:"barcode" binding:"required,min=1,max=64"`
	Quantity int64  `json:"quantity" binding:"required,min=1"`
}

type FilterPOItem struct {
	POID int64  `json:"po_id"`
	SKU  string `json:"sku"`
}

type PODetail struct {
	PO    PurchaseOrder
	Items []PurchaseOrderItem
}

type StockInRes struct {
	Location    []StoctInLoc
	IsCompleted bool `json:"is_completed"`
}

type StoctInLoc struct {
	SKU           string `json:"sku"`
	ItemName      string `json:"item_name"`
	ShelfCode     string `json:"shelf_code"`
	ShelfLocation string `json:"shelf_location"`
	LocationCode  string `json:"location_code"`
	LocationLevel int64  `json:"location_level"`
	Quantity      int64  `json:"quantity"`
}

type TransactionNew struct {
	POID          int64  `json:"po_id"`
	PONumber      string `json:"po_number"`
	ItemName      string `json:"item_name"`
	SKU           string `json:"sku"`
	Quantity      int64  `json:"quantity"`
	ShelfCode     string `json:"shelf_code"`
	ShelfLocation string `json:"shelf_location"`
	LocationCode  string `json:"location_code"`
	LocationLevel int64  `json:"location_level"`
	User          string `json:"user"`
}

type POItemUpdate struct {
	POID     int64  `json:"po_id"`
	Quantity int64  `json:"quantity"`
	SKU      string `json:"sku"`
	User     string `json:"user"`
}
type ReceiveID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type ReceiveFilter struct {
	POID     int64  `form:"po_id" binding:"omitempty,min=1"`
	PONumber string `form:"po_number" binding:"omitempty,max=64,min=1"`
	SKU      string `form:"sku" binding:"omitempty,max=64,min=1"`
	PageId   int    `form:"page_id" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=5,max=200"`
}

type SalesOrderFilter struct {
	SONumber     string `form:"so_number" binding:"omitempty,max=64,min=1"`
	CustomerName string `form:"customer_name" binding:"omitempty,max=64,min=1"`
	SalesName    string `form:"sales_name" binding:"omitempty,max=64,min=1"`
	OrderDate    string `form:"order_date" binding:"omitempty,datetime=2006-01-02"`
	Status       string `form:"status" binding:"omitempty,oneof=confirmed picked packed picking invoiced draft fulfilled void partially_shipped shipped "`
	PageId       int    `form:"page_id" binding:"required,min=1"`
	PageSize     int    `form:"page_size" binding:"required,min=5,max=200"`
}

type FilterSOItem struct {
	SOID int64  `json:"so_id"`
	SKU  string `json:"sku"`
}
type SOItemUpdate struct {
	SOID     int64  `json:"so_id"`
	Quantity int64  `json:"quantity"`
	SKU      string `json:"sku"`
	User     string `json:"user"`
	Action   string `json:"action"`
}

type SalesOrderID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type SODetail struct {
	SO    SalesOrder
	Items []SalesOrderItem
}

type PickingOrderFilter struct {
	OrderNumber string `form:"order_number" binding:"omitempty,max=64,min=1"`
	UserName    string `form:"user_name" binding:"omitempty,max=64,min=1"`
	OrderDate   string `form:"order_date" binding:"omitempty,datetime=2006-01-02"`
	PageId      int    `form:"page_id" binding:"required,min=1"`
	PageSize    int    `form:"page_size" binding:"required,min=5,max=200"`
}

type PickingOrderID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type PickingOrderInfo struct {
	PickingOrder PickingOrder
	Items        []PickingOrderItem
	Details      []PickingOrderDetail
}

type FilterPickingOrderItem struct {
	POID int64  `json:"po_id"`
	SKU  string `json:"sku"`
}
type FilterPickingOrderDetail struct {
	POID         int64  `json:"po_id"`
	SKU          string `json:"sku"`
	LocationCode string `json:"location_code"`
}

type PickingOrderNew struct {
	SOID []string `json:"so_id" binding:"required,min=1"`
}
type PickingOrderDetailNew struct {
	POID           int64  `json:"picking_order_id"`
	ShelfLocation  string `json:"shelf_location"`
	ShelfCode      string `json:"shelf_code"`
	LocationLevel  string `json:"location_level"`
	LocationCode   string `json:"location_code"`
	ItemID         int64  `json:"item_id"`
	SKU            string `json:"sku"`
	ZohoItemID     string `json:"zoho_item_id"`
	Name           string `json:"name"`
	Quantity       int64  `json:"quantity"`
	QuantityPicked int64  `json:"quantity_picked"`
	TransactionID  int64  `json:"transaction_id"`
	UserName       string `json:"user_name"`
}

type PickingInfo struct {
	POID         int64  `json:"picking_order_id" binding:"required,min=1"`
	LocationCode string `json:"location_code" binding:"required,min=1"`
	Quantity     int64  `json:"quantity" binding:"required,min=1"`
}

type PickingTransactionNew struct {
	POID          int64  `db:"po_id" json:"po_id"`
	PONumber      string `db:"po_number" json:"po_number"`
	ItemName      string `db:"item_name" json:"item_name"`
	SKU           string `db:"sku" json:"sku"`
	Quantity      int64  `db:"quantity" json:"quantity"`
	ShelfCode     string `db:"shelf_code" json:"shelf_code"`
	ShelfLocation string `db:"shelf_location" json:"shelf_location"`
	LocationCode  string `db:"location_code" json:"location_code"`
	LocationLevel int64  `db:"location_level" json:"location_level"`
	UserName      string `json:"user_name"`
}

type PackingInfo struct {
	SOID     int64  `json:"sales_order_id" binding:"required,min=1"`
	Barcode  string `json:"barcode" binding:"required,min=1"`
	Quantity int64  `json:"quantity" binding:"required,min=1"`
}
type PackingTransactionNew struct {
	SOID     int64  `db:"so_id" json:"so_id"`
	SONumber string `db:"so_number" json:"so_number"`
	ItemName string `db:"item_name" json:"item_name"`
	SKU      string `db:"sku" json:"sku"`
	Quantity int64  `db:"quantity" json:"quantity"`
	UserName string `json:"user_name"`
}

type PickingResponse struct {
	TransactionID int64
	IsCompleted   bool
}
