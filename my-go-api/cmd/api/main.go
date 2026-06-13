package main

import (
	"database/sql"
	"log"
	"time"

	"my-go-api/config"
	"my-go-api/internal/handlers"
	"my-go-api/internal/middleware"
	"my-go-api/internal/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// 1. Load Configurations
	cfg := config.LoadConfig()

	// 2. Initialize Database Connection pool
	db, err := sql.Open("postgres", cfg.DBConn)
	if err != nil {
		log.Fatalf("Critical Error parsing Database configuration: %v", err)
	}
	defer db.Close()

	// Database Connection Pooling Settings
	db.SetMaxOpenConns(25)                 // Max active connections allowed
	db.SetMaxIdleConns(25)                 // Max idle connections retained
	db.SetConnMaxLifetime(5 * time.Minute) // Connection recycling limit

	// Verify connection status before booting router
	if err := db.Ping(); err != nil {
		log.Fatalf("Critical Error: Database is unreachable: %v", err)
	}

	r := gin.Default()

	// 3. Initialize Repositories (Data Access Layer)
	userRepo := repository.NewUserRepository(db)
	addrRepo := repository.NewAddressRepository(db)

	// 4. Setup App Handlers & Inject Repositories (Presentation Layer)
	authHandler := handlers.NewAuthHandler(db, userRepo, cfg.JWTSecret)
	addrHandler := handlers.NewAddressHandler(addrRepo)

	// 5. Public Routes Group
	public := r.Group("/api/v1")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
	}

	// 6. Protected Routes Group (Secured via AuthMiddleware)
	protected := r.Group("/api/v1")

	// UPDATED: Added userRepo dependency injection to handle database session tracking validation checks
	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret, userRepo))
	{
		protected.POST("/refresh-token", authHandler.RefreshToken)
		protected.POST("/logout", authHandler.Logout)
		protected.GET("/profile", authHandler.GetUserProfile)

		// Address Management Endpoints
		protected.GET("/address", addrHandler.GetAddresses)
		protected.POST("/address", addrHandler.CreateAddress)
		protected.PUT("/address", addrHandler.UpdateAddress)
		protected.DELETE("/address", addrHandler.DeleteAddress)
	}

	// Start Engine Server
	log.Println("Server initialized and listening successfully on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
