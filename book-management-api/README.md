# 📚 Book Management API
A RESTful HTTP API to manage a collection of books, supporting CRUD operations (Create, Read, Update, Delete) integrated with a PostgreSQL database.

## 🚀 Features
* Full CRUD database support for books (Title, Author).
* Connects to a PostgreSQL database for persistent storage.
* Handles JSON request and response bodies.

## 🛠️ Go Concepts Demonstrated
* **SQL Database Integration**: Executing SQL queries (SELECT, INSERT, UPDATE, DELETE) using standard `database/sql` and `github.com/lib/pq`.
* **REST Routing**: Handling multiple HTTP methods (GET, POST, PUT, DELETE) on a single endpoint path using a `switch` statement in handlers.
* **JSON Processing**: Decoding and encoding structs directly from/to HTTP streams with `json.NewDecoder` and `json.NewEncoder`.
* **Parameter Extraction**: Retrieving query parameters from request URLs using `r.URL.Query()`.

## 📖 How to Run
Ensure your PostgreSQL server is running and a database named `bookdb` exists, then start the server:
```bash
go run .
```
Example Output:
```bash
Connected to PostgreSQL!
Server is running on port 8080...
```

To add a new book:
```bash
curl -X POST http://localhost:8080/books -d '{"title":"The Go Programming Language","author":"Alan A. A. Donovan"}'
```
Example Output:
```bash
{"id":1,"title":"The Go Programming Language","author":"Alan A. A. Donovan"}
```
