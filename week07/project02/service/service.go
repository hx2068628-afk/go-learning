package service

import (
	"net/http"
	"project02/global"
	"project02/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Loginhandler(c *gin.Context) {
	var u model.User
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "数据绑定错误",
			"data": err.Error(),
		})
		return
	}
	err = global.Db.QueryRow("select user_id from email where user_id=? and password = ?", u.UserId, u.Password).Scan(&u.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "账号或密码错误",
			"data": err.Error(),
		})
		return
	}
	global.Rdb.HSet(global.Ctx, "UserHash", "userid", u.UserId, "password", u.Password)
	u.NotBefore = jwt.NewNumericDate(time.Now().Add(-time.Second))
	u.Issuer = "userid"
	u.ExpiresAt = jwt.NewNumericDate(time.Now().Add(30 * time.Second))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	s, err := token.SignedString(global.MySigningKey)
	if err != nil {
		panic(err.Error())
	}
	// c.Header("test", "123")
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": s,
	})
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
