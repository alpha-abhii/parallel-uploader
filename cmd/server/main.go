package main

import (
	"log"
	"net/http"

	"github.com/alpha-abhii/parallel-uploader/internal/modules/uploads"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	uploadService := uploads.NewService()
	uploadHandler := uploads.NewHandler(uploadService)

	api := router.Group("/api/v1")

	uploads.RegisterUploadRoutes(api, uploadHandler)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	serverAddr := ":8080"
	log.Printf("Starting server on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}