package auth

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/middleware"
	"net/http"
)

type httpHandler struct {
	u UseCase
	c *config.Config
	j *jwtauth.JWTAuth
}

func NewHTTPHandler(u UseCase, c *config.Config, j *jwtauth.JWTAuth) chi.Router {
	a := &httpHandler{u, c, j}
	r := chi.NewRouter()

	r.Post("/login", a.Login)

	// Authenticated routes
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verify(j, func(r *http.Request) string {
			c, err := r.Cookie(c.AppJWTCookieName())
			if err != nil {
				return ""
			}
			return c.Value
		}))

		r.Use(middleware.Authenticator)

		r.Get("/whoami", a.WhoAmI)
	})

	return r
}

func (h httpHandler) Login(w http.ResponseWriter, r *http.Request) {
	p := &LoginRequest{}
	if err := render.Decode(r, p); err != nil {
		_ = render.Render(w, r, apperror.ErrBadRequest(err))
		return
	}

	fmt.Println()
}

func (h httpHandler) WhoAmI(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(jwtauth.TokenCtxKey)

	user, err := h.u.WhoAmI(r.Context(), token.(jwt.Token))
	if err != nil {
		_ = render.Render(w, r, apperror.ErrLoggedOut)
	}

	_ = render.Render(w, r, WhoAmIResponse{
		FirstName:              user.FirstName,
		LastName:               user.LastName,
		Email:                  user.Email,
		Points:                 user.Points,
		PointsAwardedCount:     user.PointsAwardedCount,
		PointsAwardedResetTime: user.PointsAwardedResetTime,
		GodMode:                user.GodMode,
	})
}
