package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/SergioLNeves/OAuth2/back/internal/core/ports"
)

type HealthCheckHandlerImpl struct {
	healthCheckService ports.HealthCheckerService
}

func NewHealthCheckHandler(healthCheckService ports.HealthCheckerService) (ports.HealthCheckHandler, error) {
	if healthCheckService == nil {
		return nil, fmt.Errorf("Failed to initialize health check service dependency")
	}
	return &HealthCheckHandlerImpl{
		healthCheckService: healthCheckService,
	}, nil
}

func (h HealthCheckHandlerImpl) Check(ctx echo.Context) error {
	check, errs := h.healthCheckService.Check()

	statusCode := http.StatusOK
	if len(errs) > 0 {
		statusCode = http.StatusServiceUnavailable
	}

	return ctx.JSON(statusCode, check)
}
