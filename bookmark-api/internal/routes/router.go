package routes

import (
	"bookmark-api/internal/handlers"
	"bookmark-api/internal/repository"
	"bookmark-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	repo := repository.NewBookmarkRepository(db)
	svc := service.NewBookingService(repo)
	handler := handlers.NewBookingHandler(svc)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := router.Group("/api/bookmarks")
	{
		api.POST("", handler.Create)
		api.GET("", handler.List)
		api.DELETE("/:id", handler.Delete)
	}

	return router
}
