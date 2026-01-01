package route

import (
	serivcemarket "adaptor/api/v1/future/usdtm/market"
	serviceconfig "adaptor/config"
	handlermarket "adaptor/handler/future/usdm/market"
	inframarket "adaptor/infrastructure/future/usdm/market"
	"net/http"

	"github.com/labstack/echo/v4"
	bncaller "github.com/non26/tradepkg/pkg/bn/bn_caller"
	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bnrequest "github.com/non26/tradepkg/pkg/bn/bn_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/bn_response"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
)

func RouteFutureUsdmMarket(
	app *echo.Echo,
	config *serviceconfig.ServiceConfig,
) {
	httptransport := bntransport.NewBinanceTransport(&http.Transport{})
	httpclient := bnclient.NewBinanceSerivceHttpClient()

	bnMarketCaller := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[inframarket.KlinesCandleStickRequest](),
		bnresponse.NewBinanceServiceHttpResponse[inframarket.KlinesCandleStickResponse](),
		httptransport,
		httpclient,
		false,
		true,
		false,
	)
	_infraOrder := inframarket.NewBinanceFutureKlinesCandleStick(
		bnMarketCaller,
		config.Future.Usdm.Url,
		config.Future.Usdm.KlinesCandleStickEndpoint,
		&config.Secret)
	_serivceOrder := serivcemarket.NewMarket(_infraOrder)
	_handlerOrder := handlermarket.NewKlinesCandleStick(_serivceOrder)
	app.POST("/fapi/v1/klines", _handlerOrder.Handler)

}
