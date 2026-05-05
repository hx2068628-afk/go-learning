package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/a")
	})
	r.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusOK, "a")
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
		c.JSON(http.StatusOK, "c")
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bd")
	})
	r.Run()
}
