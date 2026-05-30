package routes

import (
	"golang-crud/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine){

	router.POST("/users", controller.CreateUser)

	router.GET("/users", controller.GetUsers)

	router.GET("/users/:id", controller.GetUserByID)

	router.PUT("/users/:id", controller.UpdateUser)

	router.DELETE("/users/:id", controller.DeleteUser)
}