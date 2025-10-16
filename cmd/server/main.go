package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alpha-abhii/parallel-uploader/internal/modules/uploads"
	"github.com/alpha-abhii/parallel-uploader/internal/platform/database"
	"github.com/alpha-abhii/parallel-uploader/internal/platform/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Falling back to system environment variables.")
	}

	redisClient := database.NewRedisClient()
	s3Client := s3.NewS3Client()

	bucketName := os.Getenv("AWS_S3_BUCKET")
	if bucketName == "" {
		log.Fatal("AWS_S3_BUCKET environment variable not set. Please check your .env or environment.")
	}
	log.Printf("Using S3 Bucket: %s", bucketName)

	uploadStore := uploads.NewRedisStore(redisClient)
	uploadService := uploads.NewS3Uploader(s3Client, uploadStore, bucketName)
	uploadHandler := uploads.NewHandler(uploadService)

	router := gin.Default()
	api := router.Group("/api/v1")
	uploads.RegisterUploadRoutes(api, uploadHandler)
	
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverAddr := fmt.Sprintf(":%s", port)

	log.Printf("Starting HTTP server on %s...", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}