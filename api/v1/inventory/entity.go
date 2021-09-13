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

type PurchaseOrder struct {
	ID              int64     `db:"id" json:"id"`
	ZohoPOID        string    `db:"zoho_po_id" json:"zoho_po_id"`
	PONumber        string    `db:"po_number" json:"po_number"`
	PODate          time.Time `db:"po_date" json:"po_date"`
	ExpectedDate    time.Time `db:"expected_delivery_date" json:"expected_delivery_date"`
	ReferenceNumber string    `db:"reference_number" json:"reference_number"`
	Status          string    `db:"status" json:"status"`
	VendorID        string    `db:"vendor_id" json:"vendor_id"`
	VendorName      string    `db:"vendor_name" json:"vendor_name"`
	Enabled         string    `db:"enabled" json:"enabled"`
	Created         time.Time `db:"created" json:"created"`
	CreatedBy       string    `db:"created_by" json:"created_by"`
	Updated         time.Time `db:"updated" json:"updated"`
	UpdatedBy       string    `db:"updated_by" json:"updated_by"`
}

type PurchaseOrderItem struct {
	ID               int64     `db:"id" json:"id"`
	POID             int64     `db:"po_id" json:"po_id"`
	ItemID           int64     `db:"item_id" json:"item_id"`
	SKU              string    `db:"sku" json:"sku"`
	ZohoItemID       string    `db:"zoho_item_id" json:"zoho_item_id"`
	Name             string    `db:"name" json:"name"`
	Quantity         int64     `db:"quantity" json:"quantity"`
	QuantityReceived int64     `db:"quantity_received" json:"quantity_received"`
	Enabled          string    `db:"enabled" json:"enabled"`
	Created          time.Time `db:"created" json:"created"`
	CreatedBy        string    `db:"created_by" json:"created_by"`
	Updated          time.Time `db:"updated" json:"updated"`
	UpdatedBy        string    `db:"updated_by" json:"updated_by"`
}
