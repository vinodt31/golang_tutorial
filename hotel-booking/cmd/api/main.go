package main

import (
	"hotel-booking/internal/config"
	"hotel-booking/internal/handler"
	"hotel-booking/internal/repository"
	"hotel-booking/internal/routes"
	"hotel-booking/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// Database
	config.ConnectDB()

	// Repositories
	userRepo := repository.NewUserRepository(
		config.DB,
	)

	// Services
	authService := service.NewAuthService(
		userRepo,
	)

	// Handlers
	authHandler := handler.NewAuthHandler(
		authService,
	)

	// Router
	r := gin.Default()

	// Register Routes
	routes.RegisterRoutes(
		r,
		authHandler,
	)

	// Start Server
	r.Run(":8080")
}
