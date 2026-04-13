package router

import (
	"project01/controller"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	rou := r.Group("/admin")
	{
		rou.GET("/", controller.AdminController{}.Index)
		rou.GET("/news", controller.AdminController{}.News)
	}
}
