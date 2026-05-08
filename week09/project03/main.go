package main

import (
	"project03/controller"
	"project03/global"
	"project03/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	global.Global()
	global.DB.AutoMigrate(model.Todo{})
	controller.Controller(r)
	r.Run()
}
