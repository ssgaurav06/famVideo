package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "fam"
)

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", user, dbname, password, host, port)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}
