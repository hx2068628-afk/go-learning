package main

import (
	"project02/global"
	"project02/middleware"
	"project02/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func main() {
	r := gin.New()
	// cfg := zap.NewProductionConfig()
	// cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// cfg.EncoderConfig.TimeKey = "ts"
	// logger, err := cfg.Build()
	// if err != nil {
	// 	panic(err.Error())
	// }
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err.Error())
	}
	r.Use(middleware.ZapLogger(logger))

	global.Global()
	router.Router(r)
	r.Run()

}
