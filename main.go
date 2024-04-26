package main


import (
	"net/http"

	"log"
	"github.com/gin-gonic/gin"

	"github.com/sauravfouzdar/sniffle-/api/routes"
	"github.com/sauravfouzdar/sniffle-/db"
)

func main() {
	r := gin.Default()

	db.InitDB()

	routes.SetupUserRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

