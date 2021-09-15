package setting

type ShelfFilter struct {
	Code     string `form:"code" binding:"omitempty,max=64,min=1"`
	Location string `form:"location" binding:"omitempty,max=64,min=1"`
	PageId   int    `form:"page_id" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,oneof=3 5 10"`
}

type ShelfNew struct {
	Code     string `json:"code" binding:"required,min=1,max=64"`
	Location string `json:"location" binding:"required,min=1,max=64"`
	Enabled  int    `json:"enabled" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,min=1,max=8"`
	User     string `json:"user" swaggerignore:"true"`
}

type ShelfID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type LocationFilter struct {
	Code     string `form:"code" binding:"omitempty,max=64,min=1"`
	Level    int    `json:"level" binding:"omitempty,min=1,max=8"`
	ShelfID  int64  `form:"shelf_id" binding:"omitempty,min=1"`
	SKU      string `form:"sku" binding:"omitempty,max=64,min=1"`
	PageId   int    `form:"page_id" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,oneof=3 5 10"`
}

type LocationNew struct {
	Code     string `json:"code" binding:"required,min=1,max=64"`
	Level    int    `json:"level" binding:"required,min=1,max=8"`
	ShelfID  int64  `json:"shelf_id" binding:"required,min=1"`
	SKU      string `json:"sku" binding:"required,min=1,max=64"`
	Capacity int    `json:"capacity" binding:"required,min=1"`
	Quantity int    `json:"quantity" binding:"omitempty"`
	Unit     string `json:"unit" binding:"required,min=1,max=64"`
	Enabled  int    `json:"enabled" binding:"required,oneof=1 2"`
	User     string `json:"user" swaggerignore:"true"`
}

type LocationID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type BarcodeFilter struct {
	Code     string `form:"code" binding:"omitempty,max=64,min=1"`
	SKU      string `form:"sku" binding:"omitempty,max=64,min=1"`
	PageId   int    `form:"page_id" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,oneof=3 5 10"`
}

type BarcodeNew struct {
	Code     string `json:"code" binding:"required,min=1,max=64"`
	SKU      string `json:"sku" binding:"required,min=1,max=64"`
	Unit     string `json:"unit" binding:"required,min=1,max=64"`
	Quantity int    `json:"quantity" binding:"required,min=1"`
	Enabled  int    `json:"enabled" binding:"required,oneof=1 2"`
	User     string `json:"user" swaggerignore:"true"`
}

type BarcodeID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type UpdateLocationStock struct {
	Code     string `json:"code"`
	Quantity int64  `json:"quantity"`
	User     string `json:"user"`
}
