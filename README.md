# 🔗 URL Shortener API (Go) - **Advanced Level**

A production-ready RESTful API for URL shortening with PostgreSQL persistence, analytics, and enterprise architecture. Demonstrates advanced Go web development, database design, and scalable microservice patterns.

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

```sh
# Setup database and environment
echo "DATABASE_URL=postgres://user:pass@localhost/urldb" > .env
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

**Author:** IAmSotiris
