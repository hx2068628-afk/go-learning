package router

import (
	"project03/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/Todo", controller.FindAllHandler())
		v1.GET("/Todo/:id", controller.FindOneHandler())
		v1.POST("/Todo", controller.InsertOneHandler())
		v1.PUT("/Todo/:id", controller.UpdateOneHandler())
		v1.DELETE("Todo/:id", controller.DeleteOneHandler())
	}
}
