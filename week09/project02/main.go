package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	Id   int     `gorm:"column:user_id;primaryKey"`
	Name *string `gorm:"default:'张三'"`
	Age  int
}

//	type UserInfo struct {
//		Id   int
//		Name string
//		Age  int
//	}
// type Student struct {
// 	Id   int `gorm:""`
// 	Name string
// 	Age  int
// }

func main() {
	dsn := "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	// db.AutoMigrate(&UserInfo{})
	// db.Debug().Create(&UserInfo{Age: 19})
	var u UserInfo
	// var u1 UserInfo
	db.AutoMigrate(&UserInfo{})
	// (u.Name) = new(string)
	// db.Debug().Create(&u)
	u.Id = 1
	name := "wangwu"
	u.Name = &name
	// u.Age = 18
	// db.Debug().Save(&u)
	// db.Debug().Model(u).Update("age", 2)
	// db.Debug().Model(UserInfo{}).Where("age>?", 18).Updates(map[string]interface{}{"age": 18, "name": "laoliu"})
	// db.Debug().Where("age is null").Delete(UserInfo{})
	// db.Debug().Where("age=?", 19).First(&u1)
	// fmt.Printf("%#v\n", *u1.Name)
	var usernums []UserInfo
	db.Debug().Find(&usernums)
	for _, v := range usernums {
		fmt.Printf("%#v\n", *v.Name)
	}

}
