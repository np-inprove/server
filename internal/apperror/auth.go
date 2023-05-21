package apperror

import (
	"github.com/gookit/validate"
	"net/http"
)

var ErrUserNotFound = &ErrResponse{
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

var ErrInvalidPassword = &ErrResponse{
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
