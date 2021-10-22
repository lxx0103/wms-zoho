package setting

import "github.com/gin-gonic/gin"

func Routers(g *gin.RouterGroup) {
	g.GET("/shelves", GetShelfList)
	g.GET("/shelves/:id", GetShelfByID)
	g.PUT("/shelves/:id", UpdateShelf)
	g.POST("/shelves", NewShelf)

	g.GET("/locations", GetLocationList)
	g.GET("/locations/:id", GetLocationByID)
	g.PUT("/locations/:id", UpdateLocation)
	g.POST("/locations", NewLocation)

	g.GET("/barcodes", GetBarcodeList)
	g.GET("/barcodes/:id", GetBarcodeByID)
	g.PUT("/barcodes/:id", UpdateBarcode)
	g.POST("/barcodes", NewBarcode)

	g.POST("/transfers", StockTransfer)
	g.GET("/transfers", GetTransferList)
}
