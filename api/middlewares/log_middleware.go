package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the incoming request
		logrus.Infof("Incoming request: %s %s", c.Request.Method, c.Request.URL)

		// Call the next middleware/handler in the chain
		c.Next()

		// After calling c.Next(), the request has been processed and you can log the response
		statusCode := c.Writer.Status()
		logrus.Infof("Response: %s %s %d", c.Request.Method, c.Request.URL, statusCode)
	}
}
