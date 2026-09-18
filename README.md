# Production HTTP Server

A production-oriented HTTP server built from scratch in **Go** using the standard `net/http` package.

## Features

* Custom HTTP router
* Middleware architecture
* Request ID generation and context propagation
* Request logging with response status and duration
* Panic recovery middleware
* Health check endpoint
* Configurable server timeouts
* Graceful shutdown with `SIGINT` / `SIGTERM`

## Project Structure

```text
.
├── cmd/
│   └── server/
│       └── main.go
└── internal/
    ├── middleware/
    │   ├── Logger.go
    │   ├── recovery.go
    │   └── requestId.go
    └── router/
        └── router.go
```

## Running

```bash
go run ./cmd/server
```

The server starts on:

```text
http://localhost:8080
```

### Available Endpoints

```text
GET /          → Server response
GET /health    → Health check
GET /panic     → Panic recovery test
```

## Purpose

This project was built to understand the fundamentals of production HTTP server architecture in Go, including routing, middleware, observability, error recovery, timeouts, and graceful lifecycle management.
