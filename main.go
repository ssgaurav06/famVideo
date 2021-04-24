package main

import (
	"fam/client"
	"fam/config"
	"fam/router"
	"fam/storage"
	"fmt"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	fmt.Println("Connecting to database...")
	db := config.ConnectDB()
	fmt.Println("Database successfully connected")
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err)
	}
	config.MigrateDB(db)
	videoDataStorage := storage.NewVideoDataStorage(db)
	youtubeClient := client.NewYoutubeClient(videoDataStorage)
	go youtubeClient.GetClient()
	r := router.Router()
	http.ListenAndServe(":8080", r)
}
