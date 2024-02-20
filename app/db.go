package app

import (
	"database/sql"
	"os"
	"time"
	"todo-app-be/helper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	DbName   string
	UserName string
}

func DBConnect() *sql.DB {
	var cfg Config

	err := godotenv.Load()
	helper.PanicIfError(err)

	cfg.Host = os.Getenv("DB_HOST")
	cfg.Port = os.Getenv("DB_PORT")
	cfg.DbName = os.Getenv("DB_NAME")
	cfg.UserName = os.Getenv("DB_USERNAME")

	connectionDb := cfg.UserName + ":" + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.DbName

	db, err := sql.Open("mysql", connectionDb)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
