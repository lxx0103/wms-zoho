package user

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"wms.com/core/queue"
	"wms.com/core/response"
	"wms.com/service"
)

// @Summary 用户列表
// @Id 47
// @Tags 用户管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param name query string false "用户名"
// @Param email query string false "邮箱"
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Success 200 object response.ListRes{data=UserProfile} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /users [GET]
func GetUserList(c *gin.Context) {
	var filter UserFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		response.ResponseError(c, "BindingError", err)
	}
	userService := NewUserService()
	count, list, err := userService.GetUserList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageId, filter.PageSize, count, list)
}

// @Summary 根据ID获取用户
// @Id 48
// @Tags 用户管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "用户ID"
// @Success 200 object response.SuccessRes{data=UserProfile} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /users/:id [GET]
func GetUserByID(c *gin.Context) {
	var uri UserUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	userService := NewUserService()
	user, err := userService.GetUserByID(uri.ID)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, user)
}

func NewUser(c *gin.Context) {
	var q UserProfile
	if err := c.ShouldBindJSON(&q); err != nil {
		response.ResponseError(c, "BindingError", err)
	}
	userService := NewUserService()
	user, err := userService.CreateNewUser(q)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
	}
	rabbit, _ := queue.GetConn()
	msg, _ := json.Marshal(user)
	err = rabbit.Publish("user", msg)
	if err != nil {
		fmt.Println(err)
	}
	response.Response(c, user)
}

// @Summary 根据ID更新用户
// @Id 49
// @Tags 用户管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "用户ID"
// @Param menu_info body UserUpdate true "用户信息"
// @Success 200 object response.SuccessRes{data=UserProfile} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /users/:id [PUT]
func UpdateUser(c *gin.Context) {
	var uri UserUri
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	var user UserUpdate
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user.User = claims.Username
	authService := NewUserService()
	new, err := authService.UpdateUser(uri.ID, user, claims.RoleID)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, new)
}
