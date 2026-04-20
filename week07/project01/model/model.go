package model

import "github.com/golang-jwt/jwt/v5"

type UserInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type User struct {
	Id       int    `json:"id"`
	UserId   int    `json:"userid" binding:"max=99999"`
	Password string `json:"password" binding:"required"`
	jwt.RegisteredClaims
}
