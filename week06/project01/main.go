package main

import (
	"project01/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Sql()
	router.DefaultRouter(r)
	router.AdminRouter(r)
	r.Run()
}
