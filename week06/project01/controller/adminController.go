package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (con AdminController) Index(c *gin.Context) {
	c.String(http.StatusOK, "后台首页")
}

func (con AdminController) News(c *gin.Context) {
	c.String(http.StatusOK, "后台新闻首页")
}
