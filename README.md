<div align="center">

# News Platform — Backend API

**A Go-based REST API backed by PostgreSQL**

[![Go](https://img.shields.io/badge/Go-1.27-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Echo](https://img.shields.io/badge/HTTP-Echo_v5-1F6FEB)](https://echo.labstack.com/)
[![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL_16-4169E1?logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Container-Docker-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)

[Repository](https://github.com/mohammad-valipour/news-api) · [Frontend](https://github.com/mohammad-valipour/news-frontend)

</div>

---

## Overview

`news-api` is the backend service for a news publishing platform. It is written in **Go** and uses **Echo v5** for HTTP routing and **PostgreSQL** for persistent data storage.

The codebase separates responsibilities into packages for configuration, database access, controllers, routing, services, repositories, models, and database migrations. The repository also includes Swagger-related dependencies and Docker assets for running the API with PostgreSQL.

> **Project status:** In active development. API behavior, database schema, routes, configuration, and deployment details may change. The project is not presented as a finished production system.

## Technology stack

| Area | Technology |
|---|---|
| Language | Go 1.27 (as declared in `go.mod`) |
| HTTP framework | Echo v5 |
| Database | PostgreSQL 16 (Docker Compose image: `postgres:16-alpine`) |
| PostgreSQL driver | `pgx/v5` |
| Environment configuration | `godotenv` |
| API documentation tooling | `swaggo/swag`, `echo-swagger` |
| Containerization | Docker |
| Local orchestration | Docker Compose |

## Architecture and project structure

The current repository contains the following top-level areas:

```text
.
├── .github/
│   └── workflows/     # GitHub workflow configuration
├── config/            # Application configuration
├── controllers/       # HTTP/controller layer
├── database/          # Database connection and related code
├── docker/            # Dockerfile and container-related assets
├── docs/              # API documentation artifacts
├── migrations/        # Database migration files
├── models/            # Domain/data models
├── repository/        # Data access layer
├── routing/           # HTTP route registration
├── services/          # Application/business logic
├── view/              # View or response-related code
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

The intended responsibility boundaries can be understood as:

- **Routing:** registers HTTP endpoints and connects them to controllers.
- **Controllers:** handle HTTP-level input and output.
- **Services:** contain application operations and business rules.
- **Repositories:** isolate data-access operations.
- **Models:** represent application data.
- **Database and migrations:** support database connectivity and schema evolution.
- **Config:** centralizes application configuration.

Refer to the source code for the exact implementation and dependencies between packages.

## Running with Docker Compose

Docker Compose is the recommended way to start the API together with its PostgreSQL dependency in a local development environment.

### Prerequisites

- Docker Engine
- Docker Compose plugin (`docker compose`)
- Git

### 1. Clone the repository

```bash
git clone https://github.com/mohammad-valipour/news-api.git
cd news-api
```

### 2. Create the environment file

The current Compose configuration references `APP_PORT`, `DB_USER`, `DB_PASSWORD`, and `DB_NAME`. Create a local `.env` file in the project root:

```env
APP_PORT=8080

DB_USER=news_user
DB_PASSWORD=change_this_local_password
DB_NAME=news
```

Use values suitable for your local environment. Check the application's configuration package for any additional variables required by the current version.

> Keep `.env` out of version control. Do not use example credentials in a public or production environment.

### 3. Build and start the services

```bash
docker compose up --build
```

Compose starts:

| Service | Container | Purpose |
|---|---|---|
| `db` | `news-db` | PostgreSQL database |
| `api` | `news-api` | Go HTTP API |

The Compose file configures PostgreSQL with a named volume (`postgres_data`) so database data can persist across container recreation. It also defines a database health check and makes the API service wait until the database reports healthy.

### 4. Follow service logs

```bash
docker compose logs -f
```

To inspect one service:

```bash
docker compose logs -f api
docker compose logs -f db
```

### 5. Stop the services

```bash
docker compose down
```

This stops and removes the containers and network while preserving the named database volume.

To also remove the local database volume and its data:

```bash
docker compose down -v
```

**Warning:** `-v` deletes the persisted PostgreSQL data for this Compose project.

## Running without Docker

For direct local development, install a compatible Go toolchain and PostgreSQL instance.

### 1. Start PostgreSQL

Create a database and user matching the values used by your local environment configuration.

### 2. Configure environment variables

Create `.env` in the project root. At minimum, the Compose configuration documents these names:

```env
APP_PORT=8080
DB_USER=news_user
DB_PASSWORD=change_this_local_password
DB_NAME=news
```

The application may require additional configuration. Verify `config/` and the database initialization code before running it outside Compose.

### 3. Download Go dependencies

```bash
go mod download
```

### 4. Run the application

```bash
go run .
```

The exact startup behavior and listening address are controlled by the current application configuration. Check the startup logs and configuration code for the effective address.

## API documentation

The project includes Swagger tooling through:

- `github.com/swaggo/swag`
- `github.com/swaggo/echo-swagger/v2`

The generated documentation artifacts are kept in the repository's `docs/` directory. The exact Swagger UI route and generation command depend on the current routing and documentation setup; consult `routing/`, `main.go`, and the documentation files rather than assuming a URL.

When API routes or request/response structures change, keep the generated or maintained API documentation synchronized with the implementation.

## Database and migrations

PostgreSQL is used as the relational database. Migration files are organized under `migrations/`.

Before applying schema changes:

1. Review the existing migration sequence.
2. Create a new migration using the project's established naming and format.
3. Test the migration against a development database.
4. Verify the application models and repository queries remain consistent.
5. Never test destructive migrations against important data.

The exact migration tool and commands should be taken from the current migration files and project scripts.

## Development conventions

Recommended practices for contributing to this codebase:

- Keep HTTP concerns in the controller/routing layer.
- Keep business operations in services.
- Keep SQL and persistence concerns within the repository/database layer.
- Return consistent HTTP status codes and response structures.
- Validate incoming data before using it in application operations.
- Handle errors explicitly; avoid silently discarding failures.
- Use parameterized database queries.
- Keep configuration outside source code.
- Add or update tests when changing behavior.
- Update API documentation when endpoint contracts change.
- Run formatting and static checks before submitting changes.

## Useful commands

```bash
# Download dependencies
go mod download

# Verify module dependencies
go mod verify

# Run the application
go run .

# Run all Go tests
go test ./...

# Format Go source files
gofmt -w .

# Build the application
go build ./...

# Start API and PostgreSQL
docker compose up --build

# Stop containers (preserve database volume)
docker compose down

# View logs
docker compose logs -f
```

Some commands may need to be adjusted as scripts and tooling evolve.

## Frontend integration

The companion frontend repository is:

[**news-frontend — Next.js application**](https://github.com/mohammad-valipour/news-frontend)

The frontend and backend are maintained independently. To integrate them locally:

1. Start PostgreSQL and the API.
2. Confirm the API's actual listening address and route prefixes from the running application.
3. Configure the frontend's API service/base URL to point to that address.
4. Verify browser-origin requirements and CORS behavior if the frontend and API use different origins.
5. Confirm the frontend request/response types match the API contract.

Do not assume endpoint paths, authentication behavior, or payload formats without checking the current implementation or Swagger documentation.

## Security notes

- Never commit `.env` files, passwords, or private tokens.
- Use strong, unique credentials outside local development.
- Do not expose PostgreSQL publicly unless the deployment explicitly requires it and access is restricted.
- Validate and normalize incoming request data.
- Use parameterized SQL queries.
- Avoid returning internal error details or secrets in API responses.
- Review CORS and allowed origins before deployment.
- Use HTTPS and an appropriate secret-management mechanism in deployed environments.

## Current scope and limitations

This repository is under active development. Its structure and dependencies are visible in the source, but this README intentionally does not claim that every planned endpoint, authentication flow, authorization policy, test suite, or production-hardening task is complete.

For the authoritative current behavior, inspect the implementation and the generated API documentation.

## Roadmap

Potential areas for continued development include:

- Expanding and refining API capabilities.
- Improving request validation and consistent error handling.
- Increasing automated test coverage.
- Maintaining database migrations and API documentation.
- Strengthening observability and operational configuration.
- Improving deployment and production-readiness practices.

## Related repository

- [Frontend application (Next.js)](https://github.com/mohammad-valipour/news-frontend)

## Maintainer

**Mohammad Valipour**  
GitHub: [@mohammad-valipour](https://github.com/mohammad-valipour)

---

<div align="center">

**Work in progress · Built with Go, Echo, PostgreSQL, and Docker**

</div>
