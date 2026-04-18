package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type UserInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type User struct {
	Id       int    `json:"id"`
	UserId   int    `json:"userid" binding:"max=99999"`
	Password string `json:"password" binding:"required"`
	jwt.RegisteredClaims
}

func main() {
	r := gin.Default()
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	mySigningKey := []byte("nihao")
	db, err1 := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4")
	if err1 != nil {
		log.Fatal(err1)
	}
	r.POST("/login", func(c *gin.Context) {
		var u User
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "数据绑定错误",
				"data": err.Error(),
			})
			return
		}
		err = db.QueryRow("select user_id from email where user_id=? and password = ?", u.UserId, u.Password).Scan(&u.UserId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "账号或密码错误",
				"data": err.Error(),
			})
			return
		}
		rdb.HSet(ctx, "UserHash", "userid", u.UserId, "password", u.Password)
		u.NotBefore = jwt.NewNumericDate(time.Now().Add(-time.Second))
		u.Issuer = "userid"
		u.ExpiresAt = jwt.NewNumericDate(time.Now().Add(30 * time.Second))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
		s, err := token.SignedString(mySigningKey)
		if err != nil {
			panic(err.Error())
		}
		// c.Header("test", "123")
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": s,
		})
	})
	r.GET("/user", func(c *gin.Context) {
		s := c.GetHeader("s")
		var u User
		token, err := jwt.ParseWithClaims(s, &u, func(t *jwt.Token) (any, error) {
			return mySigningKey, nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "",
				"data": err.Error(),
			})
			return
		}
		if token.Valid {
			var uinfos []UserInfo = make([]UserInfo, 0)
			rows, err := db.Query("select name,age from students where id=?", u.UserId)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusBadRequest,
					"msg":  "",
					"data": err.Error(),
				})
				return
			}
			defer rows.Close()
			for rows.Next() {
				var uinfo UserInfo
				rows.Scan(&uinfo.Name, &uinfo.Age)
				uinfos = append(uinfos, uinfo)
			}
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "success",
				"data": uinfos,
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "",
				"data": "token过期了",
			})
			return
		}
	})
	r.GET("/userhash", func(c *gin.Context) {
		res, err := rdb.HGetAll(ctx, "UserHash").Result()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "",
				"data": err.Error(),
			})
			return
		}
		num, _ := strconv.Atoi(res["userid"])
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "success",
			"data": num,
		})
	})

	r.Run()

}
