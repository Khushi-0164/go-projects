# 🔖 Bookmark REST API
A RESTful API for managing web bookmarks, built as a hands-on retrofit project for practicing production-grade backend patterns on top of a deliberately simple domain: Gin + GORM + PostgreSQL, with pagination, interface-based testing, structured logging, graceful shutdown, and versioned database migrations.

## 🚀 Features
* Standard CRUD operations for managing bookmarks (Title, URL, Tags).
* **Pagination & filtering** on the list endpoint (`?page=`, `?limit=`, `?tag=`).
* **Structured JSON logging** with severity levels, not plain-text log lines.
* **Graceful shutdown** — in-flight requests are given time to finish on `SIGINT`/`SIGTERM` instead of being killed abruptly.
* **Versioned, reversible database migrations** — no `AutoMigrate`; schema changes are explicit, reviewable SQL files.

## 🛠️ Go Concepts Demonstrated
* **Web Routing**: Gin routing with clean, grouped endpoints (`GET`, `POST`, `DELETE`).
* **Object-Relational Mapping (ORM)**: GORM for query generation against PostgreSQL.
* **Layered Architecture**: `handlers` → `service` → `repository`, each with a single responsibility.
* **Pagination math**: offset/limit queries plus a separate `Count()` query, with ceiling-division for `total_pages` and correct handling of empty/last-page edge cases.
* **Incremental query building**: conditionally chaining `.Where()` onto a GORM query builder so filtered and unfiltered list requests share one code path.
* **Interfaces + dependency injection for testing**: `BookmarkService` depends on a `BookmarkRepository` *interface*, not a concrete struct. A hand-written in-memory `fakeBookmarkRepository` satisfies that interface in tests — no database required, tests run in milliseconds. Proves Go's structural typing: the real repository needed zero changes to satisfy the new interface.
* **Structured logging (`log/slog`)**: JSON-formatted, leveled logs (`Info`/`Warn`/`Error`) with structured key-value fields instead of interpolated strings — built on Go's standard library, no external dependency.
* **Graceful shutdown**: the HTTP server runs in its own goroutine; `main()` blocks on an OS signal channel, then calls `srv.Shutdown(ctx)` with a bounded `context.WithTimeout`, draining in-flight requests before exit.
* **Versioned migrations (`golang-migrate`)**: numbered `up`/`down` SQL file pairs, checked into `db/migrations/`, tracked by a `schema_migrations` table — replacing `AutoMigrate` entirely, so schema changes are deliberate and reversible.

## 📖 Setup

1. Install [`golang-migrate`](https://github.com/golang-migrate/migrate): `brew install golang-migrate`
2. Copy `.env.example` to `.env` and set your DB credentials.
3. Create the database:
```bash
   createdb bookmarkdb
```
4. Run migrations to build the schema:
```bash
   migrate -database "postgres://<user>:<password>@localhost:5432/bookmarkdb?sslmode=disable" -path db/migrations up
```
5. Install dependencies and run:
```bash
   go mod tidy
   go run ./cmd
```

## 📖 API

```bash
POST   /api/bookmarks                          { "title", "url", "tags" }
GET    /api/bookmarks?page=1&limit=5&tag=go     # paginated, optionally filtered
DELETE /api/bookmarks/:id
```

Example:
```bash
curl -X POST localhost:8080/api/bookmarks \
  -H "Content-Type: application/json" \
  -d '{"title": "Go Docs", "url": "https://go.dev", "tags": "go,docs"}'
```
```json
{"id":1,"title":"Go Docs","url":"https://go.dev","tags":"go,docs","created_at":"...","updated_at":"..."}
```

```bash
curl "localhost:8080/api/bookmarks?page=1&limit=3&tag=go"
```
```json
{"data":[...],"page":1,"limit":3,"total":2,"total_pages":1}
```

## 🧪 Tests
```bash
go test ./internal/service/... -v
```
Runs entirely against an in-memory fake repository — no database connection needed, tests complete in milliseconds.

## 🗄️ Managing migrations

Create a new migration:
```bash
migrate create -ext sql -dir db/migrations -seq <migration_name>
```

Apply all pending migrations:
```bash
migrate -database "postgres://<user>:<password>@localhost:5432/bookmarkdb?sslmode=disable" -path db/migrations up
```

Roll back the last migration:
```bash
migrate -database "postgres://<user>:<password>@localhost:5432/bookmarkdb?sslmode=disable" -path db/migrations down 1
```

## 📦 Project Structure
```
cmd/main.go              entrypoint; graceful shutdown, structured logging setup
config/                  env vars + DB connection
db/migrations/           versioned up/down SQL migration files
internal/models/         Bookmark
internal/repository/     database access (implements the service-layer interface)
internal/service/        business logic, interface-based for testability, +tests
internal/handlers/       HTTP layer, pagination/filter query parsing, request logging
internal/routes/         route wiring
```