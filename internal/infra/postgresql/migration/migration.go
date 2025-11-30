package main

import (
	"fmt"
	"log"
	"os"
	"sistem-buku/internal/domain/entity"
	"sistem-buku/internal/infra/postgresql"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := postgresql.New(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	migrator := db.Migrator()

	err = migrator.AutoMigrate(
		entity.Book{},
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
