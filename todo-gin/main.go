package main

import (
	"todo-gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/todos", handlers.GetTodos)
	r.POST("/todos", handlers.CreateTodo)
	r.GET("/todos/:id", handlers.GetTodoById)
	r.DELETE("/todos/:id", handlers.DeleteById)
	r.PUT("/todos/:id", handlers.UpdateTodo)

	r.Run(":8080")
}
