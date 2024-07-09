package router

import (
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		
		RegisterUserRoutes(api)
		RegisterPostRoutes(api)

	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r

}