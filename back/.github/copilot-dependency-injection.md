# Dependency Injection Pattern

**ALL dependencies MUST be registered in `internal/dependencies.go`**

Flow: **Service → Handler → Router**

## Implementation Flow

```go
// 1. Define interface in domain/
type HealthCheckerService interface {
    Check() (model.HealthCheck, []error)
}

// 2. Implement in services/
func NewHealthCheckService() (domain.HealthCheckerService, error) {
    return &HealthCheckServiceImpl{}, nil
}

// 3. Create handler
func NewHealthCheckHandler(svc domain.HealthCheckerService) (domain.HealthCheckHandler, error) {
    if svc == nil {
        return nil, fmt.Errorf("service cannot be nil")
    }
    return &HealthCheckHandlerImpl{service: svc}, nil
}

// 4. Register in dependencies.go
container.Provide(services.NewHealthCheckService)
container.Provide(handler.NewHealthCheckHandler)
container.Provide(router.NewRouter)

// 5. Add route in router/router.go
func (r *Router) Setup(e *echo.Echo) {
    e.GET("/health", r.healthCheckHandler.Check)
}
```

## Constructor Pattern
Always return interface + error:
```go
func NewService() (domain.ServiceInterface, error) {
    return &ServiceImpl{}, nil
}

func NewHandler(svc domain.Service) (domain.Handler, error) {
    if svc == nil {
        return nil, fmt.Errorf("dependency cannot be nil")
    }
    return &HandlerImpl{service: svc}, nil
}
```

## Rules
- ✅ Register ALL dependencies in `dependencies.go`
- ✅ Validate dependencies in constructors (nil checks)
- ✅ Use interfaces from `domain/` package
- ✅ Return errors from constructors
- ❌ Don't instantiate dependencies directly in main.go
- ❌ Don't skip dependency registration in dig container
- ❌ Don't return concrete types from constructors
