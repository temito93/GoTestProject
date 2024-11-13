package utils

import (
	"github.com/gin-gonic/gin"
)

// RespondJSON sends a JSON response with a status code
func RespondJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

// RespondError sends an error response with a status code
func RespondError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}
