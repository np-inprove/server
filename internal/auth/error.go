package auth

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/apperror"
	"net/http"
)

var (
	ErrInvalidPassword = errors.New("password provided did not match hash")
	ErrUserNotFound    = errors.New("failed to find user")
	ErrUserConflict    = errors.New("failed to register new user as email already used")
	ErrTokenRevoked    = errors.New("token is revoked")
	ErrInvalidInvite   = errors.New("invalid invite code provided")
)

func mapDomainErr(err error) render.Renderer {
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

	if errors.Is(err, ErrUserConflict) {
		return &apperror.ErrResponse{
			Err:            nil,
			HTTPStatusCode: http.StatusConflict,
			AppErrCode:     http.StatusConflict,
			AppErrMessage:  "A user exists with this email",
			Fields: validate.Errors{
				"email": map[string]string{
					"conflict": "A user exists with this email",
				},
			},
		}
	}

	if errors.Is(err, ErrTokenRevoked) {
		return &apperror.ErrResponse{
			Err:            nil,
			HTTPStatusCode: http.StatusUnauthorized,
			AppErrCode:     http.StatusUnauthorized,
			AppErrMessage:  "Token is revoked",
		}
	}

	return apperror.ErrInternal(err)
}
