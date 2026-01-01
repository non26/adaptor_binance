package handler

import "github.com/labstack/echo/v4"

type IHandler interface {
	Handler(c echo.Context) error
}
