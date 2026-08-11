package main

import (
	"log"
	"net/http"
	"time"

	"chat-api/config"
	"chat-api/internal/hub"
	"chat-api/internal/middleware"
	"chat-api/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
	h := hub.NewHub()
	go h.Run()

	router := routes.SetupRouter(h)

	router.GET("/test-token/:userId", func(c *gin.Context) {
		userID := c.Param("userId")
		claims := middleware.Claims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			},
		}
		var uid uint
		fmtSscan(userID, &uid)
		claims.UserID = uid
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signed, _ := token.SignedString(config.JWTSecret())
		c.JSON(http.StatusOK, gin.H{"token": signed})
	})
	port := config.Getenv("PORT", "8080")
	log.Printf("Server starting on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
func fmtSscan(s string, v *uint) {
	var n uint
	for _, c := range s {
		if c < '0' || c > '9' {
			return
		}
		n = n*10 + uint(c-'0')
	}
	*v = n
}
