package auth

import "github.com/gin-gonic/gin"

func Routers(g *gin.RouterGroup) {
	g.POST("/signin", Signin)
	g.POST("/signup", Signup)
}

// func AuthRouter(g *gin.RouterGroup) {
// 	g.POST("/signout", Signout)
// }
