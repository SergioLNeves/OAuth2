package main

import (
	"log"

	"github.com/SergioLNeves/OAuth2/back/internal/config"
	dependecies "github.com/SergioLNeves/OAuth2/back/internal"
	validator "github.com/SergioLNeves/OAuth2/back/internal/pkg"
	"github.com/SergioLNeves/OAuth2/back/internal/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	_ = validator.GetValidator()

	container := dependecies.ProvideDependencies()

	if err := container.Invoke(func(r *router.Router) {
		r.Setup(e)
	}); err != nil {
		log.Fatalf("failed to setup router: %v", err)
	}

	server := config.NewAPI(e, cfg.Server.Port, cfg.Server.ShutdownTimeout)
	server.Start()
}
