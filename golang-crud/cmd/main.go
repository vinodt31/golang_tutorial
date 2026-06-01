package main

import (
	"golang-crud/config"
	"golang-crud/routes"
	"log"
	"net/http"
	"time"
)

func main() {

	config.ConnectDB()

	routes.SetupRoutes()

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server started on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
