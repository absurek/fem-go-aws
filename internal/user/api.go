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

func (a *Api) RegisterUser(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var data UserData
	if err := json.Unmarshal([]byte(request.Body), &data); err != nil {
		return response.BadRequest()
	}

	if data.Username == "" || data.Password == "" {
		return response.BadRequest()
	}

	err := a.service.RegisterUser(context.TODO(), data)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
			return response.BadRequest()
		}

		return response.InternalServerError()
	}

	return response.Created("user created")
}

func (a *Api) LoginUser(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var data UserData
	if err := json.Unmarshal([]byte(request.Body), &data); err != nil {
		return response.BadRequest()
	}

	if data.Username == "" || data.Password == "" {
		return response.BadRequest()
	}

	err := a.service.LoginUser(context.TODO(), data)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			response.Unauthorized()
		}

		return response.InternalServerError()
	}

	return response.Ok("login successful")
}
