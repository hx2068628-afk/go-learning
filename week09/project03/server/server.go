package server

import (
	"fmt"
	"net/http"
	"project03/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAllHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		todolist, err := model.FindAll()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  err,
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "success",
			"data": todolist,
		})
	}
}

func FindOneHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "未找到id",
				"data": nil,
			})
		}
		todo, err := model.FindOne(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  err,
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "success",
			"data": todo,
		})
	}
}

func InsertOneHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo model.Todo
		err := c.ShouldBind(&todo)
		fmt.Println(todo)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  err,
				"data": nil,
			})
			return
		}
		err = model.InsertOne(todo)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  err,
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "新增成功",
			"data": nil,
		})
	}
}

func UpdateOneHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		var todo model.Todo
		var err error
		todo.ID, err = strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "id绑定失败",
				"data": nil,
			})
			return
		}
		err = model.UpdateOne(todo, c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  err,
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "更新成功",
			"data": nil,
		})

	}
}

// v1.DELETE("Todo/:id", DeleteOneHandler)
func DeleteOneHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := model.DeleteOne(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  err,
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "删除成功",
			"data": nil,
		})
	}
}
