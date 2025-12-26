package domain

type HealthCheck struct {
	Status   string              `json:"status"`
	Database DatabaseHealthCheck `json:"database"`
}

type DatabaseHealthCheck struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}
