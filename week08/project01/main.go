package main

import (
	"project01/global"
	"project01/middleware"
	"project01/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func main() {
	r := gin.New()
	r.LoadHTMLGlob("./static/*")
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err.Error())
	}
	r.Use(middleware.ZapLogger(logger))

	global.Global()
	defer global.Db.Close()
	router.Router(r)
	r.Run()

}
