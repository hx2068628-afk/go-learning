package controller

import (
	"project01/service"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup) {
	r.GET("/", service.Firsthandler)
	r.POST("/register", service.Registerhandler())
	r.GET("/hello", service.Hellohandler())
	r.GET("/user", service.Userhandler)
	r.GET("/userhash", service.Userhash)
}
