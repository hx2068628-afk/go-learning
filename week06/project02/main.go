package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func main() {
	r := gin.Default()
	mySigningKey := []byte("nihao")
	r.POST("/", func(c *gin.Context) {
		var myclaim MyClaim
		err := c.ShouldBind(&myclaim)
		if err != nil {
			log.Fatal(err)
		}
		myclaim.NotBefore = jwt.NewNumericDate(time.Now().Add(-time.Second))
		myclaim.Issuer = "zhangsan"
		myclaim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(20 * time.Second))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, myclaim)
		s, err := token.SignedString(mySigningKey)
		if err != nil {
			fmt.Println(err)
		}
		c.String(http.StatusOK, "%v", s)

	})
	r.GET("/news", func(c *gin.Context) {
		s := c.GetHeader("s")
		var myclaim MyClaim
		token, err := jwt.ParseWithClaims(s, &myclaim, func(t *jwt.Token) (any, error) {
			return mySigningKey, nil
		})
		if err != nil {
			c.String(http.StatusUnauthorized, "%v", err)
		}
		if token.Valid {
			c.String(http.StatusOK, "%v", myclaim)
		} else {
			c.String(http.StatusUnauthorized, "token无效")
		}

	})
	r.Run()
	// myClaim := MyClaim{"张三", "123", jwt.RegisteredClaims{
	// 	NotBefore: jwt.NewNumericDate(time.Now().Add(-time.Second)),
	// 	ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Second)),
	// 	Issuer:    "张三",
	// }}
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
	// s, err := token.SignedString(mySigningKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(s)
	// var myClaim01 MyClaim
	// time.Sleep(6 * time.Second)
	// token, err = jwt.ParseWithClaims(s, &myClaim01, func(t *jwt.Token) (any, error) {
	// 	return mySigningKey, nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(token.Claims.(*MyClaim).Username)
	// fmt.Println(myClaim01)

}
