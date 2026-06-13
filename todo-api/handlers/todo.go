package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todo-api/models"
)

var Todos = []models.Todo{
	{
		ID:        1,
		Title:     "Learn Go",
		Completed: false,
	},
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetById(w, r)
	case http.MethodDelete:
		DeleteTodo(w, r)
	case http.MethodPut:
		UpdateTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getIDFromURL(r *http.Request) (int, error) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		return 0, fmt.Errorf("id missing")
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, fmt.Errorf("invalid id")
	}

	return id, nil
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Todos)
}

func PostTodos(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if todo.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	todo.ID = len(Todos) + 1
	Todos = append(Todos, todo)

	json.NewEncoder(w).Encode(todo)
}

func GetById(w http.ResponseWriter, r *http.Request) {

	id, err := getIDFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, todo := range Todos {
		if todo.ID == id {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, todo := range Todos {
		if todo.ID == id {

			// remove item from slice
			Todos = append(Todos[:i], Todos[i+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Todo deleted successfully",
			})
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var updatedTodo models.Todo
	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, todo := range Todos {
		if todo.ID == id {
			//update item
			Todos[i].Title = updatedTodo.Title
			Todos[i].Completed = updatedTodo.Completed

			json.NewEncoder(w).Encode(Todos[i])
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}
