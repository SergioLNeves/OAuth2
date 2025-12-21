package main

import (
	"log"

	dependecies "github.com/SergioLNeves/OAuth2/back/internal"
	validator "github.com/SergioLNeves/OAuth2/back/internal/pkg"
	"github.com/SergioLNeves/OAuth2/back/internal/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	_ = validator.GetValidator()

	container := dependecies.ProvideDependencies()
	err := container.Invoke(func(r *router.Router) {
		r.Setup(e)
	})
	if err != nil {
		log.Fatalf("failed to setup router: %v", err)
	}

	e.Logger.Fatal(e.Start(":8080").Error())
}
