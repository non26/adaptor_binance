package api

import (
	infrastructure "adaptor/infrastructure/future/usdm/order"
	"context"
)

func (o *order) MakeOrder(ctx context.Context, accountid string, request *infrastructure.BinanceNewOrderRequest) (*infrastructure.BinanceNewOrderResponse, error) {
	response, err := o.bnOrder.NewOrder(ctx, accountid, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
