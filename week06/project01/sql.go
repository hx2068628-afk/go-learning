package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Sql() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@(localhost:3306)/test?charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
