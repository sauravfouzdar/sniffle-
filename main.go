package main

import (
	"sniffle-/config"
	"sniffle-/controllers"
	"sniffle-/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
    config.InitDatabase()
    //defer config.DB.Close()

    r := gin.Default()

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    auth := r.Group("/auth")
    auth.Use(middlewares.AuthMiddleware())
    {
        auth.GET("/users", controllers.GetUsers)
        auth.POST("/posts", controllers.CreatePost)
        auth.GET("/posts", controllers.GetPosts)
    }

    r.Run(":8080")
}


