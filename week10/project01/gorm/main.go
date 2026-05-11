package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	ID     int
	Name   string
	Status int
}

func main() {
	dsn := "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	tx := db.Begin()
	var todo = Todo{Name: "112", Status: 1}
	tx.Create(todo)
	fmt.Println(todo.ID)
	tx.Commit()
}
