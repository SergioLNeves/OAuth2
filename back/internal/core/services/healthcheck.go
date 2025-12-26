package services

import (
	"github.com/SergioLNeves/OAuth2/back/internal/core/domain"
	"github.com/SergioLNeves/OAuth2/back/internal/core/ports"
)

const (
	WorkingStatus = "Working"
)

type HealthCheckServiceImpl struct {
}

func NewHealthCheckService() (ports.HealthCheckerService, error) {
	return HealthCheckServiceImpl{}, nil
}

func (h HealthCheckServiceImpl) Check() (domain.HealthCheck, []error) {
	healthCheck := domain.HealthCheck{Status: WorkingStatus}

	var errs []error

	return healthCheck, errs
}
