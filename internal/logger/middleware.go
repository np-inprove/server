package logger

import (
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
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
			end := time.Now()
			go func() {
				m.logger.Info("handled http request",
					String("method", r.Method),
					String("url", r.URL.String()),
					Int("status", ww.Status()),
					Int("bytesWritten", ww.BytesWritten()),
					Duration("duration", end.Sub(start)))
			}()
		}()

		next.ServeHTTP(ww, r)
	})
}
