# 📝 Blog REST API (Gin & PostgreSQL)
A RESTful API built using the Gin framework and PostgreSQL to perform CRUD operations on blogs with user registration, login, and JWT authentication.

## 🚀 Features
* **User Registration** (`POST /register`)
* **User Login** (`POST /login`) - returns a JWT token
* **Get All Blogs** (`GET /blogs`)
* **Get Blog by ID** (`GET /blogs/:id`)
* **Create Blog** (`POST /blogs`) - JWT Authentication required
* **Update Blog** (`PUT /blogs/:id`) - JWT Authentication required
* **Delete Blog** (`DELETE /blogs/:id`) - JWT Authentication required

## 🛠️ Go Concepts Demonstrated
* **Gin Web Framework**: Using Gin for routing, request binding/validation, routing group organization, and middleware execution.
* **SQL Database Integration**: Executing SQL queries (SELECT, INSERT, UPDATE, DELETE) using standard `database/sql` and `github.com/lib/pq` PostgreSQL driver.
* **JWT Authentication**: Creating, signing, and parsing JSON Web Tokens to protect endpoints.
* **Password Security**: Hashing passwords using `golang.org/x/crypto/bcrypt` prior to storage.

## 📖 How to Run
Ensure your PostgreSQL server is running and a database named `blogdb` exists, then:
1. **Start the API server:**
   ```bash
   go run main.go
   ```
   *The server runs on `http://localhost:8080`.*
2. **Test Endpoints (using curl or Postman):**
   * **Register:** `curl -X POST -H "Content-Type: application/json" -d '{"username":"khushii","email":"khushii@example.com","password":"securepassword"}' http://localhost:8080/register`
   * **Login:** `curl -X POST -H "Content-Type: application/json" -d '{"email":"khushii@example.com","password":"securepassword"}' http://localhost:8080/login`
   * **List Blogs:** `curl http://localhost:8080/blogs`
   * **Create Blog:** `curl -X POST -H "Authorization: Bearer <your_jwt_token>" -H "Content-Type: application/json" -d '{"title":"My First Post","content":"This is the content of my blog post"}' http://localhost:8080/blogs`
