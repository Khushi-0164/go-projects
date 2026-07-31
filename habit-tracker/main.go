package main

import (
	"context"
	"habit-tracker/handlers"
	"habit-tracker/middleware"
	"habit-tracker/models"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	models.InitDB()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.StructuredLogger(logger))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.POST("/habits", handlers.CreateHabit)
		protected.GET("/habits", handlers.ListHabits)
		protected.GET("/habits/:id", handlers.GetHabit)
		protected.POST("/habits/:id/checkin", handlers.CheckIn)
		protected.DELETE("/habits/:id", handlers.DeleteHabit)
	}

	// router.Run(":8080")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
