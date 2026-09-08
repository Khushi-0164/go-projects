package routes

import (
	"net/http"
	"time"

	"invoice-saas/internal/handlers"
	"invoice-saas/internal/middleware"
	"invoice-saas/internal/repository"
	"invoice-saas/internal/service"
	"invoice-saas/internal/worker"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cachedSummaryRepo *repository.CachedSummaryRepository, pool *worker.Pool) *gin.Engine {
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

	invoiceRepo := repository.NewInvoiceRepository(db, cachedSummaryRepo)
	invoiceService := service.NewInvoiceService(invoiceRepo)
	invoiceHandler := handlers.NewInvoiceHandler(invoiceService, orgRepo)

	paymentService := service.NewPaymentService(invoiceRepo)
	paymentHandler := handlers.NewPaymentHandler(paymentService, orgRepo)

	summaryService := service.NewSummaryService(cachedSummaryRepo)
	summaryHandler := handlers.NewSummaryHandler(summaryService, orgRepo)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	webhookHandler := handlers.NewWebhookHandler(pool)
	auth := router.Group("/auth")
	auth.Use(middleware.RateLimit(rate.Every(2*time.Second), 5))
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

	api.POST("/organizations/:id/invoices", invoiceHandler.CreateInvoice)
	api.GET("/organizations/:id/invoices", invoiceHandler.ListInvoices)
	api.GET("/organizations/:id/invoices/:invoiceId", invoiceHandler.GetInvoice)
	api.PATCH("/organizations/:id/invoices/:invoiceId/status", invoiceHandler.UpdateStatus)
	api.POST("/organizations/:id/invoices/:invoiceId/checkout", paymentHandler.CreateCheckoutSession)

	api.GET("/organizations/:id/summary", summaryHandler.GetSummary)

	router.POST("/webhooks/stripe", middleware.RateLimit(rate.Every(time.Second), 20), webhookHandler.StripeWebhook)
	return router
}
