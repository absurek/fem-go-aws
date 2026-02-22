package user

import (
	"context"
	"fmt"
)

type Api struct {
	service *Service
}

func NewApi(service *Service) *Api {
	return &Api{
		service: service,
	}
}

func (a *Api) HandleRegisterUser(event RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("invalid username or password")
	}

	return a.service.RegisterUser(context.TODO(), event)
}
