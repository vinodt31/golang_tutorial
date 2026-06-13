package config

import (
	"log"
	"os"
)

type Config struct {
	JWTSecret []byte
	DBConn    string
}

// LoadConfig acts as our orchestrator/facade now
func LoadConfig() *Config {
	log.Println("Initializing application configuration mapping...")

	return &Config{
		JWTSecret: loadJWTConfig(),
		DBConn:    loadDBConfig(),
	}
}

// Separate function isolated purely for JWT cryptographic contexts
func loadJWTConfig() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Fallback for local development environment
		secret = "your_ultra_secure_secret_key_string_32_characters"
	}
	return []byte(secret)
}

// Separate function isolated purely for infrastructure Database network paths
func loadDBConfig() string {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Fallback for local development environment
		dbURL = "postgres://postgres:root@localhost:5432/my_go_api?sslmode=disable"
	}
	return dbURL
}
