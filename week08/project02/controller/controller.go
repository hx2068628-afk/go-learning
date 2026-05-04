package controller

import (
	"project02/service"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/login", service.LoginHandler())
		user.POST("/index", service.IndexHandler())
		user.GET("/list", service.UserHandler())
	}
}
