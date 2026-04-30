package controller

import (
	"project02/service"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/login", service.LoginHandler())
		// user.GET("/hello")
	}
}
