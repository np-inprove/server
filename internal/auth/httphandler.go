package auth

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/gookit/validate"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/middleware"
	"net/http"
	"time"
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

	v := validate.Struct(p)
	if !v.Validate() {
		_ = render.Render(w, r, apperror.ErrValidation(v.Errors))
		return
	}

	s, err := h.u.Login(r.Context(), p.Email, p.Password)
	if err != nil {
		// TODO something better to map from domain to HTTP error
		if apperror.IsEntityNotFound(err) {
			_ = render.Render(w, r, apperror.ErrUserNotFound)
			return
		}
		if errors.Is(err, ErrInvalidPassword) {
			_ = render.Render(w, r, apperror.ErrInvalidPassword)
			return
		}
		_ = render.Render(w, r, apperror.ErrBadRequest(err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     h.c.AppJWTCookieName(),
		Value:    s.token,
		Domain:   "np-inprove.qinguan.me",
		Path:     "/",
		Expires:  time.Now().Add(30 * time.Minute),
		MaxAge:   int(30 * time.Minute.Seconds()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	})

	_ = render.Render(w, r, LoginResponse{
		FirstName:              s.user.FirstName,
		LastName:               s.user.LastName,
		Email:                  s.user.Email,
		Points:                 s.user.Points,
		PointsAwardedCount:     s.user.PointsAwardedCount,
		PointsAwardedResetTime: s.user.PointsAwardedResetTime,
		GodMode:                s.user.GodMode,
	})
}

func (h httpHandler) WhoAmI(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(jwtauth.TokenCtxKey)

	user, err := h.u.WhoAmI(r.Context(), token.(jwt.Token))
	if err != nil {
		_ = render.Render(w, r, apperror.ErrLoggedOut)
		return
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
