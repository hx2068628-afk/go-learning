package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var logger = log.Default()

func RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Println(c.Request.URL.Path)
		logger.Println(c.Request.URL.Port())
		logger.Println(c.Request.Method)
	}
}

type S struct {
	Role int `json:"role"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var s S
		start := time.Now()
		c.ShouldBind(&s)
		fmt.Println("role:", s.Role)
		if s.Role != 0 {
			c.Abort()
		} else {
			t := time.Since(start)
			fmt.Println(t)
			c.Set("time", t)
			c.Next()
			t = time.Since(start)
			fmt.Println(t)
		}
	}
}

func main() {
	var r = gin.Default()
	r.LoadHTMLGlob("./template/*")
	api := r.Group("/api")
	api.Use(AuthMiddleware(), RequestMiddleware())
	{
		api.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "upload.html", nil)
		})
		api.POST("/upload", func(c *gin.Context) {
			f, err := c.FormFile("file")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			c.SaveUploadedFile(f, "./1.txt")
		})
		api.POST("/old", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "/api/new")
		})
		api.GET("/new", func(c *gin.Context) {
			t, ok := c.Get("time")
			if !ok {
				c.String(200, "time fail")
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"msg":  "new",
				"time": t,
			})
		})
	}
	r.Run(":8080")
}
