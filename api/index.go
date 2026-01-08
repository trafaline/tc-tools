// package main

// import (
// 	"log"
// 	"os"
// 	"tc-tools/controllers"
// 	"tc-tools/middleware"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	r := gin.Default()

// 	r.Use(middleware.CORSMiddleware())

// 	api := r.Group("/api/v1")
// 	{
// 		api.Use(middleware.AuthMiddleware())

// 		api.POST("/url-preview", controllers.HandlePreview)
// 	}

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}
// 	r.Run(":" + port)
// }

package handler // Nama package harus handler untuk folder api/ di Vercel

import (
	"net/http"
	"tc-tools/controllers"
	"tc-tools/middleware"

	"github.com/gin-gonic/gin"
)

// Declare r sebagai variable global agar tidak di-init ulang terus (Optimization)
var r *gin.Engine

func init() {
	// Pindahkan konfigurasi Gin ke sini
	gin.SetMode(gin.ReleaseMode)
	r = gin.Default()

	// Pasang Middleware
	r.Use(middleware.CORSMiddleware())

	// Pasang Route
	api := r.Group("/api/v1")
	{
		api.Use(middleware.AuthMiddleware())
		api.POST("/url-preview", controllers.HandlePreview)
	}
}

// Handler adalah fungsi yang WAJIB ada dan di-Export (Huruf Kapital)
func Handler(w http.ResponseWriter, req *http.Request) {
	// Vercel akan memanggil fungsi ini, lalu kita teruskan ke Gin
	r.ServeHTTP(w, req)
}
