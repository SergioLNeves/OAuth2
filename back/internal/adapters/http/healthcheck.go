package http

import (
	"fmt"
	"net/http"

	"github.com/SergioLNeves/OAuth2/back/internal/core/ports"
	"github.com/labstack/echo/v4"
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
	check, err := h.healthCheckService.Check()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, check.Status)
}
