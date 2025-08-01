package route

import (
	serviceconfig "adaptor/config"
	route1 "adaptor/route/future/usdm"

	"github.com/labstack/echo/v4"
)

func RouteFuture(
	app *echo.Echo,
	config *serviceconfig.ServiceConfig,
) {
	route1.RouteFutureUsdm(app, config)
}
