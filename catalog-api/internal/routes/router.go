package routes

import (
	"catalog-api/internal/handlers"
	"catalog-api/internal/repository"
	"catalog-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cache *redis.Client) *gin.Engine {
	router := gin.Default()

	productRepo := repository.NewProductRepository(db, cache)
	productService := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	api := router.Group("/api")
	{
		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products", productHandler.ListProducts)
		api.GET("/products/:id", productHandler.GetProduct)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)
	}
	return router
}
