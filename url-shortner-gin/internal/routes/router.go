package routes

import (
	"net/http"
	"url-shortener/internal/handlers"
	"url-shortener/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	AuthHandler := handlers.NewAuthHandler(db)
	linkHandler := handlers.NewLinkHandler(db)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	auth := router.Group("/auth")
	{
		auth.POST("/signup", AuthHandler.Signup)
		auth.POST("/login", AuthHandler.Login)
	}

	router.GET("/r/:code", linkHandler.Redirect)

	api := router.Group("/api")
	api.Use(middleware.AuthRequired())
	{
		api.POST("/links", linkHandler.CreateLink)
		api.GET("/links", linkHandler.ListLinks)
		api.DELETE("/links/:id", linkHandler.DeleteLink)
	}
	return router
}
