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
	User     string `json:"user"`
}

type ShelfID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
