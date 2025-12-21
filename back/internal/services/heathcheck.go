package services

import (
	"github.com/SergioLNeves/OAuth2/back/internal/domain"
	"github.com/SergioLNeves/OAuth2/back/internal/model"
)

const (
	WorkingStatus = "Working"
)

type HealthCheckServiceImpl struct {
}

func NewHealthCheckService() (domain.HealthCheckerService, error) {
	return HealthCheckServiceImpl{}, nil
}

func (h HealthCheckServiceImpl) Check() (model.HealthCheck, []error) {
	healthCheck := model.HealthCheck{Status: WorkingStatus}

	var errs []error

	return healthCheck, errs
}
