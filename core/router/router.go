package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"wms.com/core/config"
	_ "wms.com/docs"
	"wms.com/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func RunServer(r *gin.Engine) {
	host := config.ReadConfig("web.host")
	port := config.ReadConfig("web.port")

	r.Run(host + ":" + port)
}
func InitPublicRouter(r *gin.Engine, options ...func(*gin.RouterGroup)) {
	g := r.Group("")
	for _, opt := range options {
		opt(g)
	}

}

func InitAuthRouter(r *gin.Engine, options ...func(*gin.RouterGroup)) {
	g := r.Group("")
	g.Use(middleware.AuthorizeJWT())
	g.Use(middleware.RbacCheck())
	for _, opt := range options {
		opt(g)
	}
}
