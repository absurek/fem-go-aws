package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
		return fmt.Errorf("checking if user exists: %w", err)
	}

	if exists {
		return ErrUserAlreadyExists
	}

	user, err := newUser(data)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}

	if err := s.repository.InsertUser(ctx, user); err != nil {
		return fmt.Errorf("inserting user: %w", err)
	}

	return nil
}

// TODO(absurek): create and return JWT
func (s *Service) LoginUser(ctx context.Context, data UserData) (string, error) {
	user, err := s.repository.GetUser(ctx, data.Username)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}

		return "", fmt.Errorf("getting user: %w", err)
	}

	if !validatePassword(user.PasswordHash, data.Password) {
		return "", ErrInvalidCredentials
	}

	token, err := newToken(user)
	if err != nil {
		return "", fmt.Errorf("creating token: %w", err)
	}

	return token, nil
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

func newToken(user User) (string, error) {
	now := time.Now()
	validUntil := now.Add(time.Hour * 1)
	claims := jwt.MapClaims{
		"user":    user.Username,
		"expires": validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	secret := "secret" // TODO(absurek): store secret somewhere safe
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
