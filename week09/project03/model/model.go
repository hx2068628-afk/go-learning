package model

import (
	"fmt"
	"project03/global"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func FindAll() (todolist []Todo, err error) {
	err = global.DB.Find(&todolist).Error
	return
}
func FindOne(id string) (todo Todo, err error) {
	err = global.DB.First(&todo, id).Error
	return
}
func InsertOne(todo Todo) error {
	err := global.DB.Debug().Create(&todo).Error
	return err
}
func UpdateOne(todo Todo, c *gin.Context) error {
	err := global.DB.First(&todo, todo.ID).Error
	if err != nil {
		panic(err)
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		fmt.Println(err)
	}
	err = global.DB.Debug().Save(&todo).Error
	return err
}
func DeleteOne(id string) error {
	err := global.DB.Where("id=?", id).Delete(Todo{}).Error
	return err
}
