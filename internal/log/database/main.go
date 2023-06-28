package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)


func Connect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		print("Error connecting to database")
		panic(err)
	}
	return db
}
