package main

import (
	"project02/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.Controller(r)
	r.Run()
}
