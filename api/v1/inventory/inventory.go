package inventory

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"wms.com/api/v1/setting"
	"wms.com/core/queue"
	"wms.com/core/response"
	"wms.com/service"
)

// @Summary 商品列表
// @Id 15
// @Tags 商品管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param sku query string false "商品SKU"
// @Param name query string false "商品名称"
// @Success 200 object response.ListRes{data=[]Item} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /items [GET]
func GetItemList(c *gin.Context) {
	var filter ItemFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	count, list, err := inventoryService.GetItemList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary 根据ID获取商品
// @Id 16
// @Tags 商品管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "商品ID"
// @Success 200 object response.SuccessRes{data=Item} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /items/:id [GET]
func GetItemByID(c *gin.Context) {
	var uri ItemID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	item, err := inventoryService.GetItemByID(uri.ID)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, item)

}

// @Summary 采购单列表
// @Id 17
// @Tags 采购单管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param po_number query string false "采购订单编码"
// @Param vendor_name query string false "供应商名称"
// @Param receive_date query string false "预计到货日期"
// @Success 200 object response.ListRes{data=[]PurchaseOrder} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /purchaseorders [GET]
func GetPurchaseOrderList(c *gin.Context) {
	var filter PurchaseOrderFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	count, list, err := inventoryService.GetPurchaseOrderList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary 根据ID获取采购单
// @Id 18
// @Tags 采购单管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "采购订单ID"
// @Success 200 object response.SuccessRes{data=PODetail} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /purchaseorders/:id [GET]
func GetPurchaseOrderByID(c *gin.Context) {
	var uri PurchaseOrderID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	po, err := inventoryService.GetPurchaseOrderByID(uri.ID)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	filter := FilterPOItem{
		POID: uri.ID,
		SKU:  "",
	}
	item, err := inventoryService.FilterPOItem(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	var res PODetail
	res.PO = *po
	res.Items = *item
	response.Response(c, res)

}

// @Summary 创建新收货单
// @Id 19
// @Tags 收货管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param receive_info body ReceiveNew true "收货单信息"
// @Success 200 object response.SuccessRes{data=StockInRes} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /receives [POST]
func NewReceive(c *gin.Context) {
	//Binding
	var receive ReceiveNew
	if err := c.ShouldBindJSON(&receive); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	//Get Barcode information
	settingService := setting.NewSettingService()
	barcode, err := settingService.GetBarcodeByCode(receive.Barcode)
	if err != nil {
		response.ResponseError(c, "BarcodeError", err)
		return
	}
	toReceive := barcode.Quantity * receive.Quantity
	toZoho := toReceive
	inventoryService := NewInventoryService()
	POInfo, err := inventoryService.GetPurchaseOrderByID(receive.POID)
	if err != nil {
		response.ResponseError(c, "PurchaseOrderError", err)
		return
	}
	filter := FilterPOItem{
		POID: receive.POID,
		SKU:  barcode.SKU,
	}
	items, err := inventoryService.FilterPOItem(filter)
	if err != nil {
		response.ResponseError(c, "PurchaseOrderError", err)
		return
	}
	item, err := inventoryService.GetItemBySKU(barcode.SKU)
	if err != nil {
		response.ResponseError(c, "ItemError", err)
		return
	}
	var canReceive, canStore int64
	canReceive = 0
	for i := 0; i < len(*items); i++ {
		canReceive += (*items)[i].Quantity - (*items)[i].QuantityReceived
	}
	if canReceive < toReceive {
		response.ResponseError(c, "QuantityError", errors.New("RECEIVE QUANTITY GREATER THAN PO"))
		return
	}
	locations, err := settingService.GetLocationBySKU(barcode.SKU)
	if err != nil {
		response.ResponseError(c, "LocationError", err)
		return
	}
	canStore = 0
	var res StockInRes
	isCompleted := false
	var gCapacity, elseCapacity, gQuantity, elseQuantity int64
	var gLocations, elseLocations []setting.Location
	gCapacity = 0
	elseCapacity = 0
	gQuantity = 0
	elseQuantity = 0

	for j := 0; j < len(*locations); j++ {
		if (*locations)[j].Level == "G" {
			gCapacity += (*locations)[j].Capacity - (*locations)[j].Quantity
			gQuantity += (*locations)[j].Quantity
			gLocations = append(gLocations, (*locations)[j])
		} else {
			elseCapacity += (*locations)[j].Capacity - (*locations)[j].Quantity
			elseQuantity += (*locations)[j].Quantity
			elseLocations = append(elseLocations, (*locations)[j])
		}
	}
	if elseQuantity == 0 { //如果其他层没有库存, 先G后其他
		canStore = elseCapacity + gCapacity
		if canStore < toReceive {
			response.ResponseError(c, "LocationError", errors.New("NOT ENOUGH SPACE TO STORE THE ITEMS"))
			return
		}
		for k := 0; k < len(gLocations); k++ {
			shelf, err := settingService.GetShelfByID(gLocations[k].ShelfID)
			if err != nil {
				response.ResponseError(c, "ShelfError", err)
				return
			}
			var toStore StoctInLoc
			toStore.LocationCode = gLocations[k].Code
			toStore.LocationLevel = gLocations[k].Level
			toStore.ShelfCode = shelf.Code
			toStore.ShelfLocation = shelf.Location
			toStore.ItemName = item.Name
			toStore.SKU = item.SKU
			if gLocations[k].Available >= toReceive {
				toStore.Quantity = toReceive
				toReceive = 0
				res.Location = append(res.Location, toStore)
				break
			}
			toStore.Quantity = gLocations[k].Available
			toReceive = toReceive - toStore.Quantity
			res.Location = append(res.Location, toStore)
		}
		if toReceive > 0 {
			for j := 0; j < len(elseLocations); j++ {
				shelf, err := settingService.GetShelfByID(elseLocations[j].ShelfID)
				if err != nil {
					response.ResponseError(c, "ShelfError", err)
					return
				}
				var toStore StoctInLoc
				toStore.LocationCode = elseLocations[j].Code
				toStore.LocationLevel = elseLocations[j].Level
				toStore.ShelfCode = shelf.Code
				toStore.ShelfLocation = shelf.Location
				toStore.ItemName = item.Name
				toStore.SKU = item.SKU
				if elseLocations[j].Available >= toReceive {
					toStore.Quantity = toReceive
					toReceive = 0
					res.Location = append(res.Location, toStore)
					break
				}
				toStore.Quantity = elseLocations[j].Available
				toReceive = toReceive - toStore.Quantity
				res.Location = append(res.Location, toStore)
			}
		}
	} else { //如果其他层有库存, 只能放其他
		canStore = elseCapacity
		if canStore < toReceive {
			response.ResponseError(c, "LocationError", errors.New("NOT ENOUGH SPACE TO STORE THE ITEMS"))
			return
		}
		for j := 0; j < len(elseLocations); j++ {
			shelf, err := settingService.GetShelfByID(elseLocations[j].ShelfID)
			if err != nil {
				response.ResponseError(c, "ShelfError", err)
				return
			}
			var toStore StoctInLoc
			toStore.LocationCode = elseLocations[j].Code
			toStore.LocationLevel = elseLocations[j].Level
			toStore.ShelfCode = shelf.Code
			toStore.ShelfLocation = shelf.Location
			toStore.ItemName = item.Name
			toStore.SKU = item.SKU
			if elseLocations[j].Available >= toReceive {
				toStore.Quantity = toReceive
				toReceive = 0
				res.Location = append(res.Location, toStore)
				break
			}
			toStore.Quantity = elseLocations[j].Available
			toReceive = toReceive - toStore.Quantity
			res.Location = append(res.Location, toStore)
		}
	}
	rabbit, _ := queue.GetConn()
	for l := 0; l < len(res.Location); l++ {
		//stock in
		var locationUpdateInfo setting.UpdateLocationStock
		locationUpdateInfo.Code = res.Location[l].LocationCode
		locationUpdateInfo.Quantity = res.Location[l].Quantity
		locationUpdateInfo.User = claims.Username
		_, err := settingService.UpdateLocationStock(locationUpdateInfo)
		if err != nil {
			response.ResponseError(c, "UpdateLocationError", err)
			return
		}
		//Update PO
		var poUpdateInfo POItemUpdate
		poUpdateInfo.POID = receive.POID
		poUpdateInfo.Quantity = res.Location[l].Quantity
		poUpdateInfo.SKU = res.Location[l].SKU
		poUpdateInfo.User = claims.Username
		isCompleted, err = inventoryService.UpdatePOItem(poUpdateInfo)
		if err != nil {
			response.ResponseError(c, "UpdateLocationError", err)
			return
		}
		var newEvent NewReceiveCreated
		newEvent.POID = receive.POID
		newEvent.SKU = barcode.SKU
		newEvent.ItemName = item.Name
		newEvent.Quantity = res.Location[l].Quantity
		newEvent.ShelfCode = res.Location[l].ShelfCode
		newEvent.ShelfLocation = res.Location[l].ShelfLocation
		newEvent.LocationCode = res.Location[l].LocationCode
		newEvent.LocationLevel = res.Location[l].LocationLevel
		newEvent.User = claims.Username
		msg, _ := json.Marshal(newEvent)
		err = rabbit.Publish("NewReceiveCreated", msg)
		if err != nil {
			response.ResponseError(c, "PublishError", err)
			return
		}
	}
	res.IsCompleted = isCompleted
	var eventToZoho NewReceiveToZoho
	eventToZoho.POID = POInfo.ZohoPOID
	eventToZoho.SKU = item.SKU
	eventToZoho.Quantity = toZoho
	msg, _ := json.Marshal(eventToZoho)
	err = rabbit.Publish("NewReceiveToZoho", msg)
	if err != nil {
		response.ResponseError(c, "SyncToZohoError", err)
		return
	}
	response.Response(c, res)
}

// @Summary 收货列表
// @Id 20
// @Tags 收货管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param po_id query int false "采购订单id"
// @Param po_number query string false "采购订单编码"
// @Param sku query string false "SKU"
// @Success 200 object response.ListRes{data=[]Transaction} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /receives [GET]
func GetReceiveList(c *gin.Context) {
	var filter ReceiveFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	count, list, err := inventoryService.GetReceiveList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary 销售订单列表
// @Id 21
// @Tags 销售订单管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param so_number query string false "销售订单编码"
// @Param customer_name query string false "顾客名称"
// @Param sales_name query string false "销售人员名称"
// @Param order_date query string false "销售订单日期"
// @Success 200 object response.ListRes{data=[]SalesOrder} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /salesorders [GET]
func GetSalesOrderList(c *gin.Context) {
	var filter SalesOrderFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	count, list, err := inventoryService.GetSalesOrderList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary 根据ID获取销售订单
// @Id 22
// @Tags 销售订单管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "销售订单ID"
// @Success 200 object response.SuccessRes{data=SODetail} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /salesorders/:id [GET]
func GetSalesOrderByID(c *gin.Context) {
	var uri SalesOrderID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	so, err := inventoryService.GetSalesOrderByID(uri.ID)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	filter := FilterSOItem{
		SOID: uri.ID,
		SKU:  "",
	}
	item, err := inventoryService.FilterSOItem(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	var res SODetail
	res.SO = *so
	res.Items = *item
	response.Response(c, res)

}

// @Summary 捡货订单列表
// @Id 23
// @Tags 捡货管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param order_number query string false "捡货单编码"
// @Param user_name query string false "用户名称"
// @Param order_date query string false "捡货单日期"
// @Success 200 object response.ListRes{data=[]PickingOrder} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /pickingorders [GET]
func GetPickingOrderList(c *gin.Context) {
	var filter PickingOrderFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	count, list, err := inventoryService.GetPickingOrderList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary 根据ID获取捡货订单
// @Id 24
// @Tags 捡货管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "销售订单ID"
// @Success 200 object response.SuccessRes{data=PickingOrderInfo} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /pickingorders/:id [GET]
func GetPickingOrderByID(c *gin.Context) {
	var uri PickingOrderID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	pickingOrder, err := inventoryService.GetPickingOrderByID(uri.ID)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	filter := FilterPickingOrderItem{
		POID: uri.ID,
		SKU:  "",
	}
	item, err := inventoryService.FilterPickingOrderItem(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	filter2 := FilterPickingOrderDetail{
		POID: uri.ID,
		SKU:  "",
	}
	detail, err := inventoryService.FilterPickingOrderDetail(filter2)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	var res PickingOrderInfo
	res.PickingOrder = *pickingOrder
	res.Items = *item
	res.Details = *detail
	response.Response(c, res)

}

// @Summary 创建新捡货单
// @Id 25
// @Tags 捡货管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param receive_info body PickingOrderNew true "销售订单信息"
// @Success 200 object response.SuccessRes{data=int64} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /pickingorders [POST]
func NewPickingOrder(c *gin.Context) {
	var salesOrders PickingOrderNew
	if err := c.ShouldBindJSON(&salesOrders); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	claims := c.MustGet("claims").(*service.CustomClaims)
	pickingID, err := inventoryService.CreatePickingOrder(salesOrders, claims.Username)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	rabbit, _ := queue.GetConn()
	var newEvent NewPickingOrderCreated
	newEvent.PickingID = pickingID
	newEvent.User = claims.Username
	msg, _ := json.Marshal(newEvent)
	err = rabbit.Publish("NewPickingOrderCreated", msg)
	if err != nil {
		response.ResponseError(c, "PublishError", err)
		return
	}
	response.Response(c, pickingID)
}

// @Summary 捡货
// @Id 26
// @Tags 捡货管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param receive_info body PickingInfo true "捡货信息"
// @Success 200 object response.SuccessRes{data=PickingResponse} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /pickings [POST]
func NewPicking(c *gin.Context) {
	var info PickingInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	settingService := setting.NewSettingService()
	claims := c.MustGet("claims").(*service.CustomClaims)
	// fmt.Println(info)
	location, err := settingService.GetLocationByCode(info.LocationCode)
	if err != nil {
		response.ResponseError(c, "LocationError", err)
		return
	}
	// fmt.Println(location)
	item, err := inventoryService.GetItemBySKU(location.SKU)
	if err != nil {
		response.ResponseError(c, "ItemError", err)
		return
	}
	if item.StockPicking < info.Quantity {
		response.ResponseError(c, "StockError", errors.New("ITEM PICK TOO MUCH"))
		return
	}

	filter := FilterPickingOrderDetail{
		POID:         info.POID,
		LocationCode: info.LocationCode,
	}
	pickingDetails, err := inventoryService.FilterPickingOrderDetail(filter)
	if err != nil {
		response.ResponseError(c, "PickingOrderItemError", err)
		return
	}
	fmt.Println(*pickingDetails)
	if (*pickingDetails)[0].Quantity-(*pickingDetails)[0].QuantityPicked < info.Quantity {
		response.ResponseError(c, "StockError", errors.New("LOCATION PICK TOO MUCH"))
		return
	}
	shelf, err := settingService.GetShelfByID(location.ShelfID)
	if err != nil {
		response.ResponseError(c, "ShelfError", err)
		return
	}
	var newTransaction PickingTransactionNew
	newTransaction.POID = info.POID
	newTransaction.ItemName = item.Name
	newTransaction.LocationCode = location.Code
	newTransaction.LocationLevel = location.Level
	newTransaction.ShelfCode = shelf.Code
	newTransaction.ShelfLocation = shelf.Location
	newTransaction.Quantity = info.Quantity
	newTransaction.SKU = item.SKU
	newTransaction.UserName = claims.Username
	transactionID, isFullPicked, err := inventoryService.CreatePickingTransaction(newTransaction)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	if isFullPicked {
		rabbit, _ := queue.GetConn()
		var newEvent PickingOrderPicked
		newEvent.PickingID = info.POID
		newEvent.User = claims.Username
		msg, _ := json.Marshal(newEvent)
		err = rabbit.Publish("PickingOrderPicked", msg)
		if err != nil {
			response.ResponseError(c, "PublishError", err)
			return
		}
	}
	var res PickingResponse
	res.IsCompleted = isFullPicked
	res.TransactionID = transactionID
	response.Response(c, res)
}

// @Summary 打包
// @Id 27
// @Tags 捡货管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param receive_info body PackingInfo true "捡货信息"
// @Success 200 object response.SuccessRes{data=int64} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /packings [POST]
func NewPacking(c *gin.Context) {
	var info PackingInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	settingService := setting.NewSettingService()
	claims := c.MustGet("claims").(*service.CustomClaims)
	// fmt.Println(location)
	barcode, err := settingService.GetBarcodeByCode(info.Barcode)
	if err != nil {
		response.ResponseError(c, "BarcodeError", err)
		return
	}
	info.Quantity = info.Quantity * barcode.Quantity
	item, err := inventoryService.GetItemBySKU(barcode.SKU)
	if err != nil {
		response.ResponseError(c, "ItemError", err)
		return
	}
	if item.StockPacking < info.Quantity {
		response.ResponseError(c, "StockError", errors.New("ITEM PACK TOO MUCH"))
		return
	}
	var filter FilterSOItem
	filter.SOID = info.SOID
	filter.SKU = barcode.SKU
	soItems, err := inventoryService.FilterSOItem(filter)
	if err != nil {
		response.ResponseError(c, "SalesOrder Item Error", err)
		return
	}
	fmt.Println((*soItems)[0])
	if info.Quantity > (*soItems)[0].Quantity-(*soItems)[0].QuantityPacked {
		response.ResponseError(c, "StockError", errors.New("SO ITEM PACK TOO MUCH"))
		return
	}
	var newTransaction PackingTransactionNew
	newTransaction.SOID = info.SOID
	newTransaction.ItemName = item.Name
	newTransaction.Quantity = info.Quantity
	newTransaction.SKU = item.SKU
	newTransaction.UserName = claims.Username
	transactionID, err := inventoryService.CreatePackingTransaction(newTransaction)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	soInfo, err := inventoryService.GetSalesOrderByID(info.SOID)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	rabbit, _ := queue.GetConn()
	var eventToZoho NewPackedToZoho
	eventToZoho.SOID = soInfo.ZohoSOID
	eventToZoho.SKU = item.SKU
	eventToZoho.Quantity = info.Quantity
	msg, _ := json.Marshal(eventToZoho)
	err = rabbit.Publish("NewPackedToZoho", msg)
	if err != nil {
		response.ResponseError(c, "SyncToZohoError", err)
		return
	}
	response.Response(c, transactionID)
}

// @Summary 取消收货
// @Id 51
// @Tags 收货管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "采购订单ID"
// @Success 200 object response.SuccessRes{data=int64} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /cancelreceives/:id [POST]
func CancelReceive(c *gin.Context) {
	var uri PurchaseOrderID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	claims := c.MustGet("claims").(*service.CustomClaims)
	res, err := inventoryService.CancelReceive(uri.ID, claims.Username)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, res)
}

// @Summary 取消捡货
// @Id 52
// @Tags 捡货管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "捡货订单ID"
// @Success 200 object response.SuccessRes{data=int64} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /cancelpickings/:id [POST]
func CancelPicking(c *gin.Context) {
	var uri PickingOrderID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	claims := c.MustGet("claims").(*service.CustomClaims)
	res, err := inventoryService.CancelPicking(uri.ID, claims.Username)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, res)
}

// @Summary 取消打包
// @Id 53
// @Tags 捡货管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "销售订单ID"
// @Success 200 object response.SuccessRes{data=int64} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /cancelpackings/:id [POST]
func CancelPacking(c *gin.Context) {
	var uri SalesOrderID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	claims := c.MustGet("claims").(*service.CustomClaims)
	res, err := inventoryService.CancelPacking(uri.ID, claims.Username)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, res)
}

// @Summary 库存调整
// @Id 54
// @Tags 库存管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param info body AdjustmentInfo true "调整信息"
// @Success 200 object response.SuccessRes{data=int64} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /adjustments [POST]
func NewAdjustment(c *gin.Context) {
	var info AdjustmentInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	claims := c.MustGet("claims").(*service.CustomClaims)
	info.UserName = claims.Username
	res, err := inventoryService.CreateAdjustment(info)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, res)
}

// @Summary Adjustment列表
// @Id 55
// @Tags 库存管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param sku query string false "SKU"
// @Param location_code query string false "Location Code"
// @Success 200 object response.ListRes{data=[]Adjustment} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /adjustments [GET]
func GetAdjustmentList(c *gin.Context) {
	var filter AdjustmentFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	count, list, err := inventoryService.GetAdjustmentList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary Pallet列表
// @Id 57
// @Tags Pallet管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param code query string false "Pallet编码"
// @Param level query string false "Pallet层"
// @Param so_id query string false "销售订单id"
// @Success 200 object response.ListRes{data=[]SalesOrderPallet} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /pallets [GET]
func GetPalletList(c *gin.Context) {
	var filter FilterSOPallet
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	count, list, err := inventoryService.GetPalletList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary 新建Pallet
// @Id 58
// @Tags Pallet管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param pallet_info body PalletNew true "Pallet信息"
// @Success 200 object response.SuccessRes{data=SalesOrderPallet} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /pallets [POST]
func NewPallet(c *gin.Context) {
	var pallet PalletNew
	if err := c.ShouldBindJSON(&pallet); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	pallet.UserName = claims.Username
	inventoryService := NewInventoryService()
	new, err := inventoryService.NewPallet(pallet)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, new)
}

// @Summary 根据ID获取Pallet
// @Id 59
// @Tags Pallet管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "PalletID"
// @Success 200 object response.SuccessRes{data=SalesOrderPallet} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /pallets/:id [GET]
func GetPalletByID(c *gin.Context) {
	var uri SOPalletID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	inventoryService := NewInventoryService()
	pallet, err := inventoryService.GetPalletByID(uri.ID)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, pallet)

}

// @Summary 根据ID更新Pallet
// @Id 60
// @Tags Pallet管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "PalletID"
// @Param pallet_info body PalletNew true "Pallet信息"
// @Success 200 object response.SuccessRes{data=SalesOrderPallet} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /pallets/:id [PUT]
func UpdatePallet(c *gin.Context) {
	var uri SOPalletID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	var pallet PalletNew
	if err := c.ShouldBindJSON(&pallet); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	pallet.UserName = claims.Username
	inventoryService := NewInventoryService()
	new, err := inventoryService.UpdatePallet(uri.ID, pallet)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, new)
}
