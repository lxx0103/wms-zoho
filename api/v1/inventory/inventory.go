package inventory

import (
	"github.com/gin-gonic/gin"
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
