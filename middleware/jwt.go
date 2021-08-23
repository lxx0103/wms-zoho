package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"wms.com/core/response"
	"wms.com/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= len(BEARER_SCHEMA) {
			response.ResponseUnauthorized(c, "AuthError", errors.New("NO AUTH HEADER"))
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		if tokenString == "" {
			response.ResponseUnauthorized(c, "AuthError", errors.New("JWT AUTH ERROR"))
			return
		}
		claims, err := service.JWTAuthService().ParseToken(tokenString)
		if err != nil {
			response.ResponseUnauthorized(c, "AuthError", errors.New("JWT AUTH ERROR"))
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
