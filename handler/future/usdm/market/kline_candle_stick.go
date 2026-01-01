package handler

import (
	api "adaptor/api/v1/future/usdtm/market"
	"adaptor/helper"
	infrastructure "adaptor/infrastructure/future/usdm/market"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KlinesCandleStickRequest struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

func (k *KlinesCandleStickRequest) ToInfrastructureRequest() *infrastructure.KlinesCandleStickRequest {
	return &infrastructure.KlinesCandleStickRequest{
		Symbol:    k.Symbol,
		Interval:  k.Interval,
		StartTime: k.StartTime,
		EndTime:   k.EndTime,
	}
}

type klinesCandleStick struct {
	service api.IMarket
}

func NewKlinesCandleStick(service api.IMarket) IHandler {
	return &klinesCandleStick{
		service: service,
	}
}

func (k *klinesCandleStick) Handler(c echo.Context) error {
	request := new(KlinesCandleStickRequest)
	if err := c.Bind(request); err != nil {
		response := helper.NewErrorResponse(err.Error()).ValidateRequestErrorCode()
		return c.JSON(http.StatusBadRequest, response)
	}

	response, err := k.service.KlinesCandleStick(c.Request().Context(), request.ToInfrastructureRequest())
	if err != nil {
		response := helper.NewErrorResponse(err.Error()).ValidateServiceErrorCode()
		return c.JSON(http.StatusInternalServerError, response)
	}
	return c.JSON(http.StatusOK, response)
}
