package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:mysql@tcp(127.0.0.1)/user?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func queryRowDemo() {
	fmt.Printf("1 =====> 🚀🚀🚀 %v\n", 1)
	sqlStr := "select id, name, age from user where id >?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
		return
	}
	fmt.Printf("users =====> 🚀🚀🚀 %v\n", users)
}

func main() {
	if err := initDB(); err != nil {
		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
		return
	}
	fmt.Printf("success =====> 🚀🚀🚀")
	queryRowDemo()
}
