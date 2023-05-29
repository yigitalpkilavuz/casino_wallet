package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yigitalpkilavuz/casino_wallet/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Skip authentication for "authenticate" route
		if strings.Contains(ctx.FullPath(), "authenticate") {
			ctx.Next()
			return
		}
		// Fetch the Authorization header from the request
		authHeader := ctx.Request.Header.Get("Authorization")
		// If no authorization header is found, return an error
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			return
		}

		// Split the Authorization header to extract the Bearer token
		bearerToken := strings.Split(authHeader, " ")
		// If the bearer token does not consist of two parts, return an error
		if len(bearerToken) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Extract the token from the bearerToken array
		token := bearerToken[1]
		// Validate the token and extract claims
		claims, err := auth.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Set the username in the context for subsequent handlers to use
		ctx.Set("username", claims.Username)
		// Process the next handler
		ctx.Next()
	}
}
