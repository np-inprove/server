package auth

import (
	"errors"
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/apperror"
	"net/http"
)

var (
	ErrInvalidPassword = errors.New("password provided did not match hash")
	ErrUserNotFound    = errors.New("failed to find user")
)

func mapDomainErr(err error) *apperror.ErrResponse {
	if errors.Is(err, ErrInvalidPassword) {
		return &apperror.ErrResponse{
			Err:            nil,
			HTTPStatusCode: http.StatusUnauthorized,
			AppErrCode:     http.StatusUnauthorized,
			AppErrMessage:  "Email or password is invalid",
			Fields: validate.Errors{
				"password": map[string]string{
					"invalid": "Email or password is invalid",
				},
			},
		}
	}

	if errors.Is(err, ErrUserNotFound) {
		return &apperror.ErrResponse{
			Err:            nil,
			HTTPStatusCode: http.StatusNotFound,
			AppErrCode:     http.StatusNotFound,
			AppErrMessage:  "The user was not found",
			Fields: validate.Errors{
				"email": map[string]string{
					"notFound": "A user does not exist with this email",
				},
			},
		}
	}
	return apperror.ErrInternal(err)
}
