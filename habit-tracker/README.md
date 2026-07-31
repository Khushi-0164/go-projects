# ✅ Habit Tracker API
A REST API built with Gin that lets users track daily habits, check in each day, and automatically calculates their current streak and completion rate.

## 🚀 Features
* User signup and login with hashed passwords.
* JWT-based authentication protecting all habit routes.
* Create, view, and delete habits — scoped to the logged-in user only.
* Daily check-ins per habit (one check-in per calendar day).
* Automatic current streak calculation (consecutive days completed).
* Automatic completion rate calculation (% of days completed since creation).
* Structured JSON logging for every request.
* Graceful shutdown on `SIGINT`/`SIGTERM`.

## 🛠️ Go & Gin Concepts Demonstrated
* REST routing and route groups (`router.Group`) with `gin.HandlerFunc`.
* Request binding and validation using `binding` struct tags.
* Custom middleware for JWT authentication (`c.Abort`, `c.Set`, `c.MustGet`).
* Password hashing and verification with `bcrypt`.
* JWT generation and parsing with `golang-jwt/jwt`.
* Database integration with GORM (SQLite) — models, `AutoMigrate`, ownership-scoped queries.
* Structured logging with `log/slog`, injected via closure-based middleware.
* Graceful server shutdown using `http.Server`, `signal.Notify`, and `context.WithTimeout`.
* Environment-based configuration with `godotenv`.
* Multi-package project structure (`models`, `handlers`, `middleware`).

## 📁 Project Structure
```
habit-tracker/
├── main.go
├── models/
│   ├── models.go
│   └── db.go
├── handlers/
│   ├── auth.go
│   └── habit.go
├── middleware/
│   ├── auth.go
│   ├── jwt.go
│   └── logger.go
├── .env.example
└── go.mod
```

## ⚙️ Setup

1. Clone the repo and move into the project folder.
2. Copy the example environment file and fill in your own secret:
```bash
cp .env.example .env
```
3. Install dependencies:
```bash
go mod tidy
```
4. Run the server:
```bash
go run main.go
```
Server starts on `http://localhost:8080`.

## 📖 API Endpoints

| Method | Endpoint              | Auth required | Description                          |
|--------|------------------------|:--------------:|--------------------------------------|
| POST   | `/signup`               | ❌             | Create a new user account            |
| POST   | `/login`                 | ❌             | Log in, returns a JWT token          |
| GET    | `/health`                | ❌             | Health check                         |
| POST   | `/habits`                | ✅             | Create a new habit                   |
| GET    | `/habits`                 | ✅             | List all habits with streak & % rate |
| GET    | `/habits/:id`             | ✅             | Get one habit and its check-in history |
| POST   | `/habits/:id/checkin`     | ✅             | Mark today as done for a habit       |
| DELETE | `/habits/:id`             | ✅             | Delete a habit                       |

Protected routes require a header:
```
Authorization: Bearer <token>
```

## 🧪 Example Usage

**Signup**
```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"khushi@test.com","password":"mypassword123"}'
```

**Login**
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"khushi@test.com","password":"mypassword123"}'
```

**Create a habit**
```bash
curl -X POST http://localhost:8080/habits \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name":"Drink water"}'
```

**Check in**
```bash
curl -X POST http://localhost:8080/habits/1/checkin \
  -H "Authorization: Bearer <token>"
```

**List habits**
```bash
curl http://localhost:8080/habits \
  -H "Authorization: Bearer <token>"
```
Example Output:
```json
[
  {
    "id": 1,
    "name": "Drink water",
    "current_streak": 1,
    "completion_rate": 100
  }
]
```
