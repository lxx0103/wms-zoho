package inventory

type ItemFilter struct {
	SKU      string `form:"sku" binding:"omitempty,max=64,min=1"`
	Name     string `form:"name" binding:"omitempty,max=64,min=1"`
	PageId   int    `form:"page_id" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,oneof=3 5 10"`
}

type ItemID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type PurchaseOrderFilter struct {
	PONumber   string `form:"po_number" binding:"omitempty,max=64,min=1"`
	VendorName string `form:"vendor_name" binding:"omitempty,max=64,min=1"`
	PageId     int    `form:"page_id" binding:"required,min=1"`
	PageSize   int    `form:"page_size" binding:"required,oneof=3 5 10"`
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
	ShelfID  int64  `json:"shelf_id"`
	Code     string `json:"code"`
	Quantity int64  `json:"quantity"`
}
