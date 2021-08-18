package router

import (
	"github.com/gin-gonic/gin"
	"vandacare.com/core/config"
	"vandacare.com/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
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
	for _, opt := range options {
		opt(g)
	}
}
