package dependecies

import (
	"log"

	"go.uber.org/dig"

	"github.com/SergioLNeves/OAuth2/back/internal/adapters/config"
	"github.com/SergioLNeves/OAuth2/back/internal/adapters/http"
	"github.com/SergioLNeves/OAuth2/back/internal/core/services"
	"github.com/SergioLNeves/OAuth2/back/internal/infrastructure/database"
)

func ProvideDependencies() *dig.Container {
	container := dig.New()

	if err := container.Provide(config.NewConfig); err != nil {
		log.Fatalf("failed to provide config: %v", err)
	}

	if err := container.Provide(NewDatabaseConfig); err != nil {
		log.Fatalf("failed to provide database config: %v", err)
	}

	if err := container.Provide(database.NewDatabase); err != nil {
		log.Fatalf("failed to provide database: %v", err)
	}

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

func NewDatabaseConfig(cfg *config.Config) *database.Config {
	return &database.Config{
		DBPath:      cfg.AuthDB.DBPath,
		Environment: cfg.Env,
		MaxConn:     cfg.SQLite.MaxConn,
		MaxIdle:     cfg.SQLite.MaxIdle,
		MaxLifeTime: cfg.SQLite.MaxLifeTime,
	}
}
