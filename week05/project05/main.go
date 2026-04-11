package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()
	db, err1 := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4")
	if err1 != nil {
		log.Fatal(err1)
	}
	// createtext := "create table if not exists user(id int,name varchar(10),age int);"
	// _, err2 := db.Exec(createtext)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// inserttext := "insert into user (id,name,age) values (1,'张三',18),(2,'李四',19),(2,'王五',20)"
	// _, err3 := db.Exec(inserttext)
	// if err3 != nil {
	// 	log.Fatal(err3)
	// }
	r.GET("/api/user", func(c *gin.Context) {

		Rows, _ := db.Query("select id,name,age from user")
		var usersilce = make([]UserInfo, 0)
		for Rows.Next() {
			var user UserInfo
			err := Rows.Scan(&user.Id, &user.Name, &user.Age)
			if err != nil {
				log.Fatal(err)
			}
			usersilce = append(usersilce, user)
		}
		defer Rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"data": usersilce,
		})
	})
	r.GET("/api/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user UserInfo
		db.QueryRow("select id,name,age from user where id=?", id).Scan(&user.Id, &user.Name, &user.Age)
		c.String(http.StatusOK, "查询结果为%v", user)
	})
	r.POST("/api/user", func(c *gin.Context) {
		var user UserInfo
		c.ShouldBind(&user)
		_, err := db.Exec("insert into user (id,name,age) values(?,?,?)", user.Id, user.Name, user.Age)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, user)
	})
	r.PUT("/api/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user UserInfo
		c.ShouldBind(&user)
		user.Id, _ = strconv.Atoi(id)
		_, err := db.Exec("update user set name=?,age=? where id=?", user.Name, user.Age, user.Id)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, user)
	})
	r.DELETE("/api/user/:id", func(c *gin.Context) {
		strId := c.Param("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("delete from user where id=?", id)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, 200)
	})
	r.Run()

}
