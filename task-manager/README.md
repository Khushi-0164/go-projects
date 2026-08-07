# 🗂️ Task Manager API
A REST API for managing projects and tasks with per-project role-based access control, demonstrating clean layered architecture, JWT authentication, and database transactions in Go.

## 🚀 Features
* **Auth**: Signup/login with bcrypt-hashed passwords and JWT-based sessions.
* **Projects**: Create projects, list your own, invite members by email.
* **Per-Project Roles**: Each user has a role (`owner` / `admin` / `member`) scoped to a specific project — not a global permission.
* **Tasks**: Create, list, and update task status within a project, restricted to project members only.
* **Database**: PostgreSQL via GORM, with auto-migrations.

## 🛠️ Go Concepts Demonstrated
* **Layered Architecture**: Codebase segmented into `handlers` (HTTP), `service` (business rules), `repository` (database access), and `models` (data shape) — each layer depends only on the one beneath it.
* **Transactions**: Creating a project and its owner membership atomically (`db.Transaction(...)`), so a failure midway never leaves an orphaned project.
* **Join Tables with Payload**: `ProjectMember` connects `User` and `Project` while also carrying a `role`, modeling many-to-many relationships that need extra data on the relationship itself.
* **Composite Unique Constraints**: A DB-level `uniqueIndex` across `(project_id, user_id)` guarantees no duplicate memberships, rather than trusting application code alone.
* **Pointer Fields for Optional Data**: `AssigneeID *uint` distinguishes "unassigned" (`nil`) from "assigned to user 0" — a plain `uint` couldn't represent that.
* **Custom Sentinel Errors**: The service layer returns typed errors (`ErrNotAuthorized`, `ErrUserNotFound`, `ErrTaskNotFound`); handlers use `errors.Is` to map them to the correct HTTP status, keeping HTTP concerns out of business logic.
* **Middleware & Request Context**: JWT validation runs once in middleware and stores `user_id` in `gin.Context`, available to every downstream handler via `c.GetUint("user_id")`.
* **Integration Testing**: `go test` coverage on the service layer, verifying both success paths and permission-rejection paths against a real test database.

## 📖 Setup

1. Copy `.env.example` to `.env` and set your DB credentials.
2. Create the database:
```bash
   createdb taskmanager
```
3. Install dependencies and run:
```bash
   go mod tidy
   go run ./cmd
```
   Tables are auto-migrated on startup.

## 📖 API

### Auth
```bash
POST /auth/signup   { "email", "name", "password" }
POST /auth/login    { "email", "password" }   -> returns JWT
```

### Projects (require `Authorization: Bearer <token>`)
```bash
POST   /api/projects              { "name", "description" }
GET    /api/projects              # projects you're a member of
POST   /api/projects/:id/members  { "email", "role" }   # owner/admin only
```

### Tasks (require `Authorization: Bearer <token>`)
```bash
POST   /api/projects/:id/tasks       { "title", "description", "assignee_id" }
GET    /api/projects/:id/tasks
PATCH  /api/tasks/:taskId/status     { "status" }   # todo | in_progress | done
```

## 🧪 Tests
```bash
go test ./internal/service/... -v
```
Covers: automatic ownership on project creation, permission checks on member invites, permission checks on task creation/status updates for non-members.

## 📦 Project Structure
```
cmd/main.go              entrypoint
config/                  env vars + DB connection
internal/models/         User, Project, ProjectMember, Task
internal/repository/     pure database access, no business rules
internal/service/        business rules + permission checks, no HTTP knowledge
internal/handlers/       thin HTTP layer: parse request, call service, format response
internal/middleware/     JWT auth middleware
internal/routes/         route wiring
internal/utils/          password hashing, small helpers
```