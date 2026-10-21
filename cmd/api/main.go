package main

import (
	"log"
	"os"

	"github.com/fahmiarya/Article-Backend/config"
	"github.com/fahmiarya/Article-Backend/internal/handler"
	"github.com/fahmiarya/Article-Backend/internal/repository"
	"github.com/fahmiarya/Article-Backend/internal/router"
	"github.com/fahmiarya/Article-Backend/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load, falling back to system environment variables")
	}

	// Initialize Database & AutoMigrate
	db := config.InitDB()

	// Initialize Layers
	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	// Setup Router
	r := router.SetupRouter(postHandler)

	// Start Server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is starting on port %s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
