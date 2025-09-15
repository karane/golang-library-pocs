package utils

import "github.com/gin-gonic/gin"

// JSONError sends a structured JSON error response
func JSONError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
	})
}
