package main

import (
	"net/http"

	"github.com/absurek/fem-go-aws/internal/application"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	app := application.New()

	// TODO(absurek): never return errors from handlers as they override the response object with a generic error
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			return app.UserApi.RegisterUser(request)
		case "/login":
			return app.UserApi.LoginUser(request)
		default:
			return events.APIGatewayProxyResponse{
				Body:       "Not Found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
}
