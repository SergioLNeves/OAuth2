package dependecies

import (
	"log"

	"github.com/SergioLNeves/OAuth2/back/internal/handler"
	"github.com/SergioLNeves/OAuth2/back/internal/router"
	"github.com/SergioLNeves/OAuth2/back/internal/services"
	"go.uber.org/dig"
)

func ProvideDependencies() *dig.Container {
	container := dig.New()

	if err := container.Provide(services.NewHealthCheckService); err != nil {
		log.Fatalf("failed to provide health check service: %v", err)
	}

	if err := container.Provide(handler.NewHealthCheckHandler); err != nil {
		log.Fatalf("failed to provide health check handler: %v", err)
	}

	if err := container.Provide(router.NewRouter); err != nil {
		log.Fatalf("failed to provide router: %v", err)
	}

	return container
}
