package routes

import (
	"chat-api/internal/handlers"
	"chat-api/internal/hub"
	"chat-api/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *hub.Hub) *gin.Engine {
	router := gin.Default()

	chatHandler := handlers.NewChatHandler(h)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	ws := router.Group("/ws")
	ws.Use(middleware.AuthRequired())
	{
		ws.GET("/chat", chatHandler.ServeWS)
	}

	return router
}
