# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

CertSuite Overview is a dashboard tool that consolidates usage data from multiple sources to monitor the adoption and usage of CertSuite tools and services. It fetches data from:

1. **Quay Image Pulls**: Tracks container image pull statistics from the CertSuite repository on Quay.io
2. **DCI (Distributed CI) Test Runs**: Monitors CertSuite test executions performed by partners via Red Hat's DCI infrastructure
3. **CertSuite Collector Data**: Visualizes test metrics and execution statistics from the CertSuite Collector

All data is consolidated into a MySQL database, enabling real-time reporting and visualization through Grafana dashboards.

## Build Commands

```bash
# Build the project
make build

# Run go vet for static analysis
make vet

# Run golangci-lint for linting
make lint
```

### Manual Build

```bash
# Build from the cmd directory
cd cmd
go build -o ../certsuite-overview
```

## Running the CLI

The CLI uses Cobra for command handling. The main command is:

```bash
# Fetch data from Quay and DCI, store in database
./certsuite-overview fetch
```

## Test Commands

```bash
# Run unit tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

The project uses `go-sqlmock` for database mocking in tests.

## Code Organization

```
certsuite-overview/
├── cmd/                    # CLI application entry points
│   ├── main.go            # Root command and CLI initialization
│   └── certsuite_usage.go # Fetch command implementation
├── config/                 # Configuration management
│   └── config.go          # Viper-based config loading from environment
├── pkg/                    # Core business logic
│   ├── database.go        # MySQL database operations (local/AWS)
│   ├── database_test.go   # Database unit tests
│   ├── dci.go             # DCI data fetching logic
│   └── quay.go            # Quay image pull data fetching
├── grafana/               # Grafana configuration
│   ├── dashboard/         # Dashboard JSON definitions
│   └── datasource/        # MySQL datasource configuration
├── .github/workflows/     # GitHub Actions
│   ├── daily-sync.yaml    # Daily scheduled data sync job
│   └── pre-main.yaml      # PR validation (lint, vet)
├── go.mod                 # Go module definition
├── go.sum                 # Dependency checksums
└── Makefile               # Build automation
```

## Key Dependencies

| Package | Purpose |
|---------|---------|
| `github.com/spf13/cobra` | CLI framework |
| `github.com/spf13/viper` | Configuration management |
| `github.com/go-sql-driver/mysql` | MySQL database driver |
| `github.com/sebrandon1/go-quay` | Quay.io API client |
| `github.com/sebrandon1/go-dci` | DCI (Distributed CI) API client |
| `github.com/sirupsen/logrus` | Structured logging |
| `github.com/DATA-DOG/go-sqlmock` | SQL mocking for tests |
| `github.com/stretchr/testify` | Test assertions |

## Configuration

Configuration is loaded from environment variables using Viper:

| Variable | Description |
|----------|-------------|
| `DB_USER` | MySQL database username |
| `DB_PASSWORD` | MySQL database password |
| `DB_URL` | MySQL database host URL |
| `DB_PORT` | MySQL database port |
| `DB_CHOICE` | Database mode: "local" or "aws" |
| `CLIENTID` | DCI API client ID |
| `APISECRET` | DCI API secret |
| `BEARERTOKEN` | Quay.io API bearer token |
| `NAMESPACE` | Quay.io namespace (e.g., "redhat-best-practices-for-k8s") |
| `REPOSITORY` | Quay.io repository (e.g., "certsuite") |

## Database Schema

The application creates two tables in the `certsuite_usage_db` database:

### aggregated_logs
Stores Quay image pull data:
- `datetime` (DATE) - Date of the pull activity
- `count` (INT) - Number of pulls
- `kind` (VARCHAR) - Type of activity

### dci_components
Stores DCI test run data:
- `job_id` (VARCHAR) - DCI job identifier (primary key)
- `commit_hash` (VARCHAR) - Git commit hash
- `createdAt` (TIMESTAMP) - Job creation timestamp
- `totalSuccess` (INT) - Number of successful tests
- `totalFailures` (INT) - Number of failed tests
- `totalErrors` (INT) - Number of test errors
- `totalSkips` (INT) - Number of skipped tests

## Development Guidelines

### Go Version
This project uses Go 1.25.4. Ensure your local environment matches this version.

### Linting
The project uses `golangci-lint` for code quality. Run before committing:
```bash
make lint
```

### Testing
Write tests using the standard Go testing package with testify assertions. Use `go-sqlmock` for database-related tests.

### CI/CD Workflows

1. **PR Validation** (`pre-main.yaml`): Runs on pull requests to main branch
   - Executes `make vet`
   - Runs `golangci-lint`

2. **Daily Sync** (`daily-sync.yaml`): Scheduled job that runs daily
   - Builds the project
   - Executes `./certsuite-overview fetch` to sync data to AWS MySQL

### Local Development with MySQL

For local development, the application expects a MySQL server running on `localhost:3306` with:
- Username: `root`
- Password: `mypassword`

Set `DB_CHOICE=local` to use local database mode.
