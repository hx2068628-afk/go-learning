package main

import (
	"fmt"
	_ "fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Email    string `form:"email" binding:"required,checkEmail"`
	PassWord string `form:"password" binding:"required"`
}
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	var response = Response{http.StatusOK, "success", data}
	c.JSON(http.StatusOK, response)
}
func Fail(c *gin.Context, data interface{}) {
	var response = Response{http.StatusBadRequest, "fail", data}
	c.JSON(http.StatusBadRequest, response)
}
func checkEmail(f validator.FieldLevel) bool {
	email := f.Field().String()
	return strings.HasSuffix(email, "@qq.com")
}
func MyRecover(c *gin.Context) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	defer fmt.Println(567)
	c.Next()
	defer fmt.Println(123)
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	// r.Use(MyRecover)
	r.LoadHTMLGlob("static/*")

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("checkEmail", checkEmail)
	}
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.POST("/login", MyRecover, func(c *gin.Context) {
		var user User
		err := c.ShouldBind(&user)

		if err != nil {
			Fail(c, err.Error())
			panic(err)
		}
		Success(c, user)
	})
	r.Run()
}
