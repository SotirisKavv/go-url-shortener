# URL Shortener API — PostgreSQL persistence and analytics in Go

A production-ready RESTful API for URL shortening with PostgreSQL persistence, analytics, and enterprise architecture. Demonstrates advanced Go web development, database design, and scalable microservice patterns.

Quick links:
- Entrypoint: `urlshorten.go`
- Router: `router/router.go`
- Handler: `handler/url_handler.go`
- Model: `model/url.go`
- Repository: `repository/postgress_repo.go`

---

## 🚀 What is this?

A comprehensive URL shortening service with click tracking, expiration management, and analytics. Showcases advanced Go patterns including clean architecture, database integration, and production-ready API design.

---

## ✨ Features

- **URL Shortening:** Generate unique 6-character codes with SHA256-based hashing
- **PostgreSQL Integration:** Robust database persistence with proper connection pooling
- **Click Analytics:** Track usage statistics and top-performing URLs
- **Expiration Management:** Custom TTL for shortened URLs with automatic cleanup
- **RESTful Design:** Clean API endpoints with proper HTTP status codes
- **Scalable Architecture:** Repository pattern supporting multiple storage backends

---

## 🦄 Go Skills Demonstrated

- **Database Programming:** PostgreSQL integration with `pgx` driver
- **Clean Architecture:** Repository pattern with dependency injection
- **Advanced HTTP:** Custom routing with Gorilla Mux and middleware
- **Cryptography:** SHA256 hashing and Base62 encoding for URL generation
- **Environment Configuration:** Secure config management with `.env` files
- **Error Handling:** Comprehensive validation and HTTP status management

---

## 🛠️ Usage

```powershell
# Setup database and environment
"DATABASE_URL=postgres://user:pass@localhost/urldb" | Out-File -Encoding utf8 .env
go mod tidy

# Start API server
go run urlshorten.go  # Server listens on :8080

# Shorten URL
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com", "expires_in": "24h"}'

# Access shortened URL
curl http://localhost:8080/{short_code}  # Redirects to original URL

# Get analytics
curl http://localhost:8080/clicks  # Top 5 most-clicked URLs
```

---

## 🎯 Learning Objectives

This project demonstrates:
- **Enterprise Architecture:** Production-ready microservice design patterns
- **Database Design:** PostgreSQL schema design and query optimization
- **Security:** URL validation, input sanitization, and secure hashing
- **Scalability:** Repository pattern enabling horizontal scaling

---

## Folder map

- `urlshorten.go`: App entry and wiring
- `router/`: Route registration
- `handler/url_handler.go`: HTTP handlers
- `model/url.go`: URL domain model
- `repository/postgress_repo.go`: Postgres repository
- `utils/utils.go`: Helpers


## Next steps (ideas)

- Add rate limiting and auth for URL creation
- Add Redis cache for hot shortened URLs
- Add custom alias reservation and conflict resolution


**Author:** IAmSotiris
