package main

import (
	"log"

	dependecies "github.com/SergioLNeves/OAuth2/back/internal"
	"github.com/SergioLNeves/OAuth2/back/internal/adapters/http"
	validator "github.com/SergioLNeves/OAuth2/back/internal/pkg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Validator = validator.NewValidator()

	container := dependecies.ProvideDependencies()
	err := container.Invoke(func(r *http.Router) {
		r.Setup(e)
	})
	if err != nil {
		log.Fatalf("failed to setup router: %v", err)
	}

	e.Logger.Fatal(e.Start(":8080").Error())
}
