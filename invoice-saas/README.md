# 💼 Invoice SaaS — Multi-Tenant Billing API
A production-hardened multi-tenant SaaS backend for invoicing and payments — organizations, customers, invoices with line items, real Stripe payment processing with verified webhooks, and every production concern (tests, pagination, caching, rate limiting, graceful shutdown) integrated into one coherent system.

This is the flagship project bringing together everything learned across a series of smaller, focused projects: JWT auth and RBAC (task-manager), interface-based testable architecture (bookmark-api), Redis cache-aside (catalog-api), background job worker pools (newsletter-api), database-level locking discipline (booking-api), and real-time systems thinking (chat-api).

## 🚀 Features
* **Multi-tenant**: every user belongs to one or more Organizations, each with its own isolated Customers and Invoices. Tenant isolation is enforced on every single query, not just at the API boundary.
* **Org-scoped RBAC**: owner/admin/member roles, scoped per-organization — a user's permissions differ depending on which org they're acting within.
* **Invoicing**: invoices with line items; totals are always computed server-side from line items, never trusted from client input. Money is stored as integer cents throughout, never floating point.
* **Real Stripe payments**: Stripe Checkout Sessions for one-off invoice payment, with cryptographically verified webhooks confirming payment success.
* **Background job processing**: webhook-triggered invoice status updates and cache invalidation happen asynchronously via a worker pool, so Stripe's webhook delivery is never blocked on database writes.
* **Cached org summaries**: revenue/outstanding-balance aggregation is cached with Redis (cache-aside pattern), invalidated automatically whenever an invoice's status changes.
* **Rate limited**: auth and webhook endpoints are protected with per-IP token-bucket rate limiting.
* **Paginated & filterable**: invoice listings support pagination and status filtering.
* **Graceful shutdown**: in-flight requests are drained on SIGINT/SIGTERM instead of being killed abruptly.
* **Versioned migrations**: schema is managed entirely through `golang-migrate`, no `AutoMigrate`, from day one.
* **Tested**: service-layer tests run against in-memory fake repositories — no database required, tests complete in milliseconds — covering tenant isolation, permission enforcement, and the invoice total calculation specifically.

## 🛠️ Architecture & Concepts Demonstrated

* **Layered architecture**: `handlers` (HTTP only) → `service` (business rules, interface-based) → `repository` (data access) — interfaces were defined from day one this time, not retrofitted, so every service was testable from its first commit.
* **Tenant isolation as a structural property, not a convention**: every repository method that touches org-scoped data *requires* an `orgID` parameter — there is no method that can fetch a customer or invoice by ID alone, making the most severe class of bug (cross-tenant data leakage) structurally hard to introduce.
* **Money handling**: all monetary values are `int64` cents, never `float64`, avoiding floating-point rounding errors in a domain where that would mean real accounting discrepancies.
* **Transactional writes**: organization creation (org + owner membership) and invoice creation (invoice + line items, with server-computed total) are both wrapped in database transactions — partial writes are structurally impossible.
* **Webhook signature verification**: every incoming Stripe webhook is cryptographically verified against a shared secret before being trusted — without this, anyone could forge a "payment succeeded" event and mark any invoice paid for free.
* **Cache-aside with write-path invalidation**: the org summary cache is invalidated from *two* separate write paths (direct status updates and the async webhook worker), since both can make it stale — a good example of a caching discipline that has to be applied everywhere data changes, not just once.
* **Real debugging, on a real integration**: this project involved diagnosing and fixing a duplicated file, a wrong SDK method call, a hardcoded prefix-matching bug, missing-arguments-after-refactor errors, a Stripe CLI OAuth/sandbox context mismatch, and a genuine Stripe API-version incompatibility — all resolved by reading actual error messages and reasoning through them systematically.

## 📖 Setup

### Prerequisites
* PostgreSQL, Redis running locally (or via Docker Compose)
* A Stripe account (test mode) — get your test secret key from `dashboard.stripe.com/test/apikeys`
* [`golang-migrate`](https://github.com/golang-migrate/migrate) CLI
* [Stripe CLI](https://stripe.com/docs/stripe-cli) for local webhook forwarding

### Local development
```bash
cp .env.example .env   # fill in DB, Redis, JWT, and Stripe credentials
createdb invoicesaas
migrate -database "postgres://<user>:<pass>@localhost:5432/invoicesaas?sslmode=disable" -path db/migrations up
go mod tidy
go run ./cmd
```

In a separate terminal, forward Stripe webhooks to your local server:
```bash
stripe listen --forward-to localhost:8080/webhooks/stripe --api-key <your_stripe_test_secret_key>
```
Copy the printed `whsec_...` signing secret into `.env` as `STRIPE_WEBHOOK_SECRET`.

### With Docker
```bash
docker compose up --build
```

## 📖 API

### Auth (rate limited)
```bash
POST /auth/signup   { "email", "name", "password" }
POST /auth/login    { "email", "password" }
```

### Organizations
```bash
POST /api/organizations                  { "name" }
GET  /api/organizations
POST /api/organizations/:id/members      { "email", "role" }
```

### Customers
```bash
POST /api/organizations/:id/customers    { "name", "email" }
GET  /api/organizations/:id/customers
```

### Invoices
```bash
POST  /api/organizations/:id/invoices                       { "customer_id", "line_items": [...] }
GET   /api/organizations/:id/invoices?page=&limit=&status=
GET   /api/organizations/:id/invoices/:invoiceId
PATCH /api/organizations/:id/invoices/:invoiceId/status      { "status" }
```

### Payments
```bash
POST /api/organizations/:id/invoices/:invoiceId/checkout    # returns a Stripe Checkout URL
POST /webhooks/stripe                                        # Stripe calls this (rate limited)
```

### Summary (cached)
```bash
GET /api/organizations/:id/summary
```

## 🧪 Tests
```bash
go test ./internal/service/... -v
```
11 tests covering transactional ownership creation, permission enforcement, cross-tenant isolation on every resource type, and correct invoice total computation — all against in-memory fakes, no database needed.

## 📦 Project Structure
```
cmd/main.go               entrypoint: graceful shutdown, Stripe init, DB/Redis/worker wiring
config/                   env vars, Postgres + Redis connections
db/migrations/            versioned schema, applied via golang-migrate
internal/models/          User, Organization, OrgMember, Customer, Invoice, InvoiceLineItem
internal/repository/      tenant-scoped data access + Redis cache-aside layer
internal/service/         business rules, interface-based, +tests with in-memory fakes
internal/handlers/        thin HTTP layer, including Stripe webhook verification
internal/middleware/      JWT auth, per-IP rate limiting
internal/worker/          background job pool for async webhook processing
internal/routes/          route wiring
Dockerfile                 multi-stage build
docker-compose.yml          app + Postgres + Redis
```