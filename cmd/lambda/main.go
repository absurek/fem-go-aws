package main

import (
	"context"

	"github.com/absurek/fem-go-aws/internal/application"
	"github.com/absurek/fem-go-aws/internal/response"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	app := application.New()

	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			return app.UserApi.RegisterUser(context.Background(), request), nil
		case "/login":
			return app.UserApi.LoginUser(context.Background(), request), nil
		default:
			return response.NotFound(), nil
		}
	})
}
