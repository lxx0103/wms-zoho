package user

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"wms.com/core/queue"
	"wms.com/core/response"
)

func GetUserList(c *gin.Context) {
	var filter UserFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
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
