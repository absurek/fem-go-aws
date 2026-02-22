package dynamo

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewClient(ctx context.Context) *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.RetryMaxAttempts = 3
	})
}
