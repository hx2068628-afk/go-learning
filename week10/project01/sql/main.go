package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}
	s, err := tx.Exec("insert into todos(name,status) values ('512',1)")
	id, err := s.LastInsertId()
	var name string
	tx.QueryRow("select name from todos where id =?", id).Scan(&name)
	fmt.Println(name)
	tx.QueryRow("select name from students where id =?", 2).Scan(&name)
	fmt.Println(name)
	tx.Rollback()
	tx.Commit()

}
