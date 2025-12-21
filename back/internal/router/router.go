package router

import (
	"github.com/SergioLNeves/OAuth2/back/internal/domain"
	"github.com/labstack/echo/v4"
)

type Router struct {
	healthCheckHandler domain.HealthCheckHandler
}

func NewRouter(healthCheckHandler domain.HealthCheckHandler) *Router {
	return &Router{
		healthCheckHandler: healthCheckHandler,
	}
}

func (r *Router) Setup(e *echo.Echo) {
	e.GET("/health", r.healthCheckHandler.Check)
}
