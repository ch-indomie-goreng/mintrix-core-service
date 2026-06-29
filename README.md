# Mintrix Backend

![Go](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/250px-Go_Logo_Blue.svg.png)

Production-ready Go REST API boilerplate built with **Clean Architecture**. Follows Handler → Service → Repository layering with Gin, GORM, and PostgreSQL.

---

## Quick Start

```bash
# Prerequisites: Go 1.26+, PostgreSQL 16+, Redis 7+
cp .env.example .env             # then fill in required variables
go mod download                  # install dependencies
go run ./cmd/server              # http://localhost:8080
```

### Docker

```bash
docker-compose up -d             # PostgreSQL + Redis
go run ./cmd/server              # run the API server
```

---

## Tech Stack

<table>
  <tr>
    <td align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/250px-Go_Logo_Blue.svg.png" height="40" /></td>
    <td align="center"><img src="https://miro.medium.com/v2/resize:fit:1400/0*kMJr4U_zX9g-Au-k.png" height="40" /></td>
    <td align="center"><img src="https://miro.medium.com/v2/resize:fit:1400/1*XBvxUxqycRC8B8KGCuzJVw.png" height="40" /></td>
    <td align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/29/Postgresql_elephant.svg/1280px-Postgresql_elephant.svg.png" height="40" /></td>
    <td align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/3c/Logo-redis_%28old%29.svg/3840px-Logo-redis_%28old%29.svg.png" height="40" /></td>
    <td align="center"><img src="https://s3.typoniels.de/typoniels-strapi/production/scalar_24e94da42a.webp" height="40" /></td>
    <td align="center"><img src="https://1000logos.net/wp-content/uploads/2020/08/Nginx-Logo.png" height="40" /></td>
  </tr>
  <tr>
    <td align="center"><strong>Go</strong><br/>Language</td>
    <td align="center"><strong>Gin</strong><br/>HTTP Router</td>
    <td align="center"><strong>GORM</strong><br/>ORM</td>
    <td align="center"><strong>PostgreSQL</strong><br/>Database</td>
    <td align="center"><strong>Redis</strong><br/>Cache</td>
    <td align="center"><strong>Scalar</strong><br/>API Docs</td>
    <td align="center"><strong>Nginx</strong><br/>Reverse Proxy</td>
  </tr>
  <tr>
    <td align="center"><img src="https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/png/jwt-io-light.png" height="40" /></td>
    <td align="center"><img src="https://media2.dev.to/dynamic/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Fuh5cdsepfg71q22ngiil.png" height="40" /></td>
    <td align="center"><img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQhrAOm3cpNRk3XMSQ11bj7oQLmLPfWGYZyE6nJRbmzPsyChZNdAuXXhas&s=10" height="40" /></td>
    <td align="center"></td>
    <td align="center"><img src="https://dpcksph9iv8aa.cloudfront.net/assets/branding/BrandIconWhitePurple_2024-12-03-111937_gtap.svg" height="40" /></td>
    <td align="center"><img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSO-NVwyGWKCwvNj16RnH0pN6gcsoTRPEUu2-HbOY7CVAFGdwzq3Oub1Tc&s=10" height="40" /></td>
    <td align="center"><img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQRKCnwDp8FRg3fvOxXrtxIiBLiP0N9nLbx0UzHbf_40w&s=10" height="40" /></td>
  </tr>
  <tr>
    <td align="center"><strong>JWT + bcrypt</strong><br/>Authentication</td>
    <td align="center"><strong>golang-migrate</strong><br/>Migrations</td>
    <td align="center"><strong>Docker</strong><br/>Containerization</td>
    <td align="center"><strong>slog</strong><br/>Structured Logging</td>
    <td align="center"><strong>Testify</strong><br/>Testing</td>
    <td align="center"><strong>GitHub Actions</strong><br/>CI/CD</td>
    <td align="center"><strong>golangci-lint</strong><br/>Linting</td>
  </tr>
</table>

**Also:** go-playground/validator · gosec · gofmt · goimports · WebSocket

---

## Environment Variables

| Variable | Description |
|---|---|
| `PORT` | Server port (default: `8080`) |
| `DATABASE_URL` | PostgreSQL connection string |
| `JWT_SECRET` | Secret key for JWT signing |
| `JWT_EXPIRY` | Access token expiry duration (e.g. `15m`) |
| `REFRESH_EXPIRY` | Refresh token expiry duration (e.g. `168h`) |

> **Security**: Never commit `.env` files to git. Only `.env.example` should be tracked in version control.

---

## Available Commands

```bash
# Build & Run
go build ./cmd/server           # Build the server binary
go run ./cmd/server             # Run the server

# Testing
go test ./...                   # Run all tests
go test -v ./...                # Run tests with verbose output
go test -v -run TestName ./pkg  # Run a single test

# Code Quality
gofmt -w .                      # Format code
goimports -w .                  # Format imports
golangci-lint run               # Lint
gosec ./...                     # Security scan

# Database Migrations
migrate -path migrations -database "$DATABASE_URL" up
migrate -path migrations -database "$DATABASE_URL" down

# Docker
docker-compose up -d             # Start PostgreSQL + Redis
docker-compose down              # Stop services
```

---

## Architecture

```
HTTP Request → Gin Router → Middleware → Handler → Service → Repository → GORM → PostgreSQL
```

### Request Flow

1. **HTTP Request** → Incoming request hits Gin Router
2. **Middleware Chain** → Request ID, CORS, Logging, Auth, Error Handler
3. **Handler** → Parses request, validates input, calls Service layer
4. **Service** → Business logic, orchestration, depends on Repository interfaces
5. **Repository** → Data access via GORM, returns domain models
6. **GORM + PostgreSQL** → Query execution and persistence

### Clean Architecture Layers

| Layer | Responsibility | Location |
|---|---|---|
| **Handler** | Parse request, call service, write response. No business logic. | `internal/handler/` |
| **Service** | All business logic. Depends on repository interfaces. | `internal/service/` |
| **Repository** | Data access via GORM. Returns domain models. | `internal/repository/` |
| **Models** | Domain entities and DTOs with GORM tags. | `internal/models/` |
| **Middleware** | Cross-cutting concerns: auth, logging, CORS, error handling. | `internal/middleware/` |

---

## API Design

| Concern | Convention |
|---|---|
| **Versioning** | All endpoints prefixed with `/api/v1` |
| **Health** | `GET /health` — liveness probe |
| **Readiness** | `GET /ready` — readiness probe |
| **Resources** | RESTful, plural nouns, no verbs in URLs |
| **Auth** | JWT Bearer token in `Authorization` header |
| **Errors** | Consistent JSON `{ "error": "...", "details": ... }` |
| **Docs** | Scalar (OpenAPI) at `/docs` |

### Endpoints (Initial)

| Method | Path | Description |
|---|---|---|
| `GET` | `/health` | Health check |
| `GET` | `/ready` | Readiness check |
| `POST` | `/api/v1/auth/register` | User registration |
| `POST` | `/api/v1/auth/login` | User login |
| `POST` | `/api/v1/auth/refresh` | Refresh JWT token |
| `GET` | `/api/v1/users` | List users (paginated) |
| `GET` | `/api/v1/users/:id` | Get user by ID |
| `PUT` | `/api/v1/users/:id` | Update user |
| `DELETE` | `/api/v1/users/:id` | Delete user |

---

## Project Folder Structure

```
mintrix-backend/
├── cmd/
│   └── server/
│       └── main.go              # Entry point, wires dependencies
│
├── config/
│   └── config.go                # Environment-based configuration
│
├── docs/                        # OpenAPI / Scalar documentation
├── migrations/                  # golang-migrate SQL files (up/down)
│
├── internal/
│   ├── auth/                    # JWT generation, validation, refresh token
│   │   └── jwt.go
│   │
│   ├── database/                # GORM connection, connection pooling
│   │   └── database.go
│   │
│   ├── handler/                 # Gin HTTP handlers (thin)
│   │   ├── auth_handler.go      # Auth endpoints
│   │   └── docs_handler.go      # Scalar API docs
│   │
│   ├── middleware/              # Cross-cutting middleware
│   │   ├── auth.go              # JWT auth middleware
│   │   ├── cors.go              # CORS configuration
│   │   ├── error_handler.go     # Centralized error handling
│   │   ├── logger.go            # Structured request logging
│   │   └── request_id.go        # Request ID injection
│   │
│   ├── models/                  # Domain models with GORM tags
│   │   ├── auth_dto.go          # Auth request/response DTOs
│   │   ├── refresh_token.go     # Refresh token model
│   │   └── user.go              # User model
│   │
│   ├── repository/              # Data access interfaces + implementations
│   │   ├── token_repository.go  # Refresh token persistence
│   │   └── user_repository.go   # User CRUD
│   │
│   ├── service/                 # Business logic interfaces + implementations
│   │   └── auth_service.go      # Auth service
│   │
│   ├── websocket/               # WebSocket hub, client management (planned)
│   └── utils/                   # Shared helpers
│       └── response.go          # Standard JSON response helpers
│
├── pkg/                         # Reusable library packages
│
├── docker-compose.yml           # PostgreSQL + Redis
├── Dockerfile                   # Multi-stage Go build
├── .env.example                 # Environment variable template
├── .gitignore
├── go.mod
├── go.sum
└── Makefile
```

---

## Development Principles

- **Clean Architecture** — clear separation of concerns
- **SOLID Principles** — maintainable and testable code
- **Repository Pattern** — abstracted data access
- **Service Layer Pattern** — business logic encapsulation
- **Dependency Injection** — wired in `main.go`, no global state
- **RESTful API Design** — consistent, predictable endpoints
- **OpenAPI-first** — Scalar for interactive API documentation
- **Graceful Shutdown** — clean connection teardown on SIGTERM

---

## Documentation

- **[CLAUDE.md](./CLAUDE.md)** — detailed architecture, conventions, build commands, and coding standards
- **[EDD.md](./EDD.md)** — engineering design document with feature roadmap and future architecture
- **[`.agents/`](./.agents/)** — planning logs and AI agent context

---

© 2026 Mintrix. All rights reserved.
