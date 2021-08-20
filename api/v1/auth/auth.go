package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"wms.com/core/queue"
	"wms.com/core/response"
	"wms.com/service"
)

// @Summary 登录
// @Id 1
// @Tags 用户权限
// @Summary 用户登录获取token
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param login_info body SigninRequest true "登录类型"
// @Success 200 object SigninResponse 登录成功
// @Failure 400 object response.ErrorRes 内部错误
// @Failure 401 object response.ErrorRes 登录失败
// @Router /signin [POST]
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
