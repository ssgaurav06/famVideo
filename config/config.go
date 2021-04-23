package config

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

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

func MigrateDB(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	migration, err := migrate.NewWithDatabaseInstance(
		"file://migrations", "fam", driver)
	if err != nil {
		panic(err)
	}
	fmt.Println("Applying database migrations!")
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
	fmt.Println("Ran all migrations")
}
