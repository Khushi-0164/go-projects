# 📝 Notes API
A RESTful API to manage text notes, supporting full CRUD operations (Create, Read, Update, Delete) using the Gin web framework and PostgreSQL for persistence.

## 🚀 Features
* Standard CRUD operations on notes (Title, Content).
* Connects to a PostgreSQL database for persistent data storage.
* Utilizes the Gin framework for routing and JSON parsing/binding.

## 🛠️ Go Concepts Demonstrated
* **Web Framework Integration**: Utilizing Gin for clean HTTP routing, path parameters (`/notes/:id`), and response builders.
* **SQL Database Access**: Running SQL queries (SELECT, INSERT, UPDATE, DELETE) using standard `database/sql` with PostgreSQL drivers.
* **JSON Binding**: Automatic binding of JSON request bodies into Go structures using `c.ShouldBindJSON`.
* **Resource Mapping**: Mapping HTTP methods (`GET`, `POST`, `PUT`, `DELETE`) to corresponding RESTful handlers.

## 📖 How to Run
Ensure your PostgreSQL database `notesdb` is running, then start the server:
```bash
go run .
```
Example Output:
```bash
2026/08/18 11:47:50 Connected to PostgreSQL!
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
[GIN-debug] GET    /                         --> main.welcomeHandler (3 handlers)
[GIN-debug] POST   /notes                    --> main.createNoteHandler (3 handlers)
[GIN-debug] GET    /notes                    --> main.getAllNotesHandler (3 handlers)
...
[GIN-debug] Listening and serving HTTP on :8080
```

In another terminal, create a note:
```bash
curl -X POST http://localhost:8080/notes \
  -H "Content-Type: application/json" \
  -d '{"title": "My Note", "content": "This is a test note."}'
```
Example Output:
```json
{
  "message": "Note created successfully",
  "note": {
    "id": 1,
    "title": "My Note",
    "content": "This is a test note."
  }
}
```
