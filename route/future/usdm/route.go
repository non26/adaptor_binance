package route

import (
	serviceconfig "adaptor/config"
	route1 "adaptor/route/future/usdm/trade"

	"github.com/labstack/echo/v4"
)

func RouteFutureUsdm(
	app *echo.Echo,
	config *serviceconfig.ServiceConfig,
) {
	route1.RouteFutureUsdmOrder(app, config)
}
