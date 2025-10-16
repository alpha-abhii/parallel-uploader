package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/alpha-abhii/parallel-uploader/internal/modules/uploads"
)

func main() {
	router := gin.Default()

	api := router.Group("/api/v1")

	uploads.RegisterUploadRoutes(api)

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