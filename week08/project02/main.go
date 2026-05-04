package main

import (
	"project02/controller"
	"project02/global"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//解析模板
	r.LoadHTMLGlob("./template/*")
	controller.Controller(r)
	global.Global()
	r.Run()
}
