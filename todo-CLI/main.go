package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"done"`
}

func GetNextID(todos []Todo) int {
	maxID := 0

	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}

	return maxID + 1
}

func LoadTodos() ([]Todo, error) {
	a, err := os.ReadFile("todos.json")
	if err != nil {
		return nil, err
	}

	if len(a) == 0 {
		return []Todo{}, nil
	}

	var todos []Todo

	err = json.Unmarshal(a, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
func saveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("todos.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func AddTodo(todos []Todo, title string) []Todo {
	id := GetNextID(todos)
	newTodo := Todo{
		ID:        id,
		Title:     title,
		Completed: false,
	}
	todos = append(todos, newTodo)
	return todos

}
func ListTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("No todos avilable")
		return
	}
	fmt.Println("Your Todo's :")
	for _, todo := range todos {
		status := "[]"
		if todo.Completed {
			status = "[x]"
		}
		fmt.Println(todo.ID, status, todo.Title)
	}
}
func DoneTodo(todos []Todo, id int) []Todo {
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			fmt.Println("Marked Done:", todos[i].Title)
			return todos
		}
	}
	fmt.Println("No todo found with ID:", id)
	return todos
}

func DeleteTodo(todos []Todo, id int) []Todo {
	for i := range todos {
		if todos[i].ID == id {
			title := todos[i].Title
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Deleted successfully:", title)
			return todos
		}
	}
	fmt.Println("No todo found with ID:", id)
	return todos
}

func UpdateTodo(todos []Todo, id int, newTitle string) []Todo {
	for i := range todos {

		if todos[i].ID == id {

			oldTitle := todos[i].Title

			todos[i].Title = newTitle

			fmt.Println("Updated:", oldTitle, "->", newTitle)

			return todos
		}
	}

	fmt.Println("No todo found with ID:", id)
	return todos
}

func main() {
	todos, err := LoadTodos()
	if err != nil {
		fmt.Println("Error loading:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("go run main.go add \"task name \"")
		fmt.Println("go run main.go list")
		fmt.Println("go run main.go done <id>")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go add \"task\"")
			return
		}
		title := os.Args[2]
		todos = AddTodo(todos, title)

	case "list":
		ListTodos(todos)
		return

	case "completed":
		if len(os.Args) < 3 {
			fmt.Println("Please provide an ID")
			fmt.Println("Example: go run main.go done 1")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be a number, you gave:", os.Args[2])
			return
		}
		todos = DoneTodo(todos, id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide an ID")
			fmt.Println("Example: go run main.go delete 1")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be a number, you gave:", os.Args[2])
			return
		}
		todos = DeleteTodo(todos, id)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Please provide an ID")
			fmt.Println("Example: go run main.go delete 1")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be a number, you gave:", os.Args[2])
			return
		}
		NewTitle := os.Args[3]
		todos = UpdateTodo(todos, id, NewTitle)

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Try: add, list, update,delete")
		return
	}
	err = saveTodos(todos)
	if err != nil {
		fmt.Println("Error saving:", err)
	}
}
