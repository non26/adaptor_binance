package infrastructure

import (
	serviceconfig "adaptor/config"
	"adaptor/helper"
	"context"
	"net/http"

	bncaller "github.com/non26/tradepkg/pkg/bn/bn_caller"
)

type BinanceNewOrderRequest struct {
	Symbol           string `json:"symbol"`
	Side             string `json:"side"`
	PositionSide     string `json:"positionSide"`
	Quantity         string `json:"quantity"`
	Type             string `json:"type"`
	NewClientOrderId string `json:"newClientOrderId"`
	Timestamp        string `json:"timestamp"`
}

func (n *BinanceNewOrderRequest) PrepareRequest() {
	n.Timestamp = helper.GetTimestamp()
}

func (n *BinanceNewOrderRequest) GetData() interface{} {
	return n
}

// for binance response when calling new order endpoint
// if success binance will return the information of the order
// if error binance will return the error code and message
type BinanceNewOrderResponse struct {
	// for error
	Code    *int    `json:"code"`
	Message *string `json:"msg"`
	// for success
	ClientOrderID           *string `json:"clientOrderId"`
	CumQty                  *string `json:"cumQty"`
	CumQuote                *string `json:"cumQuote"`
	ExecutedQty             *string `json:"executedQty"`
	OrderID                 *int    `json:"orderId"`
	AvgPrice                *string `json:"avgPrice"`
	OrigQty                 *string `json:"origQty"`
	Price                   *string `json:"price"`
	ReduceOnly              *bool   `json:"reduceOnly"`
	Side                    *string `json:"side"`
	PositionSide            *string `json:"positionSide"`
	Status                  *string `json:"status"`
	StopPrice               *string `json:"stopPrice"`
	ClosePosition           *bool   `json:"closePosition"`
	Symbol                  *string `json:"symbol"`
	TimeInForce             *string `json:"timeInForce"`
	Type                    *string `json:"type"`
	OrigType                *string `json:"origType"`
	ActivatePrice           *string `json:"activatePrice"`
	PriceRate               *string `json:"priceRate"`
	UpdateTime              *int64  `json:"updateTime"`
	WorkingType             *string `json:"workingType"`
	PriceProtect            *bool   `json:"priceProtect"`
	PriceMatch              *string `json:"priceMatch"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode"`
	GoodTillDate            *int64  `json:"goodTillDate"`
}

type IBinanceFutureNewOrder interface {
	NewOrder(
		ctx context.Context,
		accountId string,
		request *BinanceNewOrderRequest,
	) (*BinanceNewOrderResponse, error)
}

type binanceFutureNewOrder struct {
	service          bncaller.ICallBinance[BinanceNewOrderRequest, BinanceNewOrderResponse]
	baseUrl          string
	newOrderEndPoint string
	secret           *serviceconfig.Secrets
}

func NewBinanceFutureNewOrder(
	service bncaller.ICallBinance[BinanceNewOrderRequest, BinanceNewOrderResponse],
	baseUrl string,
	newOrderEndPoint string,
	secret *serviceconfig.Secrets,
) IBinanceFutureNewOrder {
	return &binanceFutureNewOrder{
		service:          service,
		baseUrl:          baseUrl,
		newOrderEndPoint: newOrderEndPoint,
		secret:           secret,
	}
}

func (b *binanceFutureNewOrder) NewOrder(
	ctx context.Context,
	accountId string,
	request *BinanceNewOrderRequest,
) (*BinanceNewOrderResponse, error) {
	// select account
	apikey, secretkey, err := helper.SelectAccount(accountId, b.secret)
	if err != nil {
		return nil, err
	}
	respone, err := b.service.CallBinance(
		request,
		b.baseUrl,
		b.newOrderEndPoint,
		http.MethodPost,
		secretkey,
		apikey,
	)
	if err != nil {
		return nil, err
	}

	return respone, nil
}
