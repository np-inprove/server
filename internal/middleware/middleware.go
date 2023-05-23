package middleware

import (
	"context"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/np-inprove/server/internal/apperror"
	"net/http"

	"github.com/lestrrat-go/jwx/v2/jwt"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := FromContext(r.Context())

		if err != nil {
			_ = render.Render(w, r, apperror.ErrLoggedOut)
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			_ = render.Render(w, r, apperror.ErrLoggedOut)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

func GodAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := FromContext(r.Context())

		if err != nil {
			_ = render.Render(w, r, apperror.ErrLoggedOut)
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			_ = render.Render(w, r, apperror.ErrLoggedOut)
			return
		}

		if v, ok := token.Get("god"); !ok || !v.(bool) {
			_ = render.Render(w, r, apperror.ErrLoggedOut)
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

func FromContext(ctx context.Context) (jwt.Token, map[string]interface{}, error) {
	token, _ := ctx.Value(jwtauth.TokenCtxKey).(jwt.Token)

	var err error
	var claims map[string]interface{}

	if token != nil {
		claims, err = token.AsMap(context.Background())
		if err != nil {
			return token, nil, err
		}
	} else {
		claims = map[string]interface{}{}
	}

	err, _ = ctx.Value(jwtauth.ErrorCtxKey).(error)

	return token, claims, err
}
