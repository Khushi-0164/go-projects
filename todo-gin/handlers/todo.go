package handlers

import (
	"strconv"
	"todo-gin/models"

	"github.com/gin-gonic/gin"
)

var Todos = []models.Todo{
	{
		ID:        1,
		Title:     "work",
		Completed: true,
	},
}

func GetTodos(c *gin.Context) {
	c.JSON(200, Todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	todo.ID = len(Todos) + 1
	Todos = append(Todos, todo)
	c.JSON(201, todo)
}

func GetTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	for _, todo := range Todos {
		if todo.ID == id {
			c.JSON(200, todo)
			return
		}
	}

	c.JSON(404, gin.H{
		"error": "Todo not found",
	})
}

func DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Id",
		})
		return
	}

	for i, todo := range Todos {
		if todo.ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			c.JSON(200, gin.H{
				"message": "Deleted Succesfully!",
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error": "Todo not found",
	})
}

func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Id",
		})
		return
	}

	var UpdatedTodo models.Todo

	if err := c.ShouldBindJSON(&UpdatedTodo); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Json",
		})
		return
	}

	for i, todo := range Todos {
		if todo.ID == id {
			Todos[i].Title = UpdatedTodo.Title
			Todos[i].Completed = UpdatedTodo.Completed

			c.JSON(200, Todos[i])
			return
		}
	}
	c.JSON(404, gin.H{
		"error": "Todo not found",
	})
}
