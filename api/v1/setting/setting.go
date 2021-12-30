package setting

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"wms.com/core/response"
	"wms.com/service"
)

// @Summary 货架列表
// @Id 3
// @Tags 货架管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param code query string false "货架编码"
// @Param location query string false "货架位置"
// @Success 200 object response.ListRes{data=[]Shelf} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /shelves [GET]
func GetShelfList(c *gin.Context) {
	var filter ShelfFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	settingService := NewSettingService()
	count, list, err := settingService.GetShelfList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary 新建货架
// @Id 4
// @Tags 货架管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param shelf_info body ShelfNew true "货架信息"
// @Success 200 object response.SuccessRes{data=Shelf} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /shelves [POST]
func NewShelf(c *gin.Context) {
	var shelf ShelfNew
	if err := c.ShouldBindJSON(&shelf); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	shelf.User = claims.Username
	settingService := NewSettingService()
	new, err := settingService.NewShelf(shelf)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, new)
}

// @Summary 根据ID获取货架
// @Id 5
// @Tags 货架管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "货架ID"
// @Success 200 object response.SuccessRes{data=Shelf} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /shelves/:id [GET]
func GetShelfByID(c *gin.Context) {
	var uri ShelfID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	settingService := NewSettingService()
	shelf, err := settingService.GetShelfByID(uri.ID)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, shelf)

}

// @Summary 根据ID更新货架
// @Id 6
// @Tags 货架管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "货架ID"
// @Param shelf_info body ShelfNew true "货架信息"
// @Success 200 object response.SuccessRes{data=Shelf} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /shelves/:id [PUT]
func UpdateShelf(c *gin.Context) {
	var uri ShelfID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	var shelf ShelfNew
	if err := c.ShouldBindJSON(&shelf); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	shelf.User = claims.Username
	settingService := NewSettingService()
	new, err := settingService.UpdateShelf(uri.ID, shelf)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, new)
}

// @Summary 货位列表
// @Id 7
// @Tags 货位管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param code query string false "货位编码"
// @Param level query string false "货位层"
// @Param shelf_id query string false "货架id"
// @Param sku query string false "SKU"
// @Param is_alert query bool false "是否预警"
// @Param is_active query bool false "是否启用"
// @Success 200 object response.ListRes{data=[]Location} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /locations [GET]
func GetLocationList(c *gin.Context) {
	var filter LocationFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	fmt.Println(filter)
	settingService := NewSettingService()
	count, list, err := settingService.GetLocationList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary 新建货位
// @Id 8
// @Tags 货位管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param location_info body LocationNew true "货位信息"
// @Success 200 object response.SuccessRes{data=Location} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /locations [POST]
func NewLocation(c *gin.Context) {
	var location LocationNew
	if err := c.ShouldBindJSON(&location); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	location.User = claims.Username
	settingService := NewSettingService()
	new, err := settingService.NewLocation(location)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, new)
}

// @Summary 根据ID获取货位
// @Id 9
// @Tags 货位管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "货位ID"
// @Success 200 object response.SuccessRes{data=Location} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /locations/:id [GET]
func GetLocationByID(c *gin.Context) {
	var uri LocationID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	settingService := NewSettingService()
	location, err := settingService.GetLocationByID(uri.ID)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, location)

}

// @Summary 根据ID更新货位
// @Id 10
// @Tags 货位管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "货位ID"
// @Param location_info body LocationNew true "货位信息"
// @Success 200 object response.SuccessRes{data=Location} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /locations/:id [PUT]
func UpdateLocation(c *gin.Context) {
	var uri LocationID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	var location LocationNew
	if err := c.ShouldBindJSON(&location); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	location.User = claims.Username
	settingService := NewSettingService()
	new, err := settingService.UpdateLocation(uri.ID, location)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, new)
}

// @Summary 条码列表
// @Id 11
// @Tags 条码管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param code query string false "条码编码"
// @Param sku query string false "条码位置"
// @Success 200 object response.ListRes{data=[]Barcode} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /barcodes [GET]
func GetBarcodeList(c *gin.Context) {
	var filter BarcodeFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	settingService := NewSettingService()
	count, list, err := settingService.GetBarcodeList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary 新建条码
// @Id 12
// @Tags 条码管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param barcode_info body BarcodeNew true "条码信息"
// @Success 200 object response.SuccessRes{data=Barcode} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /barcodes [POST]
func NewBarcode(c *gin.Context) {
	var barcode BarcodeNew
	if err := c.ShouldBindJSON(&barcode); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	barcode.User = claims.Username
	settingService := NewSettingService()
	new, err := settingService.NewBarcode(barcode)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, new)
}

// @Summary 根据ID获取条码
// @Id 13
// @Tags 条码管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "条码ID"
// @Success 200 object response.SuccessRes{data=Barcode} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /barcodes/:id [GET]
func GetBarcodeByID(c *gin.Context) {
	var uri BarcodeID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	settingService := NewSettingService()
	barcode, err := settingService.GetBarcodeByID(uri.ID)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, barcode)

}

// @Summary 根据ID更新条码
// @Id 14
// @Tags 条码管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "条码ID"
// @Param barcode_info body BarcodeNew true "条码信息"
// @Success 200 object response.SuccessRes{data=Barcode} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /barcodes/:id [PUT]
func UpdateBarcode(c *gin.Context) {
	var uri BarcodeID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	var barcode BarcodeNew
	if err := c.ShouldBindJSON(&barcode); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	barcode.User = claims.Username
	settingService := NewSettingService()
	new, err := settingService.UpdateBarcode(uri.ID, barcode)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, new)
}

// @Summary 库存转移
// @Id 28
// @Tags 货位管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param location_info body LocationStockTransfer true "库存转移信息"
// @Success 200 object response.SuccessRes{data=int64} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /transfers [POST]
func StockTransfer(c *gin.Context) {
	var info LocationStockTransfer
	if err := c.ShouldBindJSON(&info); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	settingService := NewSettingService()
	new, err := settingService.StockTransfer(info, claims.Username)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, new)
}

// @Summary 库存转移列表
// @Id 29
// @Tags 货位管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param From query string false "来源货位编码"
// @Param To query string false "目标货位编码"
// @Success 200 object response.ListRes{data=[]TransferTransaction} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /transfers [GET]
func GetTransferList(c *gin.Context) {
	var filter TransferFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	settingService := NewSettingService()
	count, list, err := settingService.GetTransferList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary 获取最早一个批次的货位
// @Id 29
// @Tags 货位管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "转入货位ID"
// @Success 200 object response.SuccessRes{data=string} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /transferfrom/:id [GET]
func GetNextTransactionLocation(c *gin.Context) {
	var filter TranferFromFilter
	err := c.ShouldBindUri(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	settingService := NewSettingService()
	res, err := settingService.GetNextTransactionLocation(filter)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			response.Response(c, "")
			return
		}
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, res)
}
