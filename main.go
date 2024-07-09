package main

import (
    "os"
    "sniffle-/config"
    "sniffle-/router"
    "github.com/joho/godotenv"
    "log"

)

var err error   


func main() {

    // Load .env file
    err = godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    config.InitDatabase()
    config.InitSES()

    r := router.SetupRouter() 

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    r.Run(":" + port)
}


