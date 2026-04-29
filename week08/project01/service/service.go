package service

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"project01/global"
	"project01/model"

	"strconv"
	_ "time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rabbitmq/amqp091-go"
)

func Firsthandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Registerhandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u model.User
		var body string
		err := c.ShouldBind(&u)
		if err != nil {
			panic(err)
		}
		var id int
		err = global.Db.QueryRow("select user_id from email where user_id=?", u.UserId).Scan(&id)
		if err != nil {
			if err == sql.ErrNoRows {
				res, _ := global.Db.Exec("insert into email (user_id,password) values(?,?)", u.UserId, u.Password)
				i, _ := res.RowsAffected()
				fmt.Println(i)
				body = "success"
			} else {
				body = "fail"
			}
		} else {
			body = "fail"
		}
		cnn := Amqp()
		defer cnn.Close()
		ch, err := cnn.Channel()
		if err != nil {
			panic(err)
		}
		defer ch.Close()
		q, err := ch.QueueDeclare("hello", true, false, false, false, nil)
		if err != nil {
			panic(err)
		}
		ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
		c.Redirect(http.StatusSeeOther, "/api/hello")

	}
}
func Hellohandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cnn := Amqp()
		defer cnn.Close()
		ch, err := cnn.Channel()
		if err != nil {
			panic(err)
		}
		defer ch.Close()
		q, err := ch.QueueDeclare("hello", true, false, false, false, nil)
		if err != nil {
			panic(err)
		}
		msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
		if err != nil {
			panic(err)
		}
		b := <-msgs
		if string(b.Body) == "success" {
			c.HTML(http.StatusOK, "hello.html", nil)
		} else {
			c.HTML(http.StatusOK, "fail.html", nil)
		}

	}
}

func Userhandler(c *gin.Context) {
	s := c.GetHeader("s")
	var u model.User
	token, err := jwt.ParseWithClaims(s, &u, func(t *jwt.Token) (any, error) {
		return global.MySigningKey, nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "fail",
			"data": err.Error(),
		})
		return
	}
	if token.Valid {
		var uinfos []model.UserInfo = make([]model.UserInfo, 0)
		rows, err := global.Db.Query("select name,age from students where id=?", u.UserId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "fail",
				"data": err.Error(),
			})
			return
		}
		defer rows.Close()
		for rows.Next() {
			var uinfo model.UserInfo
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
			"msg":  "fail",
			"data": "token过期了",
		})
		return
	}
}

func Userhash(c *gin.Context) {
	res, err := global.Rdb.HGetAll(global.Ctx, "UserHash").Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "fail",
			"data": err.Error(),
		})
		return
	}
	if len(res) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "fail",
			"data": "无数据",
		})
		return
	}
	num, _ := strconv.Atoi(res["userid"])
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": num,
	})
}
