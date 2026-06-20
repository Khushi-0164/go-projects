# ⚡ Todo REST API (Gin Framework)
A modernized RESTful API for managing tasks, built using the third-party web framework **Gin**.

## 🚀 Features
* **Get All Todos** (`GET /todos`)
* **Get Todo by ID** (`GET /todos/:id`)
* **Create Todo** (`POST /todos`)
* **Update Todo** (`PUT /todos/:id`)
* **Delete Todo** (`DELETE /todos/:id`)

## 🛠️ Go Concepts Demonstrated
* **Web Frameworks**: Setting up default Gin engine routers (`gin.Default()`).
* **Route Parameter Mapping**: Retrieving dynamic URL route values seamlessly via `c.Param("id")`.
* **JSON Binding**: Automatic request decoding and validation checking via `c.ShouldBindJSON()`.
* **Gin Response Builders**: Utilizing the `c.JSON()` method along with data map helper shortcuts (`gin.H`).
* **Clean Routing**: Registering endpoints directly to target handler methods matching HTTP verbs (`r.GET`, `r.POST`, etc.).

## 📖 How to Run
1. **Start the API server:**
   ```bash
   go run main.go
   ```
   *The server runs on `http://localhost:8080`.*
2. **Test Endpoints (using curl or Postman):**
   * **List:** `curl http://localhost:8080/todos`
   * **Create:** `curl -X POST -H "Content-Type: application/json" -d '{"title":"Learn Gin Framework"}' http://localhost:8080/todos`
   * **Get by ID:** `curl http://localhost:8080/todos/1`
   * **Delete:** `curl -X DELETE http://localhost:8080/todos/1`
