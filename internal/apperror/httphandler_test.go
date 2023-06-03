package apperror

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/logger"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HTTPHandlerTestSuite struct {
	suite.Suite
}

func TestHTTPHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HTTPHandlerTestSuite))
}

func (suite *HTTPHandlerTestSuite) TestErrResponseRender() {
	type args struct {
		e func() ErrResponse
		w func() http.ResponseWriter
		r func() *http.Request
	}

	sampleErr := errors.New("sample err")

	tests := []struct {
		name      string
		args      args
		assertion func(r *http.Request, e ErrResponse)
	}{
		{
			name: "Renders status code",
			args: args{
				e: func() ErrResponse {
					return ErrResponse{
						Err:            nil,
						HTTPStatusCode: http.StatusTeapot,
						AppErrCode:     http.StatusTeapot,
						AppErrMessage:  sampleErr.Error(),
						Fields:         nil,
					}
				},
				w: func() http.ResponseWriter {
					return httptest.NewRecorder()
				},
				r: func() *http.Request {
					return httptest.NewRequest("", "/", nil)
				},
			},
			assertion: func(r *http.Request, e ErrResponse) {
				suite.Equal(http.StatusTeapot, r.Context().Value(render.StatusCtxKey))
			},
		},
		{
			name: "Sets validation error context",
			args: args{
				e: func() ErrResponse {
					return ErrResponse{
						Err:            nil,
						HTTPStatusCode: http.StatusTeapot,
						AppErrCode:     http.StatusTeapot,
						AppErrMessage:  sampleErr.Error(),
						Fields: validate.Errors{
							"field": map[string]string{
								"err": "bad err",
							},
						},
					}
				},
				w: func() http.ResponseWriter {
					return httptest.NewRecorder()
				},
				r: func() *http.Request {
					return httptest.NewRequest("", "/", nil)
				},
			},
			assertion: func(r *http.Request, e ErrResponse) {
				suite.Equal(e.Fields, r.Context().Value(logger.ErrValidationCtxKey))
			},
		},
		{
			name: "Sets error context",
			args: args{
				e: func() ErrResponse {
					return ErrResponse{
						Err:            sampleErr,
						HTTPStatusCode: http.StatusTeapot,
						AppErrCode:     http.StatusTeapot,
						AppErrMessage:  sampleErr.Error(),
						Fields:         nil,
					}
				},
				w: func() http.ResponseWriter {
					return httptest.NewRecorder()
				},
				r: func() *http.Request {
					return httptest.NewRequest("", "/", nil)
				},
			},
			assertion: func(r *http.Request, e ErrResponse) {
				suite.Equal(e.Err, r.Context().Value(logger.ErrCtxKey))
			},
		},
		{
			name: "Sets both error and validation context",
			args: args{
				e: func() ErrResponse {
					return ErrResponse{
						Err:            sampleErr,
						HTTPStatusCode: http.StatusTeapot,
						AppErrCode:     http.StatusTeapot,
						AppErrMessage:  sampleErr.Error(),
						Fields: validate.Errors{
							"field": map[string]string{
								"err": "bad err",
							},
						},
					}
				},
				w: func() http.ResponseWriter {
					return httptest.NewRecorder()
				},
				r: func() *http.Request {
					return httptest.NewRequest("", "/", nil)
				},
			},
			assertion: func(r *http.Request, e ErrResponse) {
				suite.Equal(e.Err, r.Context().Value(logger.ErrCtxKey))
				suite.Equal(e.Fields, r.Context().Value(logger.ErrValidationCtxKey))
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			r := tc.args.r()
			err := tc.args.e()
			ret := err.Render(tc.args.w(), r)
			tc.assertion(r, err)
			suite.Nil(ret)
		})
	}
}
