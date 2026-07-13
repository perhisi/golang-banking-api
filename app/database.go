package app

import (
	"database/sql"
	"golang-banking-api/helper"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "ardhian:afnan@tcp(localhost:3306)/golang_banking_api?parseTime=true"
	}
	db, err := sql.Open("mysql", dbURL)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
