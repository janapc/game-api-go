package database

import (
	"database/sql"
	"os"
)

var Database *sql.DB

func Connection() {
	var err error
	mysqlUrl := os.Getenv("MYSQL_URL")
	Database, err = sql.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err)
	}
}
