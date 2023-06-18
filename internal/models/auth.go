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

const UserCtx string = "models.UserCtx"
