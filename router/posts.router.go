package router



import (
	"sniffle-/controllers"
	"sniffle-/middlewares"

	"github.com/gin-gonic/gin"
)


func RegisterPostRoutes(r *gin.RouterGroup) {
	postGroup := r.Group("/users")
	postGroup.Use(middlewares.AuthMiddleware())
	{
		postGroup.POST("/posts", controllers.CreatePost)
		postGroup.GET("/posts", controllers.GetPosts)
		postGroup.GET("/posts/:id", controllers.GetPost)
		// api.PATCH("/posts/:id", controllers.UpdatePost)

	}
}