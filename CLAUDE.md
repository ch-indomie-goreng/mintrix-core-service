# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

mintrix-backend is a production-ready Go REST API boilerplate built with Clean Architecture. It follows Handler → Service → Repository layering with Gin, GORM, and PostgreSQL. The project is currently in the scaffolding phase — architecture and conventions are defined in `EDD.md`, but source code is not yet implemented.

## Build & Development Commands

```bash
# Build the server
go build ./cmd/server

# Run the server
go run ./cmd/server

# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run a single test
go test -v -run TestName ./path/to/package

# Lint
golangci-lint run

# Format
gofmt -w .
goimports -w .

# Security scan
gosec ./...

# Database migrations (golang-migrate)
migrate -path migrations -database "$DATABASE_URL" up
migrate -path migrations -database "$DATABASE_URL" down

# Docker
docker-compose up -d
docker-compose down
```

## Architecture

### Request Flow

```
HTTP Request → Gin Router → Middleware → Handler → Service → Repository → GORM → PostgreSQL
```

- **Handlers** (`internal/handler/`): Thin — parse request, call service, write response. No business logic.
- **Services** (`internal/service/`): All business logic lives here. Depends on repository interfaces.
- **Repositories** (`internal/repository/`): Data access via GORM. Returns domain models, not GORM types.
- **Models** (`internal/models/`): Domain entities and DTOs. GORM tags live here, not on service types.

### Project Structure (Planned)

```
cmd/server/main.go          — Entry point, wires dependencies
config/                     — Environment-based configuration loading
migrations/                 — golang-migrate SQL files (up/down)
internal/
  auth/                     — JWT generation, validation, refresh token logic
  database/                 — GORM connection, connection pooling
  handler/                  — Gin HTTP handlers (thin)
  middleware/               — Auth middleware, request ID, logging, CORS
  models/                   — Domain models with GORM tags
  repository/               — Database access interfaces + implementations
  service/                  — Business logic interfaces + implementations
  websocket/                — WebSocket hub, client management
  utils/                    — Shared helpers
pkg/                        — Reusable library packages
Dockerfile
docker-compose.yml
.env.example
Makefile
```

### Tech Stack

| Concern          | Choice                    |
| ---------------- | ------------------------- |
| HTTP Router      | Gin                       |
| ORM              | GORM                      |
| Database         | PostgreSQL (Supabase dev) |
| Migrations       | golang-migrate            |
| Cache            | Redis                     |
| Auth             | JWT + bcrypt              |
| Validation       | go-playground/validator   |
| Config           | Environment variables     |
| API Docs         | Scalar (OpenAPI)          |
| Logging          | slog (structured)         |
| Testing          | Testify                   |
| Containerization | Docker + Docker Compose   |
| Reverse Proxy    | Nginx                     |
| CI/CD            | GitHub Actions            |

## Key Conventions

### API Design
- All endpoints MUST be prefixed with `/api/v1`.
- Health check at `/health`, readiness at `/ready`.
- RESTful resource naming: plural nouns, no verbs in URLs.

### Database
- **Never** use GORM `AutoMigrate` in production. All schema changes go through `golang-migrate` SQL migration files in `migrations/`.
- Each migration file must have both `up` and `down` scripts.

### Error Handling
- Centralized error handler middleware returns consistent JSON error responses.
- Services return wrapped errors with context; handlers translate them to HTTP status codes.

### Logging & Observability
- Use `slog` for structured logging throughout.
- Middleware must inject a request ID into every request context.
- Log at service boundaries (handler entry, service entry, repository entry).

### Dependency Injection
- Services and repositories are defined as interfaces.
- Dependencies are wired in `main.go` — no global state, no `init()` for DI.
- Use `context.Context` in every function signature that participates in a request.

### Go Standards
- Follow Effective Go and idiomatic Go conventions.
- Packages must be small and single-purpose.
- Return errors explicitly; never use panic for control flow.
- Prefer interfaces at the consumer, not the producer.
