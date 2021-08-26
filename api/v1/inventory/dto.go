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
