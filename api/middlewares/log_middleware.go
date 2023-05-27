package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Infof("Incoming request: %s %s", c.Request.Method, c.Request.URL)
		c.Next()
		statusCode := c.Writer.Status()
		logrus.Infof("Response: %s %s %d", c.Request.Method, c.Request.URL, statusCode)
	}
}
