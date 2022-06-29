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
	LocationLevel string `json:"location_level"`
	User          string `json:"user"`
}

type NewPickingOrderCreated struct {
	PickingID int64  `json:"picking_order_id"`
	User      string `json:"user"`
}

type PickingOrderPicked struct {
	PickingID int64  `json:"picking_order_id"`
	User      string `json:"user"`
}

func Subscribe(conn *queue.Conn) {
	conn.StartConsumer("NewReceiveCreated", "NewReceiveCreated", AddTransaction)
	conn.StartConsumer("NewPickingOrderCreated", "NewPickingOrderCreated", AddPicking)
	conn.StartConsumer("PickingOrderPicked", "PickingOrderPicked", UpdateSOPicked)
	conn.StartConsumer("SalesOrderUpdated", "SalesOrderUpdated", UpdateSalesOrder)
	conn.StartConsumer("PurchaseOrderUpdated", "PurchaseOrderUpdated", UpdatePurchaseOrder)
	conn.StartConsumer("ItemUpdated", "ItemUpdated", UpdateItem)
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
	transation.LocationLevel = newReceiveCreated.LocationLevel
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
	var newPickingCreated NewPickingOrderCreated
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
	pickingOrderItem, err := repo.FilterPickingOrderItem(filter) //获取所有SKU
	if err != nil {
		fmt.Println("11")
		fmt.Println(err)
		return false
	}
	for i := 0; i < len(*pickingOrderItem); i++ {
		for q := (*pickingOrderItem)[i].Quantity; q > 0; {
			fmt.Println(q)
			transaction, err := repo.GetTransactionToPick((*pickingOrderItem)[i].SKU)
			if err != nil {
				fmt.Println("111")
				fmt.Println(err)
				return false
			}
			if (*transaction).Balance >= q {
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
				detail.Quantity = q
				detail.QuantityPicked = 0
				detail.UserName = newPickingCreated.User
				detail.TransactionID = (*transaction).ID
				_, err := repo.CreatePickingOrderDetail(detail)
				if err != nil {
					fmt.Println("112")
					fmt.Println(err)
					return false
				}
				q = 0
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
				q -= (*transaction).Balance
			}
		}
	}
	return true
}

func UpdateSOPicked(d amqp.Delivery) bool {
	if d.Body == nil {
		return false
	}
	var pickingOrderPicked PickingOrderPicked
	err := json.Unmarshal(d.Body, &pickingOrderPicked)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return false
	}
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	pickingOrder, err := repo.GetPickingOrderByID(pickingOrderPicked.PickingID)
	if err != nil {
		fmt.Println("22")
		fmt.Println(err)
		return false
	}
	err = repo.UpdateSOStatus(pickingOrder.SalesOrders, pickingOrderPicked.User)
	if err != nil {
		fmt.Println("333")
		fmt.Println(err)
		return false
	}
	return true
}

func UpdateSalesOrder(d amqp.Delivery) bool {
	if d.Body == nil {
		return false
	}
	var salesorderUpdate SalesorderUpdate
	err := json.Unmarshal(d.Body, &salesorderUpdate)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return false
	}
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	so, err := repo.GetSalesOrderByZohoID(salesorderUpdate.SalesorderID)
	if err == nil {
		err = repo.UpdateSalesorder(salesorderUpdate)
		if err != nil {
			fmt.Println("22")
			fmt.Println(err)
			return false
		}
		err = repo.UpdateSalesorderItem(so.ID, salesorderUpdate.Items)
		if err != nil {
			fmt.Println("333")
			fmt.Println(err)
			return false
		}
	} else {
		err := repo.NewSalesorder(salesorderUpdate)
		if err != nil {
			fmt.Println("22")
			fmt.Println(err)
			return false
		}
	}
	return true
}

func UpdatePurchaseOrder(d amqp.Delivery) bool {
	if d.Body == nil {
		return false
	}
	var purchaseUpdate PurchaseorderUpdate
	err := json.Unmarshal(d.Body, &purchaseUpdate)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return false
	}
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	so, err := repo.GetPurchaseOrderByZohoID(purchaseUpdate.PurchaseorderID)
	if err == nil {
		err = repo.UpdatePurchaseorder(purchaseUpdate)
		if err != nil {
			fmt.Println("22")
			fmt.Println(err)
			return false
		}
		err = repo.UpdatePurchaseorderItem(so.ID, purchaseUpdate.Items)
		if err != nil {
			fmt.Println("333")
			fmt.Println(err)
			return false
		}
	} else {
		err := repo.NewPurchaseorder(purchaseUpdate)
		if err != nil {
			fmt.Println("22")
			fmt.Println(err)
			return false
		}
	}
	return true
}

type NewReceiveToZoho struct {
	POID     string `json:"po_id"`
	SKU      string `json:"sku"`
	Quantity int64  `json:"quantity"`
}

type NewPackedToZoho struct {
	SOID     string `json:"so_id"`
	SKU      string `json:"sku"`
	Quantity int64  `json:"quantity"`
}

func UpdateItem(d amqp.Delivery) bool {
	if d.Body == nil {
		return false
	}
	var itemUpdate ItemUpdate
	err := json.Unmarshal(d.Body, &itemUpdate)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return false
	}
	db := database.InitMySQL()
	repo := NewInventoryRepository(db)
	_, err = repo.GetItemByZohoID(itemUpdate.ItemID)
	if err == nil {
		err = repo.UpdateItem(itemUpdate)
		if err != nil {
			fmt.Println("22")
			fmt.Println(err)
			return false
		}
	} else {
		err := repo.NewItem(itemUpdate)
		if err != nil {
			fmt.Println("222")
			fmt.Println(err)
			return false
		}
	}
	return true
}
