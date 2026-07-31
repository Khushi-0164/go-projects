# 📓 Notes REST API (Gin & PostgreSQL)
A RESTful HTTP API built using the Gin framework and PostgreSQL to perform CRUD operations on personal notes.

## 🚀 Features
* **Welcome Endpoint** (`GET /`)
* **Create Note** (`POST /notes`)
* **Get All Notes** (`GET /notes`)
* **Get Note by ID** (`GET /notes/:id`)
* **Update Note** (`PUT /notes/:id`)
* **Delete Note** (`DELETE /notes/:id`)

## 🛠️ Go Concepts Demonstrated
* **Gin Web Framework**: Utilizing Gin's context processing, JSON parameter binding, and response rendering helper functions.
* **SQL Database Integration**: Establishing connection to PostgreSQL via `database/sql` and performing queries/executions.
* **Structured Payload Handling**: Extracting payload structures and responding back with formatted JSON payloads.

## 📖 How to Run
Ensure your PostgreSQL server is running and a database named `notesdb` exists, then:
1. **Start the API server:**
   ```bash
   go run .
   ```
   *The server runs on `http://localhost:8080`.*
2. **Test Endpoints (using curl or Postman):**
   * **Welcome:** `curl http://localhost:8080/`
   * **Create:** `curl -X POST -H "Content-Type: application/json" -d '{"title":"Meeting Notes","content":"Discuss project milestones"}' http://localhost:8080/notes`
   * **Get All:** `curl http://localhost:8080/notes`
   * **Get by ID:** `curl http://localhost:8080/notes/1`
   * **Update:** `curl -X PUT -H "Content-Type: application/json" -d '{"title":"Meeting Notes (Updated)","content":"Discuss milestones and budget"}' http://localhost:8080/notes/1`
   * **Delete:** `curl -X DELETE http://localhost:8080/notes/1`
