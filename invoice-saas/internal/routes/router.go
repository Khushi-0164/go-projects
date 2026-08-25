package routes

import (
	"net/http"

	"invoice-saas/internal/handlers"
	"invoice-saas/internal/middleware"
	"invoice-saas/internal/repository"
	"invoice-saas/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	orgRepo := repository.NewOrganizationRepository(db)
	orgService := service.NewOrganizationService(orgRepo)
	orgHandler := handlers.NewOrganizationHandler(orgService)

	customerRepo := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handlers.NewCustomerHandler(customerService, orgRepo)

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
		api.POST("/organizations", orgHandler.CreateOrganization)
		api.GET("/organizations", orgHandler.ListMyOrganizations)
		api.POST("/organizations/:id/members", orgHandler.AddMember)
	}

	api.POST("/organizations/:id/customers", customerHandler.CreateCustomer)
	api.GET("/organizations/:id/customers", customerHandler.ListCustomers)
	return router
}
