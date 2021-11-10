package user

type UserUri struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// UserFilter represents a filter passed to FindUsers().
type UserFilter struct {
	Name     string `form:"name"`
	Email    string `form:"email" binding:"omitempty,email"`
	PageId   int    `form:"page_id" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=5,max=200"`
}

// UserUpdate represents a set of fields to be updated via UpdateUser().
type UserUpdate struct {
	Name   string `json:"name" binding:"required,min=2"`
	Email  string `json:"email" binding:"required,email"`
	RoleID int64  `json:"role_id" binding:"required,min=1"`
	User   string `json:"user" swaggerignore:"true"`
}
