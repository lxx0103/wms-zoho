package inventory

import (
	"errors"

	"github.com/gin-gonic/gin"
	"wms.com/api/v1/setting"
	"wms.com/core/response"
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
// @Tags 捡货模块
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param receive_info body ReceiveNew true "收货单信息"
// @Success 200 object response.SuccessRes{data=StockInRes} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /receives [POST]
func NewReceive(c *gin.Context) {
	var receive ReceiveNew
	if err := c.ShouldBindJSON(&receive); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	settingService := setting.NewSettingService()
	barcode, err := settingService.GetBarcodeByCode(receive.Barcode)
	if err != nil {
		response.ResponseError(c, "BarcodeError", err)
		return
	}
	toReceive := barcode.Quantity * receive.Quantity
	inventoryService := NewInventoryService()
	filter := FilterPOItem{
		POID: receive.POID,
		SKU:  barcode.SKU,
	}
	items, err := inventoryService.FilterPOItem(filter)
	if err != nil {
		response.ResponseError(c, "PurchaseOrderError", err)
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
	for j := 0; j < len(*locations); j++ {
		canStore += (*locations)[j].Capacity - (*locations)[j].Quantity
	}
	if canStore < toReceive {
		response.ResponseError(c, "LocationError", errors.New("NOT ENOUGH SPACE TO STORE THE ITEMS"))
		return
	}
	var res StockInRes
	for k := 0; k < len(*locations); k++ {
		var toStore StoctInLoc
		toStore.Code = (*locations)[k].Code
		toStore.ShelfID = (*locations)[k].ShelfID
		if (*locations)[k].Available >= toReceive {
			toStore.Quantity = toReceive
			res.Location = append(res.Location, toStore)
			break
		}
		toStore.Quantity = (*locations)[k].Available
		res.Location = append(res.Location, toStore)
	}
	response.Response(c, res)
}
