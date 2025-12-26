package ports

import (
	"github.com/labstack/echo/v4"

	"github.com/SergioLNeves/OAuth2/back/internal/core/domain"
)

type HealthCheckerService interface {
	Check() (domain.HealthCheck, []error)
}

type HealthCheckHandler interface {
	Check(ctx echo.Context) error
}
