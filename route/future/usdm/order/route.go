package route

import (
	serivceorder "adaptor/api/v1/future/usdtm/order"
	serviceconfig "adaptor/config"
	handlerorder "adaptor/handler/future/usdm/order"
	infraorder "adaptor/infrastructure/future/usdm/order"
	"net/http"

	"github.com/labstack/echo/v4"
	bncaller "github.com/non26/tradepkg/pkg/bn/bn_caller"
	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bnrequest "github.com/non26/tradepkg/pkg/bn/bn_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/bn_response"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
)

func RouteFutureUsdmOrder(
	app *echo.Echo,
	config *serviceconfig.ServiceConfig,
) {
	httptransport := bntransport.NewBinanceTransport(&http.Transport{})
	httpclient := bnclient.NewBinanceSerivceHttpClient()

	bnOrderCaller := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[infraorder.BinanceNewOrderRequest](),
		bnresponse.NewBinanceServiceHttpResponse[infraorder.BinanceNewOrderResponse](),
		httptransport,
		httpclient,
		true,
		true,
		false,
	)
	_infraOrder := infraorder.NewBinanceFutureNewOrder(bnOrderCaller, config.Future.Usdm.Url, config.Future.Usdm.NewOrderEndpoint, &config.Secrets)
	_serivceOrder := serivceorder.NewOrder(_infraOrder)
	_handlerOrder := handlerorder.NewMakeOrder(_serivceOrder)
	app.POST("/fapi/v1/order", _handlerOrder.Handler)

}
