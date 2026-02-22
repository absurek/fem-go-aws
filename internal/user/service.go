package user

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

func (s *Service) RegisterUser(ctx context.Context, reg RegisterUser) error {
	exists, err := s.repository.UserExists(ctx, reg.Username)
	if err != nil {
		return fmt.Errorf("checking if user exists (username=%s): %w", reg.Username, err)
	}

	if exists {
		return ErrUserAlreadyExists
	}

	user, err := newUser(reg)
	if err != nil {
		return fmt.Errorf("creating user (username=%s): %w", reg.Username, err)
	}

	if err := s.repository.InsertUser(ctx, user); err != nil {
		return fmt.Errorf("inserting user (username=%s): %w", reg.Username, err)
	}

	return nil
}

func newUser(reg RegisterUser) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(reg.Password), 10)
	if err != nil {
		return User{}, err
	}

	return User{
		Username:     reg.Username,
		PasswordHash: string(passwordHash),
	}, nil
}

func validatePassword(hashedPassword, plainTextPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword)) == nil
}
