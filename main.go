package main

import (
	"fam/config"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	fmt.Println("Connecting to database...")
	db := config.ConnectDB()
	fmt.Println("Database successfully connected")
	if err := db.Ping(); err != nil {
		panic(err)
	}
	config.MigrateDB(db)
}
