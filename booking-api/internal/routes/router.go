package routes

import (
	"booking-api/internal/handlers"
	"booking-api/internal/middleware"
	"booking-api/internal/repository"
	"booking-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	resourceRepo := repository.NewResourceRepository(db)
	resourceHandler := handlers.NewResourceHandler(resourceRepo)

	bookingRepo := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepo, resourceRepo)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	auth := router.Group("/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
	}
	api := router.Group("/api")
	api.Use(middleware.AuthRequired())
	{
		api.POST("/resources", resourceHandler.CreateResource)
		api.GET("/resources", resourceHandler.ListResources)

		api.POST("/resources/:id/bookings", bookingHandler.CreateBooking)
		api.GET("/resources/:id/bookings", bookingHandler.ListForResource)
		api.GET("/my-bookings", bookingHandler.ListMyBookings)
	}
	return router
}
