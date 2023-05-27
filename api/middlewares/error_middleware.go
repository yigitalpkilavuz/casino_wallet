package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) > 0 {
			statusCode := ctx.Writer.Status()
			fmt.Println("errormiddleware")
			fmt.Println("errormiddleware")
			fmt.Println("errormiddleware")
			fmt.Println(ctx.Errors)
			fmt.Println("errormiddleware")
			fmt.Println("errormiddleware")
			fmt.Println("errormiddleware")

			logger.Infof("Error processing request: %s %s %d", ctx.Request.URL, statusCode, ctx.Errors)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}
}
