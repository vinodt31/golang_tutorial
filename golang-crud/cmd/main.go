package main

import (
	"golang-crud/config"
	"golang-crud/model"
	"golang-crud/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	config.DB.AutoMigrate(&model.User{})

	router := gin.Default()

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
