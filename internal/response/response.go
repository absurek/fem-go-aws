package response

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func Ok(msg string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       msg,
		StatusCode: http.StatusOK,
	}
}

func Created(msg string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       msg,
		StatusCode: http.StatusCreated,
	}
}

func BadRequest() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "bad request",
		StatusCode: http.StatusBadRequest,
	}
}

func NotFound() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "not found",
		StatusCode: http.StatusNotFound,
	}
}

func Unauthorized() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "unauthorized",
		StatusCode: http.StatusUnauthorized,
	}
}

func InternalServerError() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "internal server error",
		StatusCode: http.StatusInternalServerError,
	}
}
