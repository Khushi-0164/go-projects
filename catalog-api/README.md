# 🗄️ Product Catalog API (Redis Cache-Aside)
A minimal product CRUD API demonstrating the cache-aside caching pattern with Redis in front of PostgreSQL — with a measured, real speedup, not just a theoretical one.

## 🚀 Features
* **Product CRUD**: create, read, update, delete, list.
* **Cache-aside reads**: `GET /products/:id` checks Redis first, falls back to Postgres on a miss, and populates the cache before returning.
* **Explicit invalidation**: updates and deletes actively remove the stale Redis entry so no request ever sees outdated data.
* **TTL as a safety net**: cached entries also expire after 5 minutes independently of explicit invalidation, as a second layer of protection against staleness.

## 🛠️ Go Concepts Demonstrated
* **Cache-aside (lazy loading) pattern**: `ProductRepository.FindByID` — check cache → on miss, query the real source of truth → populate the cache → return. The most common caching pattern in real production systems.
* **Cache invalidation on write**: `Update`/`Delete` call `Cache.Del` after a successful database write, closing the gap where stale cached data could otherwise persist.
* **`context.Context` propagation**: `ctx` is threaded from the handler (`c.Request.Context()`) through the service and into every Redis/DB call, the idiomatic Go pattern for cancellation-aware I/O — even in layers that don't use it directly yet, so the whole chain stays cancellable.
* **JSON (de)serialization for cache storage**: `json.Marshal`/`json.Unmarshal` convert Go structs to/from strings, since Redis only stores raw bytes, not native Go types.
* **Deliberate artificial latency**: `time.Sleep(300 * time.Millisecond)` simulates an expensive query, making the cache's effect visible and measurable rather than theoretical.
* **Graceful degradation on cache errors**: a failed cache write (`Set`) doesn't fail the request — a cache miss just means the next request also misses, rather than caching becoming a new source of failures.

## 📖 Setup

1. Install and start Redis locally (e.g. `brew install redis && brew services start redis`), confirm with `redis-cli ping` → `PONG`.
2. Copy `.env.example` to `.env` and set DB/Redis config.
3. Create the database:
```bash
   createdb catalogdb
```
4. Install dependencies and run:
```bash
   go mod tidy
   go run ./cmd
```

## 📖 API
```bash
POST   /api/products        { "name", "description", "price" }
GET    /api/products         # list all
GET    /api/products/:id     # cache-aside read
PUT    /api/products/:id     { "name", "description", "price" }   # invalidates cache
DELETE /api/products/:id     # invalidates cache
```

## 🧪 Measuring the cache effect
```bash
curl -X POST localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Mechanical Keyboard","description":"Hot-swappable switches","price":89.99}'

# First call — cache miss, ~300ms
curl -w "\nTime: %{time_total}s\n" localhost:8080/api/products/1

# Second call — cache hit, ~1ms
curl -w "\nTime: %{time_total}s\n" localhost:8080/api/products/1
```
Measured result during development: **0.307s → 0.001s**, roughly a 280x speedup on the cached path.

## 📦 Project Structure
```
cmd/main.go              entrypoint; connects DB + Redis, runs migrations
config/                  env vars, Postgres + Redis connections
internal/models/         Product
internal/repository/     cache-aside logic lives here — the core of this project
internal/service/        thin business layer, threads context.Context through
internal/handlers/       HTTP layer
internal/routes/         route wiring
internal/utils/          small helpers
```