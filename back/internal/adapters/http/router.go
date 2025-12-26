package http

import (
	"github.com/SergioLNeves/OAuth2/back/internal/core/ports"
	"github.com/labstack/echo/v4"
)

type Router struct {
	healthCheckHandler ports.HealthCheckHandler
}

func NewRouter(healthCheckHandler ports.HealthCheckHandler) *Router {
	return &Router{
		healthCheckHandler: healthCheckHandler,
	}
}

func (r *Router) Setup(e *echo.Echo) {
	e.GET("/health", r.healthCheckHandler.Check)
}
