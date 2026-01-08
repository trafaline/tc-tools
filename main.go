package main

import (
	"log"
	"os"
	"tc-tools/controllers"
	"tc-tools/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.StaticFile("/", "./index.html")

	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api/v1")
	{
		api.Use(middleware.AuthMiddleware())

		api.POST("/url-preview", controllers.HandlePreview)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
