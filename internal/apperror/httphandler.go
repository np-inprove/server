package apperror

import (
	"context"
	"github.com/go-chi/render"
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/logger"
	"net/http"
)

type ErrResponse struct {
	Err            error           `json:"-"`
	HTTPStatusCode int             `json:"-"`
	AppErrCode     int             `json:"code"`
	AppErrMessage  string          `json:"message"`
	Fields         validate.Errors `json:"fields,omitempty"`
}

func (e ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	if !e.Fields.Empty() {
		ctx := context.WithValue(r.Context(), logger.ErrValidationCtxKey, e.Fields)
		*r = *r.Clone(ctx)
	}

	if e.Err != nil {
		ctx := context.WithValue(r.Context(), logger.ErrCtxKey, e.Err)
		*r = *r.Clone(ctx)
	}

	render.Status(r, e.HTTPStatusCode)

	return nil
}

func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		AppErrCode:     http.StatusBadRequest,
		AppErrMessage:  http.StatusText(http.StatusBadRequest),
	}
}

func ErrValidation(fields validate.Errors) render.Renderer {
	return &ErrResponse{
		Err:            nil,
		HTTPStatusCode: http.StatusBadRequest,
		AppErrCode:     http.StatusBadRequest,
		AppErrMessage:  http.StatusText(http.StatusBadRequest),
		Fields:         fields,
	}
}

var ErrLoggedOut = &ErrResponse{
	Err:            nil,
	HTTPStatusCode: http.StatusUnauthorized,
	AppErrCode:     http.StatusUnauthorized,
	AppErrMessage:  http.StatusText(http.StatusUnauthorized),
}

func ErrInternal(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		AppErrCode:     http.StatusInternalServerError,
		AppErrMessage:  http.StatusText(http.StatusInternalServerError),
	}
}
