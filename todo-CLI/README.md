# 📝 Todo List CLI
A persistent command-line task manager application that stores user tasks in a JSON database.

## 🚀 Features
* **Create**: Add new tasks.
* **Read**: List all current tasks with their completion status (`[ ]` or `[x]`).
* **Update**: Change a task's title.
* **Complete**: Mark tasks as done.
* **Delete**: Remove tasks from the list.
* **Persistence**: Automatically saves state to a local `todos.json` file.

## 🛠️ Go Concepts Demonstrated
* Data mapping via `struct` definitions with custom JSON metadata tags (e.g. ``json:"done"``).
* File system I/O operations using `os.ReadFile` and `os.WriteFile`.
* JSON encoding/decoding using `json.MarshalIndent` and `json.Unmarshal`.
* Slice manipulation, including the idiomatic slice deletion pattern: `append(slice[:i], slice[i+1:]...)`.

## 📖 How to Run
Execute different CRUD operations by passing commands as arguments:
* **Add a Task:**
  ```bash
  go run main.go add "Learn Go interfaces"
  ```
* **List Tasks:**
  ```bash
  go run main.go list
  ```
* **Complete a Task:**
  ```bash
  go run main.go completed 1
  ```
* **Update a Task:**
  ```bash
  go run main.go update 1 "Learn Go packages"
  ```
* **Delete a Task:**
  ```bash
  go run main.go delete 1
  ```
