# Go Backend REST API Boilerplate

## Tech Stack

| Category          | Technology                            |
| ----------------- | ------------------------------------- |
| Language          | Go                                    |
| HTTP Router       | Gin                                   |
| ORM               | GORM                                  |
| Database          | PostgreSQL (Supabase for development) |
| Migration         | golang-migrate                        |
| Cache             | Redis                                 |
| Authentication    | JWT + bcrypt                          |
| Validation        | go-playground/validator               |
| Configuration     | Environment Variables (.env)          |
| API Documentation | Scalar (OpenAPI)                      |
| Logging           | slog                                  |
| Testing           | Testify                               |
| Containerization  | Docker + Docker Compose               |
| Reverse Proxy     | Nginx                                 |
| CI/CD             | GitHub Actions                        |

---

# Architecture

```text
HTTP Request
      │
      ▼
Gin Router
      │
Middleware
      │
Handler
      │
Service
      │
Repository
      │
GORM
      │
PostgreSQL
```

---

# Project Structure

```text
backend-api/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── config/
├── docs/
├── migrations/
│
├── internal/
│   ├── auth/
│   ├── database/
│   ├── handler/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── service/
│   ├── websocket/
│   └── utils/
│
├── pkg/
│
├── Dockerfile
├── docker-compose.yml
├── .env.example
├── Makefile
├── go.mod
└── README.md
```

---

# Initial Features

* Health Check
* User Authentication
* Register
* Login
* JWT Authentication
* Refresh Token
* User CRUD
* Pagination
* Filtering
* Search
* Request Validation
* Global Error Handler
* Structured Logging
* Database Migration
* Docker Support
* API Documentation (Scalar)

---

# Future Features

* WebSocket
* Redis Pub/Sub
* Notification Service
* Email Service
* File Upload
* Password Reset
* Email Verification
* Role Based Access Control (RBAC)
* Rate Limiting
* Background Worker
* Monitoring
* Metrics
* Kubernetes Deployment

---

# Development Principles

* Clean Architecture
* SOLID Principles
* Repository Pattern
* Service Layer Pattern
* Dependency Injection
* Environment-based Configuration
* Graceful Shutdown
* Consistent Error Handling
* Structured Logging
* RESTful API Design
* OpenAPI-first Documentation

---

# Coding Standards

* Follow Effective Go conventions.
* Use idiomatic Go.
* Keep packages small and focused.
* Avoid global mutable state.
* Prefer interfaces only when needed.
* Return errors explicitly.
* Use context.Context in every request path.
* Write unit tests for business logic.
* Keep handlers thin; business logic belongs in services.

---

# Future Real-time Architecture

```text
Client
    │
WebSocket
    │
Connection Hub
    │
Service Layer
    │
Redis Pub/Sub
    │
PostgreSQL
```

---

## Goal

Build a **production-ready Go Backend Boilerplate** that supports REST APIs today and can scale to real-time applications with WebSocket, Redis Pub/Sub, Docker, and cloud deployment without major architectural changes.

---

Kalau ini akan dipakai bersama **Claude Code**, saya juga menyarankan menambahkan beberapa standar sejak awal agar hasil generate kodenya lebih konsisten:

* **Configuration:** Viper atau `.env` loader yang konsisten.
* **Linting:** `golangci-lint`.
* **Formatting:** `gofmt` + `goimports`.
* **Security Scan:** `gosec`.
* **Git Hooks:** pre-commit untuk format dan lint.
* **API Versioning:** gunakan prefix seperti `/api/v1`.
* **Database Migrations:** semua perubahan skema wajib melalui migration, tanpa `AutoMigrate` di production.
* **Observability:** siapkan struktur untuk request ID, structured logging, dan health/readiness endpoints (`/health`, `/ready`) agar mudah diintegrasikan dengan monitoring nanti.

Dengan pondasi tersebut, project akan lebih siap berkembang dari aplikasi CRUD sederhana menjadi backend yang layak digunakan di lingkungan production.
