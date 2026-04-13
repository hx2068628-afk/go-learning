package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	c.String(http.StatusOK, "前台首页--")
}

func (con DefaultController) News(c *gin.Context) {
	c.String(http.StatusOK, "新闻首页--")
}
