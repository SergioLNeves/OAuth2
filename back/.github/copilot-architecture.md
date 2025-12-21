# Architecture - OAuth 2.0 Backend

## Stack
- **Go 1.25.4**
- **Echo v4** - Web framework
- **Uber Dig** - Dependency injection
- **go-env** - Environment management
- **go-playground/validator** - Request validation

## Project Structure
```
cmd/app/main.go              # Entry point
internal/
├── config/                  # Configuration & API server
│   ├── api.go              # Graceful shutdown server
│   ├── config.go           # Config loader
│   └── model.go            # Config structs
├── dependencies.go         # DI container (Uber Dig)
├── domain/                 # Interfaces
├── handler/                # HTTP handlers
├── model/                  # Data models
├── pkg/                    # Utilities
├── router/                 # Route aggregator
└── services/               # Business logic
```

## Module Path
`github.com/SergioLNeves/OAuth2/back`
