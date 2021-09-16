package inventory

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	"wms.com/core/database"
	"wms.com/core/queue"
)

type NewReceiveCreated struct {
	POID          int64  `json:"po_id"`
	SKU           string `json:"sku"`
	ItemName      string `json:"item_name"`
	ShelfID       int64  `json:"shelf_id"`
	ShelfCode     string `json:"shelf_code"`
	ShelfLocation string `json:"shelf_location"`
	Quantity      int64  `json:"quantity"`
	LocationID    int64  `json:"location_id"`
	LocationCode  string `json:"location_code"`
	LocationLevel int64  `json:"location_level"`
	User          string `json:"user"`
}

func Subscribe(conn *queue.Conn) {
	conn.StartConsumer("NewReceiveCreated", "NewReceiveCreated", AddTransaction)
}

func AddTransaction(d amqp.Delivery) bool {
	if d.Body == nil {
		return false
	}
	var newReceiveCreated NewReceiveCreated
	err := json.Unmarshal(d.Body, &newReceiveCreated)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return false
	}
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	po, err := repo.GetPurchaseOrderByID(newReceiveCreated.POID)
	if err != nil {
		fmt.Println("11")
		fmt.Println(err)
		return false
	}
	var transation TransactionNew
	transation.POID = newReceiveCreated.POID
	transation.PONumber = po.PONumber
	transation.ItemName = newReceiveCreated.ItemName
	transation.SKU = newReceiveCreated.SKU
	transation.Quantity = newReceiveCreated.Quantity
	transation.ShelfCode = newReceiveCreated.ShelfCode
	transation.ShelfLocation = newReceiveCreated.ShelfLocation
	transation.LocationCode = newReceiveCreated.LocationCode
	transation.User = newReceiveCreated.User
	err = repo.CreateTransaction(transation)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
