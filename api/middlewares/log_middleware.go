package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Infof("Incoming request: %s %s", ctx.Request.Method, ctx.Request.URL)
		ctx.Next()
		statusCode := ctx.Writer.Status()
		logger.Infof("Response: %s %s %d", ctx.Request.Method, ctx.Request.URL, statusCode)
	}
}
