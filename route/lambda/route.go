package route

import (
	serviceconfig "adaptor/config"

	"github.com/labstack/echo/v4"
)

func RouteLambda(
	app *echo.Echo, config *serviceconfig.ServiceConfig,
) {
	UpdateAWSAppConfig(app, config)
}
