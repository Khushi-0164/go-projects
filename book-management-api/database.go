package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func connectDB() {
	connStr := "user=khushii dbname=bookdb sslmode=disable"

	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL!")
}
