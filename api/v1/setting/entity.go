package setting

import "time"

type Shelf struct {
	ID        int64     `db:"id" json:"id"`
	Code      string    `db:"code" json:"code"`
	Level     int       `db:"level" json:"level"`
	Location  string    `db:"location" json:"location"`
	Enabled   string    `db:"enabled" json:"enabled"`
	Created   time.Time `db:"created" json:"created"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	Updated   time.Time `db:"updated" json:"updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}

type Location struct {
	ID        int64     `db:"id" json:"id"`
	Code      string    `db:"code" json:"code"`
	Level     int64     `db:"level" json:"level"`
	ShelfID   int64     `db:"shelf_id" json:"shelf_id"`
	SKU       string    `db:"sku" json:"sku"`
	Capacity  int64     `db:"capacity" json:"capacity"`
	Quantity  int64     `db:"quantity" json:"quantity"`
	Available int64     `db:"available" json:"available"`
	CanPick   int64     `db:"can_pick" json:"can_pick"`
	Unit      string    `db:"unit" json:"unit"`
	Enabled   string    `db:"enabled" json:"enabled"`
	Created   time.Time `db:"created" json:"created"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	Updated   time.Time `db:"updated" json:"updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}

type Barcode struct {
	ID        int64     `db:"id" json:"id"`
	Code      string    `db:"code" json:"code"`
	SKU       string    `db:"sku" json:"sku"`
	Unit      string    `db:"unit" json:"unit"`
	Quantity  int64     `db:"quantity" json:"quantity"`
	Enabled   string    `db:"enabled" json:"enabled"`
	Created   time.Time `db:"created" json:"created"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	Updated   time.Time `db:"updated" json:"updated"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}
