package ports

import (
	"github.com/SergioLNeves/OAuth2/back/internal/core/domain"
	"github.com/labstack/echo/v4"
)

type HealthCheckerService interface {
	Check() (domain.HealthCheck, []error)
}

type HealthCheckHandler interface {
	Check(ctx echo.Context) error
}
