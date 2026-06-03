package routes

import (
	"hotel-booking/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	authHandler *handler.AuthHandler,
) {

	auth := r.Group("/auth")
	{
		auth.POST(
			"/register",
			authHandler.Register,
		)

		auth.POST(
			"/login",
			authHandler.Login,
		)
	}
}
