package controller

import (
	"project03/server"

	"github.com/gin-gonic/gin"
)

func FindAllHandler() gin.HandlerFunc {
	return server.FindAllHandler()
}
func FindOneHandler() gin.HandlerFunc {
	return server.FindOneHandler()
}
func InsertOneHandler() gin.HandlerFunc {
	return server.InsertOneHandler()
}
func UpdateOneHandler() gin.HandlerFunc {
	return server.UpdateOneHandler()
}
func DeleteOneHandler() gin.HandlerFunc {
	return server.DeleteOneHandler()
}
