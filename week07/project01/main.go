package main

import (
	"project01/global"
	"project01/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	global.Global()
	router.Router(r)
	r.Run()

}
