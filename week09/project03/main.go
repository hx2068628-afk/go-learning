package main

import (
	"project03/global"
	"project03/model"
	"project03/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	global.Global()
	global.DB.AutoMigrate(model.Todo{})
	router.Router(r)
	r.Run()
}
