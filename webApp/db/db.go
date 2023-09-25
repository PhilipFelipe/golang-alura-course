package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {
	connStr := "user=testuser dbname=testdb password=qwerty host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}
