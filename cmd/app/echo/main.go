package main

import (
	serviceconfig "adaptor/config"
	route "adaptor/route/future"
	healthcheck "adaptor/route/health_check"
	routelambda "adaptor/route/lambda"
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

	healthcheck.HealthCheck(app, config.HealthcheckMessage)
	route.RouteFuture(app, config)
	routelambda.UpdateAWSAppConfig(app, config)

	app.Start(fmt.Sprintf(":%d", config.Port))
}
