package router

import (
	"github.com/fahmiarya/Article-Backend/internal/handler"
	"github.com/fahmiarya/Article-Backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(postHandler *handler.PostHandler) *gin.Engine {
	r := gin.Default()

	// Apply CORS middleware
	r.Use(middleware.CORSMiddleware())

	// Article routes
	articleGroup := r.Group("/article")
	{
		// 1. POST /article/ and /article -> Create article
		articleGroup.POST("/", postHandler.CreatePost)
		articleGroup.POST("", postHandler.CreatePost)

		// 2. GET /article/:limit/:offset -> List articles with pagination
		// Using :id/:offset avoids Gin's radix-tree wildcard parameter conflict with /:id
		articleGroup.GET("/:id/:offset", postHandler.GetPosts)

		// 3. GET /article/:id -> Get article by ID
		articleGroup.GET("/:id", postHandler.GetPostByID)

		// 4. PUT /article/:id -> Update article by ID
		articleGroup.PUT("/:id", postHandler.UpdatePost)

		// 5. DELETE /article/:id -> Delete article by ID
		articleGroup.DELETE("/:id", postHandler.DeletePost)
	}

	return r
}
