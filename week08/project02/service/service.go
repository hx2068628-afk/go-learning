package service

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"project02/global"
	"project02/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type Found struct {
	Account string `json:"account"`
	Role    int    `json:"role"`
	jwt.RegisteredClaims
}

func Zap() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("service模块的zap初始化失败")
	}
	return logger
}
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login", nil)
	}
}
func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := Zap()
		var user model.UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			logger.Info("数据绑定失败")
			return
		}

		var f Found
		err = global.Db.QueryRow("select account,role from userinfo where account=? and password=?", user.Account, user.Password).Scan(&f.Account, &f.Role)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  fmt.Sprintf("%s\t用户未注册\n", user.Account),
					"data": nil,
				})
				return
			} else {
				logger.Info("其他错误")
				return
			}
		}
		user.Role = f.Role
		f.NotBefore = jwt.NewNumericDate(time.Now().Add(-1 * time.Second))
		f.Issuer = "userinfo"
		f.ExpiresAt = jwt.NewNumericDate(time.Now().Add(30 * time.Second))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, f)
		auth, err := token.SignedString(global.Secret)
		if err != nil {
			logger.Info("token解析失败")
			return
		}
		fmt.Printf("role:%v\n", user.Role)
		c.HTML(http.StatusOK, "index", gin.H{
			"account": user.Account,
			"role":    user.Role,
			"auth":    auth,
		})
	}
}
func UserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("auth")
		var f Found
		logger := Zap()
		token, err := jwt.ParseWithClaims(auth, &f, func(t *jwt.Token) (any, error) {
			return global.Secret, nil
		})
		if err != nil {
			logger.Info(fmt.Sprintf("%v", err.Error()))
			return
		}
		if !token.Valid {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "token失效",
				"data": nil,
			})
			return
		}
		var id int
		err = global.Db.QueryRow("select id from userinfo where account=? and role=?", f.Account, f.Role).Scan(&id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				logger.Info("未查到该用户")
				return
			} else {
				logger.Info("其他错误")
				return
			}
		}
		if f.Role == 0 {
			rows, err := global.Db.Query("select id,account,password,role from userinfo where role=1")
			if err != nil {
				logger.Info(fmt.Sprintf("%v", err.Error()))
				return
			}
			var users = make([]model.UserInfo, 0)
			defer rows.Close()
			for rows.Next() {
				var user model.UserInfo
				err = rows.Scan(&user.Id, &user.Account, &user.Password, &user.Role)
				if err != nil {
					logger.Info("数据绑定失败")
					return
				}
				users = append(users, user)
			}
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "success",
				"data": users,
			})
		}

	}
}
