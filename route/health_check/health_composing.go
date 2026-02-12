package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
	healthcheckpkg "github.com/non26/tradepkg/pkg/bn/health_check"
)

func HealthCheck(app *echo.Echo, msg string) {
	app.GET(healthcheckpkg.PATH_HEALTHCHECK, func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			healthcheckpkg.NewHealthCheckResponseWith(msg),
		)
	})
}
