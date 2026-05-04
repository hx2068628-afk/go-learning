package model

type UserInfo struct {
	Id       int    `json:"id"`
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
	Role     int    `json:"role"`
}
