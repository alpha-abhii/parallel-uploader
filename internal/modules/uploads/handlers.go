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
	c.JSON(http.StatusNotImplemented, gin.H{"message": "get presigned url not implemented"})
}

func (h *Handler) HandleCompleteUpload(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "complete upload not implemented"})
}