package config

import (
	// Standard library
	"fmt"
	"log"
	"os"

	// Third-party
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres" // Kept because we use lowercase postgres.Open below
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Removed the unnecessary semicolon (Go doesn't require them)
	err := godotenv.Load() 

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbName,
		port,
	)

	// FIX: Changed Postgres.Open to postgres.Open 
	// FIX: Changed &gorm.Config() to &gorm.Config{}
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	}

	fmt.Println("Database connected successfully")

	DB = database
}