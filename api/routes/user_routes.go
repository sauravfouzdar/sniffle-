package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sauravfouzdar/sniffle-/api/handlers"
)


func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/users", handlers.CreateUser)
		userGroup.GET("/users", handlers.GetUsers)
		userGroup.GET("/users/:id", handlers.GetUser)
		userGroup.PUT("/users/:id", handlers.UpdateUser)
		userGroup.DELETE("/users/:id", handlers.DeleteUser)

		// Define other user routes here
	}
}
