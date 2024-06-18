package main

import (
	// "sniffle-/config"
    "os"
	"sniffle-/controllers"
	"sniffle-/middlewares"
    "sniffle-/config"
	"github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"

)

var err error   


// func init() {
//     // Load .env file
//     err = godotenv.Load()
//     if err != nil {
//         log.Fatalf("Error loading .env file")
//     }
// }


func main() {

    // Load .env file
    err = godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    config.InitDatabase()
    r := gin.Default()

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.GET("/health", controllers.Health)

    auth := r.Group("/auth")
    auth.Use(middlewares.AuthMiddleware())
    {
        auth.GET("/users", controllers.GetUsers)
        auth.POST("/posts", controllers.CreatePost)
        auth.GET("/posts", controllers.GetPosts)
        auth.GET("/posts/:id", controllers.GetPost)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    r.Run(":" + port)
}


