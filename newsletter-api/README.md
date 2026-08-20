# 📬 Newsletter Signup API
A small signup API demonstrating three production concerns together: async background job processing, per-IP rate limiting, and full Docker containerization — built as a focused drill after retrofitting bookmark-api with pagination, mocking, logging, graceful shutdown, and migrations.

## 🚀 Features
* **Signup**: `POST /subscribe` saves an email and responds immediately — it does not wait for the "welcome email" to send.
* **Background job processing**: a pool of concurrent worker goroutines sends welcome emails asynchronously, decoupled from the HTTP response.
* **Rate limiting**: per-IP token bucket limiting on the signup endpoint, protecting it from abuse/spam.
* **Fully containerized**: app + Postgres run together via Docker Compose, with a healthcheck ensuring the app never starts before the database is truly ready.

## 🛠️ Go Concepts Demonstrated
* **Worker pool pattern**: `worker.Pool` starts a fixed number of goroutines, all reading from one shared buffered channel (`chan Job`). Multiple jobs are processed concurrently — proven with real logs showing three different `worker_id`s processing three signups within milliseconds of each other, and all three "sent" confirmations landing together ~2 seconds later (not staggered, which would indicate sequential processing).
* **Decoupling slow work from the response**: the service calls `pool.Enqueue(job)` and returns immediately — measured at ~12ms response time despite a simulated 2-second "email send" happening after the response was already sent.
* **Token bucket rate limiting**: `golang.org/x/time/rate`, wrapped in Gin middleware, with a separate bucket per client IP (`c.ClientIP()`), each allowing a burst before throttling to a steady refill rate.
* **`sync.Mutex` for protecting shared state**: unlike the channel-based coordination used elsewhere, the rate limiter's per-IP map is accessed by multiple concurrent request goroutines directly, so it's protected with a mutex — a different concurrency-safety tool for a different situation (protecting shared state vs. passing data between goroutines).
* **A real deadlock, found and fixed**: an early version accidentally called `defer mu.Lock()` instead of `defer mu.Unlock()`, causing every request after the first to hang indefinitely waiting on an already-held lock — a first-hand encounter with one of the most classic concurrency bugs.
* **Multi-stage Docker builds**: a `builder` stage compiles a static Linux binary; the final image is minimal Alpine plus just that binary — no Go toolchain or source code shipped.
* **Docker Compose service healthchecks**: `depends_on: condition: service_healthy` on the database service fixes a real startup race condition where the app container could start and fail to connect before Postgres had actually finished initializing, even though `db` had technically started first.

## 📖 Setup

### Locally
1. Copy `.env.example` to `.env`, set DB credentials.
2. `createdb newsletterdb`
3. `go mod tidy && go run ./cmd`

### With Docker
```bash
docker compose up --build
```
Runs the app and Postgres together; no local Go or Postgres installation needed.

## 📖 API
```bash
POST /subscribe   { "email" }   # rate-limited: 3 burst, refills 1 per 2s
GET  /health
```

## 🧪 Proving it works

**Background jobs — response is fast, "email" arrives later:**
```bash
curl -w "\nTime: %{time_total}s\n" -X POST localhost:8080/subscribe \
  -H "Content-Type: application/json" -d '{"email":"you@example.com"}'
```
Measured: **~12ms** response, with `"welcome email sent"` appearing in the logs ~2 seconds afterward.

**Rate limiting — burst then throttle:**
```bash
for i in 1 2 3 4 5 6; do
  curl -s -o /dev/null -w "Request $i: %{http_code}\n" -X POST localhost:8080/subscribe \
    -H "Content-Type: application/json" -d "{\"email\":\"burst$i@example.com\"}"
done
```
Measured: requests 1–3 → `201`, requests 4–6 → `429`.

## 📦 Project Structure
```
cmd/main.go              entrypoint
config/                  env vars, DB connection
internal/models/         Subscriber
internal/repository/     database access
internal/service/        business logic; enqueues background jobs
internal/worker/         worker pool: Job, Pool, concurrent processing
internal/middleware/     per-IP token bucket rate limiting
internal/handlers/       HTTP layer
internal/routes/         route wiring
Dockerfile                multi-stage build
docker-compose.yml         app + Postgres, with healthcheck-gated startup
```