package user

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExists  = errors.New("user alredy exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) RegisterUser(ctx context.Context, data UserData) error {
	exists, err := s.repository.UserExists(ctx, data.Username)
	if err != nil {
		return fmt.Errorf("checking if user exists (username=%s): %w", data.Username, err)
	}

	if exists {
		return ErrUserAlreadyExists
	}

	user, err := newUser(data)
	if err != nil {
		return fmt.Errorf("creating user (username=%s): %w", data.Username, err)
	}

	if err := s.repository.InsertUser(ctx, user); err != nil {
		return fmt.Errorf("inserting user (username=%s): %w", data.Username, err)
	}

	return nil
}

// TODO(absurek): create and return JWT
func (s *Service) LoginUser(ctx context.Context, data UserData) error {
	user, err := s.repository.GetUser(ctx, data.Username)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return ErrInvalidCredentials
		}

		return fmt.Errorf("getting user (username=%s): %w", data.Username, err)
	}

	if !validatePassword(user.PasswordHash, data.Password) {
		return ErrInvalidCredentials
	}

	return nil
}

func newUser(data UserData) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return User{}, err
	}

	return User{
		Username:     data.Username,
		PasswordHash: string(passwordHash),
	}, nil
}

func validatePassword(passwordHash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}
