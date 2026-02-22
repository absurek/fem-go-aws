package main

import (
	"github.com/absurek/fem-go-aws/internal/application"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	app := application.New()
	lambda.Start(app.UserApi.HandleRegisterUser)
}
