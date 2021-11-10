package user

import "github.com/gin-gonic/gin"

func Routers(r *gin.RouterGroup) {
	r.GET("/users", GetUserList)
	r.GET("/users/:id", GetUserByID)
	r.POST("/users", NewUser)
	r.PUT("/users/:id", UpdateUser)
}
