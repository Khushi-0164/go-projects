# 📖 Book Library API (Gin & PostgreSQL)
A RESTful API built using the Gin framework and PostgreSQL to manage a library of books, supporting CRUD operations, pagination, and text-based search.

## 🚀 Features
* **Create Book** (`POST /books`)
* **List Books (with Pagination)** (`GET /books?page=1&limit=5`)
* **Get Book by ID** (`GET /books/:id`)
* **Update Book** (`PUT /books/:id`)
* **Delete Book** (`DELETE /books/:id`)
* **Search Books by Title** (`GET /books/search?title=...`)

## 🛠️ Go Concepts Demonstrated
* **Gin Web Framework**: Implementing handlers, query string parameters, path parameters, and request validation/binding using Gin.
* **SQL Database Integration**: Executing queries, updates, and transactions using the standard library `database/sql` and PostgreSQL.
* **Environment Configuration**: Loading application configuration parameters (database credentials, hostname, ports) dynamically from a `.env` file using the `joho/godotenv` library.
* **Pagination & Search**: Implementing pagination (offset/limit) and pattern-matching text searches (ILIKE queries) directly at the database layer.

## 📖 How to Run
1. **Set up the Database & Environment:**
   Ensure a PostgreSQL database exists and create a `.env` file in the root directory with the following variables:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=khushii
   DB_NAME=librarydb
   ```
2. **Start the API server:**
   ```bash
   go run main.go
   ```
   *The server runs on `http://localhost:8080`.*
3. **Test Endpoints (using curl or Postman):**
   * **Create:** `curl -X POST -H "Content-Type: application/json" -d '{"title":"The Go Programming Language","author":"Alan A. A. Donovan","published_year":2015,"genre":"Technology"}' http://localhost:8080/books`
   * **List with Pagination:** `curl "http://localhost:8080/books?page=1&limit=2"`
   * **Search:** `curl "http://localhost:8080/books/search?title=Go"`
   * **Update:** `curl -X PUT -H "Content-Type: application/json" -d '{"title":"The Go Programming Language (2nd Ed)","author":"Alan A. A. Donovan","published_year":2016,"genre":"Technology"}' http://localhost:8080/books/1`
   * **Delete:** `curl -X DELETE http://localhost:8080/books/1`
