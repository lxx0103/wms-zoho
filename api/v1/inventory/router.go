package inventory

import "github.com/gin-gonic/gin"

func Routers(g *gin.RouterGroup) {
	g.GET("/items", GetItemList)
	g.GET("/items/:id", GetItemByID)
}
