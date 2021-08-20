package response

import "github.com/gin-gonic/gin"

type ErrorRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

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
	var res ErrorRes
	res.Code = 1
	res.Message = err.Error()
	c.AbortWithStatusJSON(code, res)
}
