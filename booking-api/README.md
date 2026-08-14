# 📅 Booking / Reservation API
A REST API for booking time slots on shared resources, with concurrency-safe overlap prevention using database-level row locking.

## 🚀 Features
* **Auth**: Signup/login with bcrypt-hashed passwords and JWT sessions.
* **Resources**: Create bookable resources (rooms, tables, desks — anything with a name).
* **Bookings**: Reserve a resource for a start/end time range.
* **Race-condition-safe**: Two simultaneous requests for an overlapping slot can never both succeed — enforced at the database transaction level, not just in application code.

## 🛠️ Go Concepts Demonstrated
* **`SELECT ... FOR UPDATE` row locking**: `BookingRepository.CreateIfAvailable` wraps the overlap check and insert in a single transaction, using `clause.Locking{Strength: "UPDATE"}` to lock relevant rows so a concurrent request is forced to wait rather than racing past the check.
* **The check-then-act race condition**: demonstrates why a plain "query, then insert" sequence is unsafe under concurrency, and how locking closes the gap between the two steps.
* **Interval overlap logic**: two ranges `[startA, endA)` and `[startB, endB)` overlap exactly when `startA < endB AND startB < endA` — a reusable pattern for any scheduling/calendar problem.
* **Layered architecture, applied selectively**: `ResourceHandler` calls its repository directly (no service layer) since resource creation has no business rules yet, while `BookingHandler` goes through a full service layer because overlap-checking and validation genuinely belong there — the layering is applied where it earns its keep, not uniformly by rule.
* **Sentinel errors across packages**: `ErrSlotUnavailable` is defined in the repository and re-exported through the service layer via `errors.Is`, avoiding fragile string-matching on error messages.
* **Native `time.Time` binding**: Gin/JSON automatically parses ISO-8601 strings (`"2026-08-20T14:00:00Z"`) directly into `time.Time` struct fields via `binding:"required"`, no manual date parsing needed.

## 📖 Setup

1. Copy `.env.example` to `.env` and set your DB credentials.
2. Create the database:
```bash
   createdb bookingdb
```
3. Install dependencies and run:
```bash
   go mod tidy
   go run ./cmd
```

## 📖 API

### Auth
```bash
POST /auth/signup   { "email", "name", "password" }
POST /auth/login    { "email", "password" }   -> returns JWT
```

### Resources (require `Authorization: Bearer <token>`)
```bash
POST /api/resources   { "name", "description" }
GET  /api/resources
```

### Bookings (require `Authorization: Bearer <token>`)
```bash
POST /api/resources/:id/bookings   { "start_time", "end_time" }   # ISO-8601
GET  /api/resources/:id/bookings   # all bookings for a resource
GET  /api/my-bookings              # your own bookings across all resources
```

## 🧪 Manually verifying the concurrency guarantee
Fire two overlapping booking requests at the same instant using shell background jobs:
```bash
curl -X POST localhost:8080/api/resources/1/bookings \
  -H "Authorization: Bearer <token>" -H "Content-Type: application/json" \
  -d '{"start_time":"2026-09-01T10:00:00Z","end_time":"2026-09-01T11:00:00Z"}' &
curl -X POST localhost:8080/api/resources/1/bookings \
  -H "Authorization: Bearer <token>" -H "Content-Type: application/json" \
  -d '{"start_time":"2026-09-01T10:00:00Z","end_time":"2026-09-01T11:00:00Z"}' &
wait
```
Exactly one request should return `201 Created`; the other should return `409 {"error":"time slot is unavailable"}` — never both `201`.

## 📦 Project Structure
```
cmd/main.go              entrypoint
config/                  env vars + DB connection
internal/models/         User, Resource, Booking
internal/repository/     database access, including the locking transaction
internal/service/        business rules: overlap policy, validation
internal/handlers/       thin HTTP layer
internal/middleware/     JWT auth
internal/routes/         route wiring
internal/utils/          password hashing, small helpers
```