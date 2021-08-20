package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"wms.com/core/response"
	"wms.com/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= len(BEARER_SCHEMA) {
			response.ResponseError(c, http.StatusUnauthorized, errors.New("NO AUTH HEADER"))
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		if tokenString == "" {
			response.ResponseError(c, http.StatusUnauthorized, errors.New("JWT AUTH ERROR"))
			return
		}
		claims, err := service.JWTAuthService().ParseToken(tokenString)
		if err != nil {
			response.ResponseError(c, http.StatusUnauthorized, errors.New("JWT AUTH ERROR"))
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
