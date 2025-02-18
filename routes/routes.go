package routes

import (
	"github.com/gin-gonic/gin"
	"test/controller"
)

// Setup untuk API routes
func SetupRoutes(r *gin.Engine) {

	// Melakukan grouping untuk routesnya ke dalam /api
	api := r.Group("/api")

	{
		// Melakukan grouping untuk routes posts menjadi /api/posts
		posts := api.Group("/posts")
		{
			// Menggunakan controller untuk setiap routes
			posts.GET("/", controller.GetPosts)
			posts.POST("/", controller.CreatePost)
			posts.GET("/:id", controller.GetPost)
			posts.PUT("/:id", controller.UpdatePost)
			posts.DELETE("/:id", controller.DeletePost)
		}
	}
}
