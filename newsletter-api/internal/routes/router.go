package routes

import (
	"net/http"
	"newsletter-api/internal/handlers"
	"newsletter-api/internal/repository"
	"newsletter-api/internal/service"
	"newsletter-api/internal/worker"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, pool *worker.Pool) *gin.Engine {
	router := gin.Default()

	repo := repository.NewSubscriberRepository(db)
	svc := service.NewSubscriberService(repo, pool)
	handler := handlers.NewSubscriberHandler(svc)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.POST("/subscribe", handler.Subscribe)

	return router
}
