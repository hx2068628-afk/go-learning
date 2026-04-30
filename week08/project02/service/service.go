package service

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"project02/global"
	"project02/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Zap() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("service模块的zap初始化失败")
	}
	return logger
}
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := Zap()
		var user model.UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			logger.Info("数据绑定失败")
		}
		type found struct {
			account  string
			password string
		}
		var f found
		db, _, _ := global.Global()
		err = db.QueryRow("select account,password from userinfo where account=? and password=?", user.Account, user.Password).Scan(&f.account, &f.password)
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
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("%s\t用户,欢迎你\n", user.Account),
			"data": nil,
		})
	}
}
