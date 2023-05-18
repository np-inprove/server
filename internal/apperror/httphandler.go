package apperror

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	AppErrCode     int    `json:"code"`
	AppErrMessage  string `json:"message"`
	Fields         string `json:"fields,omitempty"`
}

func (e ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		AppErrCode:     http.StatusBadRequest,
		AppErrMessage:  http.StatusText(http.StatusBadGateway),
	}
}

var ErrLoggedOut = &ErrResponse{
	Err:            nil,
	HTTPStatusCode: http.StatusUnauthorized,
	AppErrCode:     http.StatusUnauthorized,
	AppErrMessage:  http.StatusText(http.StatusUnauthorized),
}
