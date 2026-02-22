package application

import (
	"context"

	"github.com/absurek/fem-go-aws/internal/platform/dynamo"
	"github.com/absurek/fem-go-aws/internal/user"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Application struct {
	UserApi *user.Api
}

func New() *Application {
	client := dynamo.NewClient(context.TODO())

	return &Application{
		UserApi: createUserApi(client),
	}
}

func createUserApi(client *dynamodb.Client) *user.Api {
	repository := user.NewDynamoReposiotry(client)
	service := user.NewService(repository)

	return user.NewApi(service)
}
