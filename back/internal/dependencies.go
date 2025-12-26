package dependecies

import (
	"log"

	"github.com/SergioLNeves/OAuth2/back/internal/adapters/http"
	"github.com/SergioLNeves/OAuth2/back/internal/core/services"
	"go.uber.org/dig"
)

func ProvideDependencies() *dig.Container {
	container := dig.New()

	if err := container.Provide(services.NewHealthCheckService); err != nil {
		log.Fatalf("failed to provide health check service: %v", err)
	}

	if err := container.Provide(http.NewHealthCheckHandler); err != nil {
		log.Fatalf("failed to provide health check handler: %v", err)
	}

	if err := container.Provide(http.NewRouter); err != nil {
		log.Fatalf("failed to provide router: %v", err)
	}

	return container
}
