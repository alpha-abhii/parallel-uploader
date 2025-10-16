package uploads

import "github.com/gin-gonic/gin"

func RegisterUploadRoutes(router *gin.RouterGroup) {
	uploadRoutes := router.Group("/uploads")
	{
		uploadRoutes.POST("/initiate", HandleInitiateUpload)
		uploadRoutes.POST("/:uploadId/presigned-url", HandleGetPresignedURL)
		uploadRoutes.POST("/:uploadId/complete", HandleCompleteUpload)
	}
}