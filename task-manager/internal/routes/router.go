package routes

import (
	"net/http"

	"task-manager/internal/handlers"
	"task-manager/internal/middleware"
	"task-manager/internal/repository"
	"task-manager/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)
	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo)
	projectHandler := handlers.NewProjectHandler(projectService)
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

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
		api.POST("/projects", projectHandler.CreateProject)
		api.GET("/projects", projectHandler.ListMyProjects)
		api.POST("/projects/:id/members", projectHandler.AddMember)

		api.POST("/projects/:id/tasks", taskHandler.CreateTask)
		api.GET("/projects/:id/tasks", taskHandler.ListTasks)
		api.PATCH("/tasks/:taskId/status", taskHandler.UpdateTaskStatus)
	}

	return router
}
