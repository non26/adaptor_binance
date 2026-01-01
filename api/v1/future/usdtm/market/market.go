package market

import (
	infrastructure "adaptor/infrastructure/future/usdm/market"
	"context"
)

type IMarket interface {
	KlinesCandleStick(ctx context.Context, request *infrastructure.KlinesCandleStickRequest) (*infrastructure.KlinesCandleStickResponse, error)
}

type market struct {
	bnKlinesCandleStick infrastructure.IBinanceFutureKlinesCandleStick
}

func NewMarket(
	bnKlinesCandleStick infrastructure.IBinanceFutureKlinesCandleStick,
) IMarket {
	return &market{
		bnKlinesCandleStick: bnKlinesCandleStick,
	}
}
