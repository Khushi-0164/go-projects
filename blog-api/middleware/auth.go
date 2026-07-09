package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my secret key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get Authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing",
			})
			c.Abort()
			return
		}

		// Check Bearer token format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header",
			})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and validate JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// Ensure correct signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}

			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})
			c.Abort()
			return
		}

		// Get user_id from claims
		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid user ID",
			})
			c.Abort()
			return
		}

		// Store user ID in Gin context
		c.Set("user_id", int(userID))

		// Continue to the next handler
		c.Next()
	}
}
