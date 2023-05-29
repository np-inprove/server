package logger

import (
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type ctxKey int

const (
	ErrCtxKey ctxKey = iota
	ErrValidationCtxKey
)

type Middleware struct {
	logger AppLogger
}

func NewMiddleware(logger AppLogger) *Middleware {
	return &Middleware{
		logger,
	}
}

func (m *Middleware) Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, r *http.Request) {
		ww := chimiddleware.NewWrapResponseWriter(res, r.ProtoMajor)

		start := time.Now()
		defer func() {
			// TODO this does not work
			fields := []Field{
				String("area", "http"),
				String("method", r.Method),
				String("url", r.URL.String()),
				Int("status", ww.Status()),
				Int("bytes_written", ww.BytesWritten()),
				Duration("duration", time.Since(start)),
			}

			v := r.Context().Value(ErrCtxKey)
			if v != nil {
				fields = append(fields, String("err", v.(error).Error()))
			}

			v = r.Context().Value(ErrValidationCtxKey)
			if v != nil {
				fields = append(fields, String("err_validation", v.(string)))
			}

			m.logger.Info("handled http request",
				fields...,
			)
		}()

		next.ServeHTTP(ww, r)
	})
}
