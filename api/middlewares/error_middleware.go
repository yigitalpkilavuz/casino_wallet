package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call the next middleware/handler in the chain
		c.Next()

		// If there was an error during processing the request, it's available in c.Errors
		if len(c.Errors) > 0 {
			// Log the error and return a 500 response
			logrus.Errorf("Error processing request: %v", c.Errors)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}
}
