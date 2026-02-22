package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

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

func (a *Api) RegisterUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data UserData
	if err := json.Unmarshal([]byte(request.Body), &data); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "bad request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	if data.Username == "" || data.Password == "" {
		return events.APIGatewayProxyResponse{
			Body:       "bad request",
			StatusCode: http.StatusBadRequest,
		}, ErrEmptyUserData
	}

	err := a.service.RegisterUser(context.TODO(), data)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
			return events.APIGatewayProxyResponse{
				Body:       "bad request",
				StatusCode: http.StatusBadRequest,
			}, err
		}

		return events.APIGatewayProxyResponse{
			Body:       "internal server error",
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "user created",
		StatusCode: http.StatusCreated,
	}, nil
}

func (a *Api) LoginUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data UserData
	if err := json.Unmarshal([]byte(request.Body), &data); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "bad request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	if data.Username == "" || data.Password == "" {
		return events.APIGatewayProxyResponse{
			Body:       "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}, ErrEmptyUserData
	}

	err := a.service.LoginUser(context.TODO(), data)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			return events.APIGatewayProxyResponse{
				Body:       "unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, err
		}

		return events.APIGatewayProxyResponse{
			Body:       "internal server error",
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "Successful login", // TODO(absurek): JWT
		StatusCode: http.StatusOK,
	}, nil
}
