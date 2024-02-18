package app

import (
	"database/sql"
	"time"
	"todo-app-be/helper"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/todo_app")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
