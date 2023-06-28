package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)


func Connect() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
