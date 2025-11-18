# Project Structure

## Overview
SaloonBook is now organized as a production-ready monorepo with clear separation between frontend and backend.

## Directory Structure

```
SaloonBook/
├── backend/                     # Go backend service
│   ├── cmd/                     # Application entrypoints
│   │   └── main.go              # Main application
│   ├── internal/                # Private application code
│   │   ├── config/              # Configuration management
│   │   │   └── config.go        # Environment config loader
│   │   ├── db/                  # Database layer
│   │   │   └── db.go            # PostgreSQL connection
│   │   ├── handlers/            # HTTP request handlers
│   │   │   ├── health.go        # Health check endpoint
│   │   │   ├── services.go      # Services endpoints
│   │   │   └── bookings.go      # Bookings endpoints
│   │   ├── middleware/          # HTTP middleware
│   │   │   └── cors.go          # CORS handling
│   │   ├── models/              # Data models/types
│   │   │   └── models.go        # Domain models
│   │   └── repository/          # Data access layer
│   │       ├── service.go       # Service repository
│   │       └── booking.go       # Booking repository
│   ├── pkg/                     # Public/shared libraries
│   │   └── logger/              # Structured logging
│   │       └── logger.go        # Logger implementation
│   ├── scripts/                 # Build/deployment scripts
│   ├── docs/                    # Backend documentation
│   ├── Dockerfile               # Container build
│   ├── go.mod                   # Go dependencies
│   ├── go.sum                   # Dependency checksums
│   └── README.md                # Backend docs
│
├── frontend/                    # Frontend applications
│   └── web/                     # React web application
│       ├── src/
│       │   ├── api/             # API client layer
│       │   │   ├── client.ts    # HTTP client
│       │   │   └── index.ts     # API methods
│       │   ├── components/      # React components
│       │   │   ├── ui/          # Reusable UI components
│       │   │   └── [feature]/   # Feature components
│       │   ├── hooks/           # Custom React hooks
│       │   ├── lib/             # Utilities/helpers
│       │   ├── pages/           # Page components
│       │   ├── types/           # TypeScript definitions
│       │   │   └── index.ts     # Shared types
│       │   ├── constants/       # Application constants
│       │   │   └── index.ts     # Const values
│       │   ├── utils/           # Utility functions
│       │   ├── App.tsx          # Root component
│       │   ├── main.tsx         # Entry point
│       │   ├── index.css        # Global styles
│       │   └── vite-env.d.ts    # Vite types
│       ├── public/              # Static assets
│       │   ├── assets/          # Images, fonts
│       │   └── favicon.png      # Favicon
│       ├── docs/                # Frontend documentation
│       ├── index.html           # HTML template
│       ├── package.json         # Dependencies
│       ├── tsconfig.json        # TypeScript config
│       ├── vite.config.ts       # Vite configuration
│       ├── tailwind.config.ts   # Tailwind CSS config
│       ├── postcss.config.cjs   # PostCSS config
│       ├── .env.development     # Dev environment vars
│       └── README.md            # Frontend docs
│
├── attached_assets/             # Design assets (images, etc.)
├── docker-compose.yml           # Docker orchestration
├── Makefile                     # Build automation
├── .env.example                 # Environment template
├── .dockerignore                # Docker ignore rules
└── README.md                    # Project documentation
```

## Architecture Layers

### Backend (Go)

#### 1. **cmd/** - Entry Points
- Contains application main packages
- Handles initialization and wiring
- Minimal business logic

#### 2. **internal/** - Private Code
Cannot be imported by external projects

- **config**: Environment variable loading and validation
- **db**: Database connection management  
- **handlers**: HTTP request/response handling (thin layer)
- **middleware**: Cross-cutting concerns (CORS, logging, auth)
- **models**: Domain models and DTOs
- **repository**: Data access abstraction

#### 3. **pkg/** - Public Libraries
Reusable across projects

- **logger**: Structured logging utilities

### Frontend (React + TypeScript)

#### 1. **api/** - Backend Communication
- HTTP client with type safety
- API method definitions
- Request/response handling

#### 2. **components/** - UI Building Blocks
- **ui/**: Reusable design system components
- Feature-specific components
- Layout components

#### 3. **types/** - TypeScript Definitions
- Shared interfaces
- API response types
- Domain models

#### 4. **constants/** - Application Constants
- Configuration values
- Enumerations
- Static data

## Design Principles

### Backend
1. **Clean Architecture**: Clear separation of layers
2. **Dependency Injection**: Testable, loosely coupled code
3. **Repository Pattern**: Abstract data access
4. **Single Responsibility**: Each package has one job
5. **Interface-Driven**: Program to interfaces

### Frontend
1. **Component Composition**: Build complex UIs from simple parts
2. **Type Safety**: Leverage TypeScript fully
3. **API Abstraction**: Centralized backend communication
4. **Separation of Concerns**: Logic separate from presentation

## File Naming Conventions

### Backend (Go)
- Files: `lowercase_snake.go` or `lowercase.go`
- Packages: `lowercase` (single word preferred)
- Types: `PascalCase`
- Functions/Methods: `PascalCase` (exported), `camelCase` (private)

### Frontend (TypeScript/React)
- Components: `PascalCase.tsx`
- Utilities: `camelCase.ts`
- Types: `camelCase.ts` or `index.ts`
- Constants: `UPPER_SNAKE_CASE` or `camelCase`

## Key Files

### Configuration
- `.env.example` - Environment variable template
- `backend/internal/config/config.go` - Backend config loader
- `frontend/web/.env.development` - Frontend dev config
- `frontend/web/vite.config.ts` - Build configuration

### Build & Deploy
- `Makefile` - Common development tasks
- `docker-compose.yml` - Local development environment
- `backend/Dockerfile` - Backend container build
- `backend/cmd/main.go` - Application entry point

### Documentation
- `README.md` - Project overview
- `backend/README.md` - Backend-specific docs
- `frontend/web/README.md` - Frontend-specific docs

## Development Workflow

1. **Start Database**: `docker compose up -d db`
2. **Start Backend**: `make backend` or `cd backend && go run ./cmd`
3. **Start Frontend**: `make frontend` or `cd frontend/web && npm run dev`
4. **Build All**: `make build`
5. **Clean**: `make clean`

## Testing Strategy

### Backend
- Unit tests: `*_test.go` files alongside source
- Integration tests: `internal/integration/`
- Run: `go test ./...`

### Frontend
- Component tests: `*.test.tsx`
- Type checking: `npm run typecheck`
- Run: `npm test`

## Next Steps for Production

1. **Add Database Migrations**: Use golang-migrate or similar
2. **Implement Authentication**: JWT or session-based auth
3. **Add Validation**: Request validation middleware
4. **Error Handling**: Standardized error responses
5. **Logging**: Structured logging throughout
6. **Monitoring**: Health checks, metrics, traces
7. **CI/CD**: Automated testing and deployment
8. **Documentation**: API docs (Swagger/OpenAPI)
9. **Security**: HTTPS, rate limiting, input sanitization
10. **Performance**: Caching, connection pooling, optimization

## Benefits of This Structure

✅ **Scalability**: Easy to add new features without conflicts  
✅ **Maintainability**: Clear where code belongs  
✅ **Testability**: Dependency injection and separation of concerns  
✅ **Team Collaboration**: Multiple developers can work independently  
✅ **Onboarding**: New developers understand structure quickly  
✅ **Flexibility**: Swap implementations without touching other layers  
✅ **Production-Ready**: Follows industry best practices
