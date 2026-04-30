package model

type UserInfo struct {
	Id       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
