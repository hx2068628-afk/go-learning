package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Man struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "你好12")
	})
	r.GET("/add", func(c *gin.Context) {
		man := Man{"1", 18}
		fmt.Printf("%+v", man)
		c.JSON(200, man)
	})
	//query
	r.GET("/query", func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")

		c.JSON(200, page)
	})
	//path
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})
	r.Run(":8000")
}
