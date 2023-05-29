package auth

import (
	"errors"
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/apperror"
	"net/http"
)

var (
	ErrInvalidPassword  = errors.New("password provided did not match hash")
	ErrUserNotFound     = errors.New("failed to find user")
	ErrUserConflict     = errors.New("failed to register new user as email already used")
	ErrDomainNotFound   = errors.New("no institution is registered with provided domain")
	ErrPasswordConflict = errors.New("passwords provided does not match")
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

	if errors.Is(err, ErrDomainNotFound) {
		return &apperror.ErrResponse{
			Err:            nil,
			HTTPStatusCode: http.StatusBadRequest,
			AppErrCode:     http.StatusBadRequest,
			AppErrMessage:  "An institution does not exist for this domain. Please get your institution administrator to register first.",
			Fields: validate.Errors{
				"email": map[string]string{
					"domain": "Email domain not found",
				},
			},
		}
	}

	if errors.Is(err, ErrPasswordConflict) {
		return &apperror.ErrResponse{
			Err:            nil,
			HTTPStatusCode: http.StatusBadRequest,
			AppErrCode:     http.StatusBadRequest,
			AppErrMessage:  "The passwords provided does not match. Please retype the passwords",
			Fields: validate.Errors{
				"confirmPassword": map[string]string{
					"invalid": "Password fields do not match",
				},
			},
		}
	}

	return apperror.ErrInternal(err)
}
