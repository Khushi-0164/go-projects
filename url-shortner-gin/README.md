# 🔗 URL Shortener (Gin & GORM)
A full-featured URL shortener service built with the Gin web framework, GORM, and PostgreSQL. It features user registration/login, JWT authentication, link management, and click tracking analytics.

## 🚀 Features
* **User Authentication**: Sign up and log in to get a JWT token.
* **Link Management**: Scoped link creation (`POST /api/links`), listing (`GET /api/links`), and deletion (`DELETE /api/links/:id`) for logged-in users.
* **Redirect & Analytics**: Public redirection via `/r/:code` that automatically increments the link's click count.
* **Auto-Migrations**: Automatically creates necessary database tables on server startup.

## 🛠️ Go Concepts Demonstrated
* **JWT-Authenticated Middleware**: Creating and verifying tokens for API endpoint security.
* **Redirection & DB State Updates**: Handling HTTP redirects (`302 Found`) while concurrently updating database records (incrementing click count).
* **GORM ORM**: Managing relationships between `User` and `Link` models and executing queries.
* **Config Management**: Dynamic loading of configuration variables using the `joho/godotenv` package.

## 📖 How to Run
Ensure your PostgreSQL database is running, copy `.env.example` to `.env` with correct database credentials, then start the server:
```bash
go run cmd/main.go
```
Example Output:
```bash
2026/08/18 11:51:00 Server starting on :8080
[GIN-debug] Listening and serving HTTP on :8080
```

To register a user:
```bash
curl -X POST http://localhost:8080/auth/signup \
  -H "Content-Type: application/json" \
  -d '{"username":"khushi","password":"securepassword"}'
```

To shorten a URL (using JWT authorization token):
```bash
curl -X POST http://localhost:8080/api/links \
  -H "Authorization: Bearer <your_jwt_token>" \
  -H "Content-Type: application/json" \
  -d '{"original_url": "https://google.com", "short_code": "ggl"}'
```
Example Output:
```json
{
  "id": 1,
  "short_code": "ggl",
  "original_url": "https://google.com",
  "user_id": 1,
  "clicks": 0,
  "created_at": "2026-08-18T11:51:00Z"
}
```
