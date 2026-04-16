package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
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
	c.Next()
}

func main() {
	db, err := sql.Open("mysql", "root:123456@(localhost:3306)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	var ctx = context.Background()
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
		rdb.HSet(ctx, "user", "email", user.Email, "password", user.PassWord)
		rdb.Expire(ctx, "user", time.Second*10)
		time.Sleep(1 * time.Second)
		val, err := rdb.HGetAll(ctx, "user").Result()
		if len(val) == 0 {
			fmt.Println(err)
			db.QueryRow("select email,password from user").Scan(&user.Email, &user.PassWord)
			rdb.HSet(ctx, "user", "email", user.Email, "password", user.PassWord)
			Success(c, user)
			return
		}
		Success(c, val)

	})
	r.Run()

}
