package user

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/absurek/fem-go-aws/internal/response"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrMissingAuthHeader      = errors.New("missing authorization header")
	ErrMalformedAuthHeader    = errors.New("malformed authorization header")
	ErrInvalidAuthToken       = errors.New("invalid authorization token")
	ErrInvalidAuthTokenClaims = errors.New("invalid authorization token claims")
)

type HandlerFunction = func(ctx context.Context, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse

func ValidateJWTMiddleware(next HandlerFunction) HandlerFunction {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
		token, err := authToken(request.Headers)
		if err != nil {
			return response.Unauthorized()
		}

		claims, err := parseJWT(token)
		if err != nil {
			return response.Unauthorized()
		}

		expired, err := isExpired(claims)
		if expired || err != nil {
			return response.Unauthorized()
		}

		return next(ctx, request)
	}
}

func authToken(headers map[string]string) (string, error) {
	authHeader, ok := headers["Authorization"]
	if !ok {
		return "", ErrMissingAuthHeader
	}

	parts := strings.Split(authHeader, "Bearer ")
	if len(parts) != 2 {
		return "", ErrMalformedAuthHeader
	}

	return parts[1], nil
}

func parseJWT(token string) (jwt.MapClaims, error) {
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		secret := "secret" // TODO(absurek): store somewhere safe
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parsing JWT")
	}

	if !parsed.Valid {
		return nil, ErrInvalidAuthToken
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidAuthTokenClaims
	}

	return claims, nil
}

// TODO(absurek): this is atrocious
func isExpired(claims jwt.MapClaims) (bool, error) {
	expires, ok := claims["expires"].(float64)
	if !ok {
		return true, ErrInvalidAuthToken
	}

	if time.Now().Unix() > int64(expires) {
		return true, nil
	}

	return false, nil
}
