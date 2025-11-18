# SaloonBook Backend API

Go backend service for SaloonBook salon booking application.

## Architecture

```
backend/
├── cmd/
│   └── saloonbook/      # Application entrypoint
├── internal/
│   ├── config/          # Configuration management
│   ├── db/              # Database connection
│   ├── handlers/        # HTTP handlers
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data models
│   └── repository/      # Data access layer
├── pkg/
│   └── logger/          # Logging utilities
├── scripts/             # Helper scripts
└── docs/                # Documentation

```

## Key Features

- Clean architecture with separation of concerns
- Repository pattern for data access
- Structured logging
- Graceful shutdown
- CORS support
- Health checks
- Configuration management

## API Endpoints

- `GET /api/health` - Health check
- `GET /api/services` - List all services
- `GET /api/bookings` - List all bookings
- `POST /api/bookings` - Create a new booking

## Development

### Option 1: Using the dev script (recommended)
```bash
cd backend
./scripts/dev.sh
```

### Option 2: With environment variables
```bash
cd backend
DATABASE_URL=postgres://postgres:postgres@localhost:5433/saloonbook_dev \
PORT=5000 \
go run ./cmd
```

### Option 3: Using .env file
```bash
cd backend
# Ensure .env exists with correct values
go run ./cmd
```

### Option 4: From project root
```bash
make backend
```

## Environment Variables

See `.env.example` in project root.
