# 🔌 Todo REST API (Standard Library)
A RESTful API built purely using Go's standard library to perform CRUD operations on tasks.

## 🚀 Features
* **Get All Todos** (`GET /todos`)
* **Get Todo by ID** (`GET /todos/:id`)
* **Create Todo** (`POST /todos/add`)
* **Update Todo** (`PUT /todos/:id`)
* **Delete Todo** (`DELETE /todos/:id`)

## 🛠️ Go Concepts Demonstrated
* **HTTP Routing**: Creating multiplex routes using `http.HandleFunc` and starting listeners with `http.ListenAndServe`.
* **Request Processing**: Routing requests based on the HTTP verb properties (`r.Method`).
* **Manual Route Parsing**: Splitting path strings (`strings.Split`) to manually read parameter IDs from URL queries.
* **JSON Streams**: Decoding incoming JSON request payloads via `json.NewDecoder` and encoding HTTP responses via `json.NewEncoder`.
* **HTTP Headers**: Setting response content types and custom HTTP response status codes.

## 📖 How to Run
1. **Start the API server:**
   ```bash
   go run main.go
   ```
   *The server runs on `http://localhost:8082`.*
2. **Test Endpoints (using curl or Postman):**
   * **List:** `curl http://localhost:8082/todos`
   * **Create:** `curl -X POST -d '{"title":"Learn Go APIs"}' http://localhost:8082/todos/add`
   * **Get by ID:** `curl http://localhost:8082/todos/1`
   * **Delete:** `curl -X DELETE http://localhost:8082/todos/1`
