package controller

import (
	"project03/server"

	"github.com/gin-gonic/gin"
)

func FindAllHandler() gin.HandlerFunc {
	return server.FindAllHandler()
}
func FindPageHandler() gin.HandlerFunc {
	return server.FindPageHandler()
}
func FindOneHandler() gin.HandlerFunc {
	return server.FindOneHandler()
}
func FindNameOneHandler() gin.HandlerFunc {
	return server.FindNameOneHandler()
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
