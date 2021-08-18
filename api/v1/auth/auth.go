package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"vandacare.com/core/queue"
	"vandacare.com/core/response"
	"vandacare.com/service"
)

func Signin(c *gin.Context) {
	var signinInfo SigninRequest
	err := c.ShouldBindJSON(&signinInfo)
	if err != nil {
		response.ResponseError(c, 400, err)
		return
	}
	authService := NewAuthService()
	jwtServices := service.JWTAuthService()
	authResult, err := authService.VerifyCredential(signinInfo)
	if err != nil {
		response.ResponseError(c, http.StatusUnauthorized, err)
		return
	}
	claims := service.CustomClaims{
		UserID:   authResult.ID,
		Username: authResult.Name,
		RoleID:   authResult.RoleID,
		Email:    authResult.Email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + 72000,
			Issuer:    "vandacare",
		},
	}
	generatedToken := jwtServices.GenerateToken(claims)
	var res SigninResponse
	res.Token = generatedToken
	res.User = authResult
	response.Response(c, res)
}

func Signup(c *gin.Context) {
	var signupInfo SignupRequest
	err := c.ShouldBindJSON(&signupInfo)
	if err != nil {
		response.ResponseError(c, 400, err)
		return
	}
	authService := NewAuthService()
	authID, err := authService.CreateAuth(signupInfo)
	if err != nil {
		response.ResponseError(c, 400, err)
		return
	}
	var newEvent NewAuthCreated
	newEvent.AuthID = authID
	newEvent.AuthType = signupInfo.AuthType
	newEvent.Credential = signupInfo.Credential
	newEvent.Identifier = signupInfo.Identifier
	newEvent.Gender = signupInfo.Gender
	newEvent.Name = signupInfo.Name
	newEvent.Email = signupInfo.Email

	rabbit, _ := queue.GetConn()
	msg, _ := json.Marshal(newEvent)
	err = rabbit.Publish("NewAuthCreated", msg)
	if err != nil {
		response.ResponseError(c, 400, err)
		return
	}
	response.Response(c, authID)
}
