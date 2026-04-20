package router

import (
	"project01/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	user := r.Group("/api")
	controller.Controller(user)
}
