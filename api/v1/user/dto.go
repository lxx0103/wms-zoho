package user

type UserUri struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// UserFilter represents a filter passed to FindUsers().
type UserFilter struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"omitempty,email"`
	PageId   int    `json:"page_id" binding:"required,min=1"`
	PageSize int    `json:"page_size" binding:"required,min=5,max=200"`
}

// UserUpdate represents a set of fields to be updated via UpdateUser().
type UserUpdate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
