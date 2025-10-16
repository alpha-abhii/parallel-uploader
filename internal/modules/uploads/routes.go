package uploads

import "github.com/gin-gonic/gin"

func RegisterUploadRoutes(router *gin.RouterGroup, h *Handler) {
	uploadRoutes := router.Group("/uploads")
	{
		uploadRoutes.POST("/initiate", h.HandleInitiateUpload)
		uploadRoutes.POST("/:uploadId/presigned-url", h.HandleGetPresignedURL)
		uploadRoutes.POST("/:uploadId/complete", h.HandleCompleteUpload)
	}
}