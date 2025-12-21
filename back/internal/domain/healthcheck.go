package domain

import (
	"github.com/SergioLNeves/OAuth2/back/internal/model"
	"github.com/labstack/echo/v4"
)

type HealthCheckerService interface {
	Check() (model.HealthCheck, []error)
}

type HealthCheckHandler interface {
	Check(ctx echo.Context) error
}
