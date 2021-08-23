package setting

import "github.com/gin-gonic/gin"

func Routers(g *gin.RouterGroup) {
	g.GET("/shelves", GetShelfList)
	g.GET("/shelves/:id", GetShelfByID)
	g.PUT("/shelves/:id", UpdateShelf)
	g.POST("/shelves", NewShelf)
}
