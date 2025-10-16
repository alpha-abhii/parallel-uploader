package uploads

import (
	"log"
	"net/http"
	"context"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uploader Uploader
}

func NewHandler(uploader Uploader) *Handler {
	return &Handler{
		uploader: uploader,
	}
}

func (h *Handler) HandleInitiateUpload(c *gin.Context) {
	var req InitiateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	ctx := context.Background()
	state, err := h.uploader.InitiateUpload(ctx, req)

	if err != nil {
		log.Printf("ERROR: failed to initiate upload: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to initiate upload"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uploadId": state.ID,
	})
}

func (h *Handler) HandleGetPresignedURL(c *gin.Context) {
	uploadID := c.Param("uploadId")

	var req PresignedURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request, missing partNumber"})
		return
	}

	ctx := context.Background()
	url, err := h.uploader.GetPresignedURL(ctx, uploadID, req.PartNumber)
	if err != nil {
		log.Printf("ERROR: failed to get presigned URL for upload %s: %v", uploadID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get presigned URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}

func (h *Handler) HandleCompleteUpload(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "complete upload not implemented"})
}