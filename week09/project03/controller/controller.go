package controller

import (
	"project03/server"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/Todo", server.FindAllHandler())
		v1.GET("/Todo/:id", server.FindOneHandler())
		v1.POST("/Todo", server.InsertOneHandler())
		v1.PUT("/Todo/:id", server.UpdateOneHandler())
		v1.DELETE("Todo/:id", server.DeleteOneHandler())
	}
}
