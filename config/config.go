package config

import (
	"database/sql"
	"fam/handler"
	"fam/service"
	"fam/storage"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), 5432)
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

func Init() handler.VideoDataHandler {
	videoDataRepository := storage.NewVideoDataStorage(ConnectDB())
	videoDataService := service.NewVideoDataService(videoDataRepository)
	videoDataHandler := handler.NewVideoDataHandler(videoDataService)
	return videoDataHandler
}
