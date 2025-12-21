# Coding Rules

## When Adding New Features
1. **Define interface** in `internal/domain/`
2. **Create model** in `internal/model/`
3. **Implement service** in `internal/services/`
4. **Implement handler** in `internal/handler/`
5. **Register in** `internal/dependencies.go`
6. **Add route** in `internal/router/router.go`

## Error Handling
```go
if err != nil {
    return fmt.Errorf("context description: %w", err)
}
```

## Configuration
- Use struct tags: `env:"VAR_NAME,default=value"`
- All config in `internal/config/model.go`
- Load via `config.NewConfig()`

## Key Environment Variables
```
SERVER_PORT=8080
SERVER_SHUTDOWN_TIMEOUT=10s
ENV=development
OAUTH_ISSUER=http://localhost:5001
OAUTH_ACCESS_TOKEN_EXPIRATION=1h
```

## Code Style
- Use **tabs** for indentation
- Package names: lowercase, no underscores
- Meaningful variable names
- Small, focused functions
- Comment only when clarification needed
- Use Echo's context and error handling
- Follow Go conventions (gofmt, goimports)

## What to DO ✅
- Register ALL dependencies in `dependencies.go`
- Validate dependencies in constructors (nil checks)
- Use interfaces from `domain/` package
- Return errors from constructors
- Use Echo's context and error handling
- Follow Go conventions (gofmt, goimports)

## What NOT to DO ❌
- Instantiate dependencies directly in main.go
- Skip dependency registration in dig container
- Return concrete types from constructors
- Ignore errors
- Commit `.env` files or `*.db` files
