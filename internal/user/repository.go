package user

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const userTable = "UserTable"

var ErrUserNotFound = errors.New("user not found")

type Repository interface {
	UserExists(ctx context.Context, username string) (bool, error)
	InsertUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, username string) (User, error)
}

type DynamoRepository struct {
	client *dynamodb.Client
	table  string
}

func NewDynamoReposiotry(client *dynamodb.Client) *DynamoRepository {
	return &DynamoRepository{
		client: client,
		table:  userTable, // TODO(absurek): os.Getenv("USERS_TABLE")
	}
}

func (dr *DynamoRepository) UserExists(ctx context.Context, username string) (bool, error) {
	usernameav, err := attributevalue.Marshal(username)
	if err != nil {
		return true, err
	}

	result, err := dr.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(dr.table),
		Key:       map[string]types.AttributeValue{"username": usernameav},
	})
	if err != nil {
		return true, err
	}

	if result.Item == nil {
		return false, nil
	}

	return true, nil
}

func (dr *DynamoRepository) InsertUser(ctx context.Context, user User) error {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return err
	}

	_, err = dr.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(dr.table),
		Item:      item,
	})
	if err != nil {
		return err
	}

	return nil
}

func (dr *DynamoRepository) GetUser(ctx context.Context, username string) (User, error) {
	usernameav, err := attributevalue.Marshal(username)
	if err != nil {
		return User{}, err
	}

	result, err := dr.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(dr.table),
		Key:       map[string]types.AttributeValue{"username": usernameav},
	})
	if err != nil {
		return User{}, err
	}

	if result.Item == nil {
		return User{}, ErrUserNotFound
	}

	var user User
	if err := attributevalue.UnmarshalMap(result.Item, &user); err != nil {
		return User{}, err
	}

	return user, nil
}
