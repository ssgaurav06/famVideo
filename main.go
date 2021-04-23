package main

import (
	"fam/config"
	"fmt"
)

func main() {
	fmt.Println("Connecting to database...")
	db := config.ConnectDB()
	fmt.Println("Database successfully connected")
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
