package response

import "github.com/gin-gonic/gin"

func ResponseList(c *gin.Context, page int, page_size int, count int, data interface{}) {
	c.JSON(200, gin.H{
		"page_id":   page,
		"page_size": page_size,
		"count":     count,
		"data":      data,
	})
}

func Response(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"data": data,
	})
}

func ResponseError(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, gin.H{
		"message": err.Error(),
	})
}
