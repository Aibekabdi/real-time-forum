package models

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	jwt.StandardClaims
	UserToken
}

type UserToken struct {
	UserId uint
	Role   uint
}

type SigningInput struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

const UserCtx string = "models.UserCtx"
