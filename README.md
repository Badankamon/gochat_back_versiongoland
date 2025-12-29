## GoChat Backend

Lightweight backend for the GoChat project — REST APIs and real-time services implemented in Go.

Module: `github.com/Badankamon/gochat_backend`
Go version: 1.25.5

### Overview

This repository contains the backend services for the GoChat application. It uses Gin for HTTP APIs, GORM with Postgres for persistence, Redis for caching, JWT for authentication and other libraries for configuration and logging.

Key features:
- RESTful APIs (Gin)
- Authentication with JWT
- Persistence via GORM (Postgres driver)
- Redis caching
- Configuration with Viper
- Structured logging (zap)

### Prerequisites

- Go 1.25 or later
- A Postgres database
- Redis (optional, for caching)

### Quick start

```bash
git clone https://github.com/Badankamon/gochat_backend.git
cd gochat_back_versiongoland
```

Create a `.env` file or export environment variables for database and JWT configuration. Example variables:

```
DATABASE_URL=postgres://user:pass@localhost:5432/gochatdb
REDIS_ADDR=localhost:6379
JWT_SECRET=your_secret_here
ENV=development
```

Install dependencies and run:

```bash
go mod download
go build -o gochat_backend ./...
./gochat_backend
# or run directly
go run ./...
```

If your project has an entrypoint under `cmd/`, run the appropriate package, e.g. `go run ./cmd/server`.

### Docker (optional)

If a `Dockerfile` or `docker-compose.yml` is present you can build and run using Docker. Example:

```bash
docker build -t gochat-backend .
docker run -e DATABASE_URL="$DATABASE_URL" -p 8080:8080 gochat-backend
```

### Configuration

The project uses Viper for configuration. Check the code for exact config keys and supported file formats (YAML/TOML). You can override with environment variables.

### Tests

Run tests with:

```bash
go test ./...
```

### Contributing

1. Fork the repository
2. Create a feature branch
3. Open a pull request describing your changes

### Contact

Email: kamonfay@gmail.com | fayamamoubarak@yahoo.com

GitHub: https://github.com/Badankamon

---
If you want I can commit and push this README for you, or update it to include precise run commands based on this repository's exact entrypoint. Tell me which you prefer.
