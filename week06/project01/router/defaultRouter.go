package router

import (
	"project01/controller"

	"github.com/gin-gonic/gin"
)

func DefaultRouter(r *gin.Engine) {
	rou := r.Group("/")
	{
		rou.GET("/", controller.DefaultController{}.Index)
		rou.GET("/news", controller.DefaultController{}.News)
	}
}
