package database

import (
	"database/sql"

	// "github.com/ben-eh/CodingOrganizer/tag"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/codingOrganizer")

	if err != nil {
		panic(err.Error())
	}

	return db
}
