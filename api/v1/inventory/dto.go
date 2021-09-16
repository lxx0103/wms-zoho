package inventory

type ItemFilter struct {
	SKU      string `form:"sku" binding:"omitempty,max=64,min=1"`
	Name     string `form:"name" binding:"omitempty,max=64,min=1"`
	PageId   int    `form:"page_id" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,oneof=5 10 15 20"`
}

type ItemID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type PurchaseOrderFilter struct {
	PONumber    string `form:"po_number" binding:"omitempty,max=64,min=1"`
	VendorName  string `form:"vendor_name" binding:"omitempty,max=64,min=1"`
	ReceiveDate string `form:"receive_date" binding:"omitempty,datetime=2006-01-02"`
	PageId      int    `form:"page_id" binding:"required,min=1"`
	PageSize    int    `form:"page_size" binding:"required,oneof=5 10 15 20"`
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
	Location []StoctInLoc
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
	PageSize int    `form:"page_size" binding:"required,oneof=5 10 15 20"`
}
