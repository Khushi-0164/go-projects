# 🔖 Bookmark REST API
A simple RESTful API for managing web bookmarks, built using the Gin web framework, GORM, and PostgreSQL.

## 🚀 Features
* Standard CRUD operations for managing bookmarks (Title, URL, Tags).
* Connects to a PostgreSQL database utilizing GORM ORM.
* Auto-migrates database schemas on startup.

## 🛠️ Go Concepts Demonstrated
* **Web Routing**: Setting up Gin routing with clean endpoints (`GET`, `POST`, `DELETE`).
* **Object-Relational Mapping (ORM)**: Using GORM to interact with PostgreSQL, automating query generation and auto-migrations.
* **Architecture**: A clean modular layer layout (`handlers`, `service`, `repository`, `models`).

## 📖 How to Run
Ensure your PostgreSQL database `bookmarkdb` is running, then copy the environment example and start the server:
```bash
go run cmd/main.go
```
Example Output:
```bash
2026/08/18 11:50:00 [info] server starting on port 8080
[GIN-debug] Listening and serving HTTP on :8080
```

In another terminal, add a bookmark:
```bash
curl -X POST http://localhost:8080/api/bookmarks \
  -H "Content-Type: application/json" \
  -d '{"title": "Google", "url": "https://google.com", "tags": "search,tech"}'
```
Example Output:
```json
{
  "id": 1,
  "title": "Google",
  "url": "https://google.com",
  "tags": "search,tech",
  "created_at": "2026-08-18T11:50:00Z",
  "updated_at": "2026-08-18T11:50:00Z"
}
```
