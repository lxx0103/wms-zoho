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

type NewPickingCreated struct {
	PickingID int64  `json:"picking_order_id"`
	User      string `json:"user"`
}

func Subscribe(conn *queue.Conn) {
	conn.StartConsumer("NewReceiveCreated", "NewReceiveCreated", AddTransaction)
	conn.StartConsumer("NewPickingCreated", "NewPickingCreated", AddPicking)
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

func AddPicking(d amqp.Delivery) bool {
	if d.Body == nil {
		return false
	}
	var newPickingCreated NewPickingCreated
	err := json.Unmarshal(d.Body, &newPickingCreated)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return false
	}
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	// settingRepo := setting.NewSettingRepository(db)
	var filter FilterPickingOrderItem
	filter.POID = newPickingCreated.PickingID
	pickingOrderItem, err := repo.FilterPickingOrderItem(filter)
	if err != nil {
		fmt.Println("11")
		fmt.Println(err)
		return false
	}
	for i := 0; i < len(*pickingOrderItem); i++ {
		for (*pickingOrderItem)[i].Quantity > 0 {
			transaction, err := repo.GetTransactionToPick((*pickingOrderItem)[i].SKU)
			if err != nil {
				fmt.Println("111")
				fmt.Println(err)
				return false
			}
			if (*transaction).Balance >= (*pickingOrderItem)[i].Quantity {
				var detail PickingOrderDetailNew
				detail.POID = (*pickingOrderItem)[i].POID
				detail.ItemID = (*pickingOrderItem)[i].ItemID
				detail.SKU = (*pickingOrderItem)[i].SKU
				detail.ZohoItemID = (*pickingOrderItem)[i].ZohoItemID
				detail.Name = (*pickingOrderItem)[i].Name
				detail.ShelfLocation = (*transaction).ShelfLocation
				detail.ShelfCode = (*transaction).ShelfCode
				detail.LocationLevel = (*transaction).LocationLevel
				detail.LocationCode = (*transaction).LocationCode
				detail.Quantity = (*pickingOrderItem)[i].Quantity
				detail.QuantityPicked = 0
				detail.UserName = newPickingCreated.User
				detail.TransactionID = (*transaction).ID
				_, err := repo.CreatePickingOrderDetail(detail)
				if err != nil {
					fmt.Println("112")
					fmt.Println(err)
					return false
				}
				(*pickingOrderItem)[i].Quantity = 0
			} else {
				var detail PickingOrderDetailNew
				detail.POID = (*pickingOrderItem)[i].POID
				detail.ItemID = (*pickingOrderItem)[i].ItemID
				detail.SKU = (*pickingOrderItem)[i].SKU
				detail.ZohoItemID = (*pickingOrderItem)[i].ZohoItemID
				detail.Name = (*pickingOrderItem)[i].Name
				detail.ShelfLocation = (*transaction).ShelfLocation
				detail.ShelfCode = (*transaction).ShelfCode
				detail.LocationLevel = (*transaction).LocationLevel
				detail.LocationCode = (*transaction).LocationCode
				detail.Quantity = (*transaction).Balance
				detail.QuantityPicked = 0
				detail.UserName = newPickingCreated.User
				detail.TransactionID = (*transaction).ID
				_, err := repo.CreatePickingOrderDetail(detail)
				if err != nil {
					fmt.Println("113")
					fmt.Println(err)
					return false
				}
				(*pickingOrderItem)[i].Quantity -= (*transaction).Balance
			}
		}
	}
	return true
}
