package shared

import (
	"fmt"
	"net/http"
)

const (
	DatabaseNotFoundError = "DB_NOT_FOUND"
	DatabaseFindError     = "DB_FIND"
)

type EchoErrorResponse struct {
	Message string `json:"message"`
}

type ApplicationError struct {
	ErrorCode        string
	ErrorDescription string
	ErrorOrigin      string
	HTTPStatusCode   int
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("[%s]: %s at %s - sending: %d", e.ErrorCode, e.ErrorDescription, e.ErrorOrigin, e.HTTPStatusCode)
}

func BuildError(httpStatusCode int, errorCode, errorDescription, errorOrigin string) *ApplicationError {
	return &ApplicationError{
		HTTPStatusCode:   httpStatusCode,
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
		ErrorOrigin:      errorOrigin,
	}
}

func GetHTTPStatusErrorCode(err error) int {
	applicationError, ok := err.(*ApplicationError)
	if ok {
		return applicationError.HTTPStatusCode
	}

	return http.StatusInternalServerError
}
