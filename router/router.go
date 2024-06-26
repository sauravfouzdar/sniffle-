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

	return r

}