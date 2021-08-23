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

func UpdateShelf(c *gin.Context) {
	var uri ShelfID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	fmt.Println(uri)
	var shelf ShelfNew
	if err := c.ShouldBindJSON(&shelf); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	fmt.Println(shelf)
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
