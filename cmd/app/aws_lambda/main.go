package main

import (
	serviceconfig "adaptor/config"
	routefuture "adaptor/route/future"
	routehealthcheck "adaptor/route/health_check"
	routelambda "adaptor/route/lambda"
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

	config, err := serviceconfig.ReadAWSAppConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(middleware.BodyLimit("10M"))
	app.Use(middleware.Secure())
	app.Use(middleware.RequestID())
	routehealthcheck.HealthCheck(app)
	routefuture.RouteFuture(app, config)
	routelambda.UpdateAWSAppConfig(app, config)

	echoLambda = echoadapter.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
