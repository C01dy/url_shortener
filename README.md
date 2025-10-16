# URL Shortener in Go

[![Go Version](https://img.shields.io/badge/Go-1.22.2-00ADD8?style=flat&logo=go)](https://golang.org/)
[![SQLite](https://img.shields.io/badge/SQLite-3-003B57?style=flat&logo=sqlite)](https://www.sqlite.org/)

> 🎯 Learning project for studying Go

## 📋 About the Project

URL Shortener is a RESTful service for shortening URLs with a custom HTTP router implementation and support for multiple storage types. The project was created for in-depth learning of Go and its standard library.

### Motivation

As a backend developer with Node.js experience, I started learning Go to expand my technology stack. This project helped me understand the fundamental differences between Go and Node.js and master the idiomatic approach to development in this language.

## 🎓 What I Learned

### Go Language Basics
- **Structural types and interfaces**: implementing `LinkStorage` interface for different storage types
- **Composition over inheritance**: using struct embedding
- **Error handling**: Go-way approach with explicit error returns
- **Goroutines and channels**: graceful shutdown using `chan` for signals

### Standard Library
- **net/http**: creating HTTP server with timeouts and middleware
- **database/sql**: working with SQLite through standard interface
- **context**: managing request lifecycle and graceful shutdown
- **log/slog**: structured logging (new API in Go 1.21+)
- **encoding/json**: JSON parsing and serialization
- **gopkg.in/yaml.v3**: working with YAML configuration

### Architectural Patterns
- **Dependency Injection**: passing dependencies through function parameters
- **Repository Pattern**: abstracting storage work through interface
- **Factory Pattern**: creating handlers through factory functions
- **Middleware Pattern**: HTTP request processing chain

### Practices
- ✅ Modular architecture with clear separation of concerns
- ✅ Unit testing using the `testing` package
- ✅ Graceful shutdown for proper server termination
- ✅ Configuration through files and environment variables
- ✅ Structured logging of all operations

## 🏗️ Architecture

```
.
├── api/              # HTTP handlers and business logic
│   ├── handler.go    # Handlers for creating and redirecting links
│   ├── response.go   # Helper functions for HTTP responses
│   └── handler_test.go
├── cmd/              # Application entry point
│   └── main.go
├── config/           # Configuration management
│   └── config.go     # Loading from YAML + ENV overrides
├── middleware/       # HTTP middleware
│   └── logger.go     # Request logging
├── router/           # Custom HTTP router
│   ├── router.go     # Routing implementation without third-party frameworks
│   └── router_test.go
├── storage/          # Data storage layer
│   ├── memory.go     # In-memory storage
│   └── sqlite.go     # SQLite storage
└── urlshort/         # Basic URL shortener handlers
    ├── handler.go
    ├── parser.go
    └── *_test.go
```

## 🚀 Quick Start

### Requirements
- Go 1.22.2 or higher
- SQLite3

### Installation and Running

```bash
# Clone the repository
git clone <repository-url>
cd url_shortener

# Install dependencies
go mod download

# Copy configuration example
cp config.yaml.example config.yaml

# Build the project
go build -o bin/urlshort ./cmd

# Run the server
./bin/urlshort
```

Server will start on `http://localhost:8080`

## 📡 API

### Create Short Link

```bash
POST /api/v1/links
Content-Type: application/json

{
  "url": "https://example.com/very/long/url"
}
```

**Response:**
```json
{
  "short_url": "localhost:8080/aB3dE5"
}
```

### Redirect to Short Link

```bash
GET /{code}
```

Performs redirect (302) to the original URL.

## ⚙️ Configuration

Configuration is loaded from `config.yaml`:

```yaml
port: ":8080"
db_path: "links.db"
```

Environment variables have priority:
- `PORT` - server port (without colon)
- `DB_PATH` - path to SQLite database

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Detailed coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🎯 Future Plans

### In Development
- [ ] **Containerization**: Creating an optimized multi-stage Dockerfile
- [ ] **Router Refactoring**: 
  - Support for path parameters (`/api/:id`)
  - Support for wildcards and regex patterns
  - Query parameters parsing
  - Route-level middleware

### Ideas to Possibly Add

#### Backend & Infrastructure
- [ ] **Prometheus metrics**: Exposing metrics for monitoring
- [ ] **Health checks**: Endpoints for service health verification
- [ ] **Rate limiting**: Protection against API abuse
- [ ] **Caching**: Redis for hot links

#### Features
- [ ] **Link TTL**: Automatic deletion of expired links
- [ ] **Batch API**: Creating multiple links in one request

#### DevOps & Quality
- [ ] **CI/CD**: GitHub Actions for tests and deployment

#### Security
- [ ] **HTTPS-only mode**: Secure cookies and HSTS
- [ ] **API keys**: Authentication for link creation

## 🛠️ Technologies

- **Go 1.22.2** 
- **SQLite3** 


