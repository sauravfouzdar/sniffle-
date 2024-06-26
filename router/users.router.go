package router

import (
	"sniffle-/controllers"
	"sniffle-/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup) {

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
		userGroup.GET("/health", controllers.Health)
	}
	userGroup.Use(middlewares.AuthMiddleware())
	{
		userGroup.GET("/users", controllers.GetUsers)
	}
}
