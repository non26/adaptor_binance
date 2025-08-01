package api

import (
	infrastructure "adaptor/infrastructure/future/usdm/order"
	"context"
)

type order struct {
	bnOrder infrastructure.IBinanceFutureNewOrder
}

type IOrder interface {
	MakeOrder(ctx context.Context, accountid string, request *infrastructure.BinanceNewOrderRequest) (*infrastructure.BinanceNewOrderResponse, error)
}

func NewOrder(
	bnOrder infrastructure.IBinanceFutureNewOrder,
) IOrder {
	return &order{
		bnOrder: bnOrder,
	}
}
