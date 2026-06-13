package main

import (
	"net/http"
	"todo-api/handlers"
)

func main() {
	http.HandleFunc("/todos", handlers.GetTodos)
	http.HandleFunc("/todos/add", handlers.PostTodos)
	http.HandleFunc("/todos/", handlers.TodoHandler)

	http.ListenAndServe(":8082", nil)
}
