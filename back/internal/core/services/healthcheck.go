package services

import (
	"github.com/SergioLNeves/OAuth2/back/internal/core/domain"
	"github.com/SergioLNeves/OAuth2/back/internal/core/ports"
)

const (
	WorkingStatus = "Working"
	FailedStatus  = "Failed"
)

type HealthCheckServiceImpl struct {
	db ports.Database
}

func NewHealthCheckService(db ports.Database) (ports.HealthCheckerService, error) {
	return &HealthCheckServiceImpl{db: db}, nil
}

func (h *HealthCheckServiceImpl) Check() (domain.HealthCheck, []error) {
	var errs []error

	dbHealth := h.checkDatabase()
	if dbHealth.Status == FailedStatus {
		errs = append(errs, nil)
	}

	status := WorkingStatus
	if len(errs) > 0 {
		status = FailedStatus
	}

	healthCheck := domain.HealthCheck{
		Status:   status,
		Database: dbHealth,
	}

	return healthCheck, errs
}

func (h *HealthCheckServiceImpl) checkDatabase() domain.DatabaseHealthCheck {
	if h.db == nil {
		return domain.DatabaseHealthCheck{
			Status: FailedStatus,
			Error:  "database not initialized",
		}
	}

	if err := h.db.Ping(); err != nil {
		return domain.DatabaseHealthCheck{
			Status: FailedStatus,
			Error:  err.Error(),
		}
	}

	return domain.DatabaseHealthCheck{
		Status: WorkingStatus,
	}
}
