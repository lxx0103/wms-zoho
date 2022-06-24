package setting

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	"wms.com/core/database"
	"wms.com/core/queue"
)

type NewLocationCreated struct {
	SKU          string `json:"sku"`
	LocationCode string `json:"location_code"`
	Quantity     int    `json:"quantity"`
}
type TransactionNew struct {
	POID          int64  `json:"po_id"`
	PONumber      string `json:"po_number"`
	ItemName      string `json:"item_name"`
	SKU           string `json:"sku"`
	Quantity      int    `json:"quantity"`
	ShelfCode     string `json:"shelf_code"`
	ShelfLocation string `json:"shelf_location"`
	LocationCode  string `json:"location_code"`
	LocationLevel string `json:"location_level"`
	User          string `json:"user"`
}

func Subscribe(conn *queue.Conn) {
	conn.StartConsumer("NewLocationCreated", "NewLocationCreated", AddTransaction)
}

func AddTransaction(d amqp.Delivery) bool {
	if d.Body == nil {
		return false
	}
	var newLocationCreated NewLocationCreated
	err := json.Unmarshal(d.Body, &newLocationCreated)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return false
	}
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	location, err := repo.GetLocationByCode(newLocationCreated.LocationCode)
	if err != nil {
		fmt.Println("11")
		fmt.Println(err)
		return false
	}
	shelf, err := repo.GetShelfByID(location.ShelfID)
	if err != nil {
		fmt.Println("22")
		fmt.Println(err)
		return false
	}
	itemName, err := repo.GetItemNameBySKU(newLocationCreated.SKU)
	if err != nil {
		fmt.Println("33", itemName, "--", newLocationCreated.SKU)
		fmt.Println(err)
		return false
	}
	var transation TransactionNew
	transation.POID = 0
	transation.PONumber = "Initial Stock"
	transation.ItemName = itemName
	transation.SKU = newLocationCreated.SKU
	transation.Quantity = newLocationCreated.Quantity
	transation.ShelfCode = shelf.Code
	transation.ShelfLocation = shelf.Location
	transation.LocationCode = location.Code
	transation.LocationLevel = location.Level
	transation.User = "SYSTEM"
	err = repo.CreateTransaction(transation)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
