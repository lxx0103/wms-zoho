package inventory

import "time"

type Item struct {
	ID             int64     `db:"id" json:"id"`
	SKU            string    `db:"sku" json:"sku"`
	Name           string    `db:"name" json:"name"`
	ZohoItemID     string    `db:"zoho_item_id" json:"zoho_item_id"`
	Unit           string    `db:"unit" json:"unit"`
	Stock          int       `db:"stock" json:"stock"`
	StockAvailable string    `db:"stock_available" json:"stock_available"`
	Enabled        string    `db:"enabled" json:"enabled"`
	Created        time.Time `db:"created" json:"created"`
	CreatedBy      string    `db:"created_by" json:"created_by"`
	Updated        time.Time `db:"updated" json:"updated"`
	UpdatedBy      string    `db:"updated_by" json:"updated_by"`
}
