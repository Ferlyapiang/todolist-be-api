package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Ferlyapiang/todolist-be-api/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Init() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		panic("Can't load configuration!")
	}

	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_DB := os.Getenv("MYSQL_DBNAME")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	// MYSQL_HOST := "localhost"
	// MYSQL_DB := "todolist"
	// MYSQL_USER := "root"
	// MYSQL_PASSWORD := ""
	// MYSQL_PORT := "3306"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DB))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(630)
	db.SetMaxOpenConns(700)

	return db
}
