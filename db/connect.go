package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDb() *sql.DB {
	conn := "host=localhost user=alura password=alura dbname=aluraloja port=5432 sslmode=disable"

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
