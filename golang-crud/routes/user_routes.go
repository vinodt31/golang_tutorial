package routes

import (
	"golang-crud/controller"
	"net/http"
)

func SetupRoutes() {

	http.HandleFunc("POST /users", controller.CreateUser)
	http.HandleFunc("GET /users", controller.GetUsers)

	http.HandleFunc("GET /users/{id}", controller.GetUserByID)
	http.HandleFunc("PUT /users/{id}", controller.UpdateUser)
	http.HandleFunc("DELETE /users/{id}", controller.DeleteUser)
}
