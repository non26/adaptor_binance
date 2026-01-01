package market

import (
	infrastructure "adaptor/infrastructure/future/usdm/market"
	"context"
)

func (m *market) KlinesCandleStick(ctx context.Context, request *infrastructure.KlinesCandleStickRequest) (*infrastructure.KlinesCandleStickResponse, error) {
	response, err := m.bnKlinesCandleStick.KlinesCandleStick(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
