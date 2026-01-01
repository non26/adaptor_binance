package infrastructure

import (
	serviceconfig "adaptor/config"
	"adaptor/helper"
	"context"
	"net/http"

	bncaller "github.com/non26/tradepkg/pkg/bn/bn_caller"
)

type KlinesCandleStickRequest struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime string `json:"startTime" binance:"optional"`
	EndTime   string `json:"endTime" binance:"optional"`
}

func (n *KlinesCandleStickRequest) PrepareRequest() {
}

func (n *KlinesCandleStickRequest) GetData() interface{} {
	return n
}

// type KlinesCandleStickResponse struct {
// 	// for error
// 	// Code    *int    `json:"code"`
// 	// Message *string `json:"msg"`
// 	// for success
// 	Klines [][]interface{}
// }

type KlinesCandleStickResponse [][]interface{}

type IBinanceFutureKlinesCandleStick interface {
	KlinesCandleStick(
		ctx context.Context,
		request *KlinesCandleStickRequest,
	) (*KlinesCandleStickResponse, error)
}

type binanceFutureKlinesCandleStick struct {
	service                   bncaller.ICallBinance[KlinesCandleStickRequest, KlinesCandleStickResponse]
	baseUrl                   string
	klinesCandleStickEndPoint string
	secret                    *serviceconfig.Secrets
}

func NewBinanceFutureKlinesCandleStick(
	service bncaller.ICallBinance[KlinesCandleStickRequest, KlinesCandleStickResponse],
	baseUrl string,
	klinesCandleStickEndPoint string,
	secret *serviceconfig.Secrets,
) IBinanceFutureKlinesCandleStick {
	return &binanceFutureKlinesCandleStick{
		service:                   service,
		baseUrl:                   baseUrl,
		klinesCandleStickEndPoint: klinesCandleStickEndPoint,
		secret:                    secret,
	}
}

func (b *binanceFutureKlinesCandleStick) KlinesCandleStick(
	ctx context.Context,
	request *KlinesCandleStickRequest,
) (*KlinesCandleStickResponse, error) {
	// select account
	apikey, secretkey, err := helper.SelectAccount("1", b.secret)
	if err != nil {
		return nil, err
	}

	respone, err := b.service.CallBinance(
		request,
		b.baseUrl,
		b.klinesCandleStickEndPoint,
		http.MethodGet,
		secretkey,
		apikey,
	)
	if err != nil {
		return nil, err
	}

	return respone, nil
}
