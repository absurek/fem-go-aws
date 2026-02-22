package user

import (
	"context"
	"errors"
	"fmt"
)

var ErrUserAlreadyExists = errors.New("user alredy exists")

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) RegisterUser(ctx context.Context, user RegisterUser) error {
	exists, err := s.repository.UserExists(ctx, user.Username)
	if err != nil {
		return fmt.Errorf("checking if user exists (username=%s): %w", user.Username, err)
	}

	if exists {
		return ErrUserAlreadyExists
	}

	if err := s.repository.InsertUser(ctx, user); err != nil {
		return fmt.Errorf("inserting user (username=%s): %w", user.Username, err)
	}

	return nil
}
