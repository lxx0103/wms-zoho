package auth

import "wms.com/api/v1/user"

type SigninRequest struct {
	AuthType   int    `json:"auth_type" binding:"required,max=9,min=1"`
	Identifier string `json:"identifier" binding:"required"`
	Credential string `json:"credential" binding:"required,min=6"`
}
type SigninResponse struct {
	Token string `json:"token"`
	User  user.UserProfile
}

type SignupRequest struct {
	AuthType   int    `json:"auth_type" binding:"required,max=9,min=1"`
	Identifier string `json:"identifier" binding:"required"`
	Credential string `json:"credential" binding:"required,min=6"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
}
