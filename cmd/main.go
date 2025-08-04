package main

import (
	serviceconfig "adaptor/config"
	route "adaptor/route/future"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(middleware.BodyLimit("10M"))
	app.Use(middleware.Secure())
	app.Use(middleware.RequestID())

	config, err := serviceconfig.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	route.RouteFuture(app, config)

	app.Start(fmt.Sprintf(":%d", config.Port))
}
