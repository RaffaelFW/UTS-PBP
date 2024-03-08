package controller

import (
	"database/sql"
	"log"
)

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_uts")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
