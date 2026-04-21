package controller

import (
	"project02/service"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup) {
	r.POST("/login", service.Loginhandler)
	r.GET("/user", service.Userhandler)
	r.GET("/userhash", service.Userhash)
}
