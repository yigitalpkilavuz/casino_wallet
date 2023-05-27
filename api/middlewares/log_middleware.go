package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logrus.Infof("Incoming request: %s %s", ctx.Request.Method, ctx.Request.URL)
		ctx.Next()
		statusCode := ctx.Writer.Status()
		logrus.Infof("Response: %s %s %d", ctx.Request.Method, ctx.Request.URL, statusCode)
	}
}
