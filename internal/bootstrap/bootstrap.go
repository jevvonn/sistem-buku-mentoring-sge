package bootstrap

import (
	"fmt"
	"log"
	"os"
	"sistem-buku/internal/infra/postgresql"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	_, err = postgresql.New(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app.Listen(":3000")
}
