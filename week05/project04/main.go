package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4")
	if err != nil {
		panic("连接错误")
	}
	// defer db.Close()
	// var stu Student
	//单条查询
	// var name string
	// err = db.QueryRow("select name from students where id = ?", 2).Scan(&name)
	// fmt.Printf("%#v\n", name)
	//多条查询
	// rows, _ := db.Query("select name from students ")
	// defer rows.Close()
	// for rows.Next() {
	// 	rows.Scan(&name)
	// 	fmt.Println(name)
	// }
	//增加
	// db.Exec("insert into students (name,age,id) values(?,?,?)", "王五", 19, 4)
	// db.QueryRow("select name from students where id = ?", 4).Scan(&name)
	// fmt.Printf("%#v\n", name)
	//删除
	db.Exec("delete from students where id=?", 4)
	//修改
	db.Exec("update students set name=? where id=?", "王五", 2)

}
