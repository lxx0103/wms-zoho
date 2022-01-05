package inventory

import "github.com/gin-gonic/gin"

func Routers(g *gin.RouterGroup) {
	g.GET("/items", GetItemList)
	g.GET("/items/:id", GetItemByID)
	g.GET("/purchaseorders", GetPurchaseOrderList)
	g.GET("/purchaseorders/:id", GetPurchaseOrderByID)
	g.POST("/receives", NewReceive)
	g.GET("/receives", GetReceiveList)
	g.GET("/salesorders", GetSalesOrderList)
	g.GET("/salesorders/:id", GetSalesOrderByID)
	g.GET("/pickingorders", GetPickingOrderList)
	g.GET("/pickingorders/:id", GetPickingOrderByID)
	g.POST("/pickingorders", NewPickingOrder)
	g.POST("/pickings", NewPicking)
	g.POST("/packings", NewPacking)
	g.POST("/cancelreceives/:id", CancelReceive)
	g.POST("/cancelpickings/:id", CancelPicking)
	g.POST("/cancelpackings/:id", CancelPacking)
	g.POST("/adjustments", NewAdjustment)
	g.GET("/adjustments", GetAdjustmentList)

	g.GET("/pallets", GetPalletList)
	g.GET("/pallets/:id", GetPalletByID)
	g.PUT("/pallets/:id", UpdatePallet)
	g.POST("/pallets", NewPallet)
}
