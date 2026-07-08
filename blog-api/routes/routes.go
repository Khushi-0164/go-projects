package routes

import (
	"blog-api/handlers"
	"blog-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)

	router.GET("/blogs", handlers.GetAllBlogs)
	router.GET("/blogs/:id", handlers.GetBlogByID)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/blogs", handlers.CreateBlog)
	//protected.PUT("/blogs/:id", handlers.UpdateBlog)
	//protected.DELETE("/blogs/:id", handlers.DeleteBlog)
}
