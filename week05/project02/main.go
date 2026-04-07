package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name string `json:"name" form:"name"`
	Age  string `json:"age"  form:"age"`
}

func main() {
	fmt.Println(1)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"username": 123,
			"password": 567,
		})
	})
	r.POST("/user1", func(c *gin.Context) {
		var stu Student
		err := c.ShouldBind(&stu)
		if err == nil {
			fmt.Println(stu)
			c.JSON(http.StatusOK, stu)
		} else {
			c.JSON(http.StatusOK, err)
		}
	})
	r.POST("/user2", func(c *gin.Context) {
		var stu Student
		err := c.ShouldBind(&stu)
		if err == nil {
			fmt.Println(stu)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "success",
				"data":    stu,
			})
		} else {
			c.JSON(http.StatusOK, err)
		}
	})
	r.POST("/user3", func(c *gin.Context) {
		var stu Student
		err := c.ShouldBind(&stu)
		if err == nil {
			fmt.Println(stu)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "success",
				"data":    stu,
			})
		} else {
			c.JSON(http.StatusOK, err)
		}
	})

	r.Run()

}
