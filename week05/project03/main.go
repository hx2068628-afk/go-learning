package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/test?charset=utf8mb4"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	//创建数据表
	db.AutoMigrate(&Student{})
	//创建数据行
	// db.Create(&Student{3, "张三", 18})
	//查询第一条数据
	var stu Student
	db.First(&stu)
	fmt.Printf("%#v\n", stu)
	//修改数据
	db.Model(&stu).Update("name", "老六")
	fmt.Printf("%#v\n", stu)
	//删除数据
	// db.Delete(&stu)

}
