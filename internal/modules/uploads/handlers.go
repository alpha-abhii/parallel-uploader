package uploads

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandleInitiateUpload(c *gin.Context) {
	var req InitiateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	uploadID, err := h.service.InitiateUpload(req)
	if err != nil {
		log.Printf("ERROR: failed to initiate upload: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to initiate upload"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uploadId": uploadID,
	})
}

func (h *Handler) HandleGetPresignedURL(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "get presigned url not implemented"})
}

func (h *Handler) HandleCompleteUpload(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "complete upload not implemented"})
}