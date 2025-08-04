package main

import (
	serviceconfig "adaptor/config"
	route "adaptor/route/future"
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(middleware.BodyLimit("10M"))
	app.Use(middleware.Secure())
	app.Use(middleware.RequestID())

	config, err := serviceconfig.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	route.RouteFuture(app, config)
	echoLambda = echoadapter.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
