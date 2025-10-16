package uploads

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleInitiateUpload(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "initiate upload not implemented"})
}

func HandleGetPresignedURL(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "get presigned url not implemented"})
}

func HandleCompleteUpload(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "complete upload not implemented"})
}