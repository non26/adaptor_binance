package handler

import (
	api "adaptor/api/v1/future/usdtm/trade"
	"adaptor/helper"
	infrastructure "adaptor/infrastructure/future/usdm/trade"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MakeOrderRequest struct {
	AccountId        string `json:"accountId" binding:"required"`
	PositionSide     string `json:"positionSide" binding:"required"`
	Side             string `json:"side" binding:"required"`
	Quantity         string `json:"quantity" binding:"required"`
	Symbol           string `json:"symbol" binding:"required"`
	NewClientOrderId string `json:"newClientOrderId" binding:"required"`
	Type             string `json:"type" binding:"required"`
}

func (n *MakeOrderRequest) ToBinanceNewOrder() *infrastructure.BinanceNewOrderRequest {
	return &infrastructure.BinanceNewOrderRequest{
		Symbol:           n.Symbol,
		Side:             n.Side,
		PositionSide:     n.PositionSide,
		Quantity:         n.Quantity,
		Type:             n.Type,
		NewClientOrderId: n.NewClientOrderId,
	}
}

type makeOrder struct {
	service api.IOrder
}

func NewMakeOrder(service api.IOrder) IHandler {
	return &makeOrder{
		service: service,
	}
}

func (n *makeOrder) Handler(c echo.Context) error {
	request := new(MakeOrderRequest)
	if err := c.Bind(request); err != nil {
		response := helper.NewErrorResponse(err.Error()).ValidateRequestErrorCode()
		return c.JSON(http.StatusBadRequest, response)
	}

	response, err := n.service.MakeOrder(c.Request().Context(), request.AccountId, request.ToBinanceNewOrder())
	if err != nil {
		response := helper.NewErrorResponse(err.Error()).ValidateServiceErrorCode()
		return c.JSON(http.StatusInternalServerError, response)
	}
	return c.JSON(http.StatusOK, response)
}
