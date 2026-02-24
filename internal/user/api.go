package user

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/absurek/fem-go-aws/internal/response"
	"github.com/aws/aws-lambda-go/events"
)

var ErrEmptyUserData = errors.New("empty username or password")

type Api struct {
	service *Service
}

func NewApi(service *Service) *Api {
	return &Api{
		service: service,
	}
}

func (a *Api) RegisterUser(ctx context.Context, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var data UserData
	if err := json.Unmarshal([]byte(request.Body), &data); err != nil {
		return response.BadRequest()
	}

	if data.Username == "" || data.Password == "" {
		return response.BadRequest()
	}

	err := a.service.RegisterUser(ctx, data)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
			return response.BadRequest()
		}

		return response.InternalServerError()
	}

	return response.Created("user created")
}

func (a *Api) LoginUser(ctx context.Context, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var data UserData
	if err := json.Unmarshal([]byte(request.Body), &data); err != nil {
		return response.BadRequest()
	}

	if data.Username == "" || data.Password == "" {
		return response.BadRequest()
	}

	token, err := a.service.LoginUser(ctx, data)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			response.Unauthorized()
		}

		return response.InternalServerError()
	}

	return response.Ok(token)
}

func (a *Api) Protected(ctx context.Context, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	return response.Ok("This is a protected endpoint")
}
