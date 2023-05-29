package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/middleware"
	"github.com/np-inprove/server/internal/payload"
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
	r.Post("/register", a.Register)

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
	p := &payload.LoginRequest{}
	if err := render.Decode(r, p); err != nil {
		_ = render.Render(w, r, apperror.ErrBadRequest(err))
		return
	}

	if v := p.Validate(); !v.Validate() {
		_ = render.Render(w, r, apperror.ErrValidation(v.Errors))
		return
	}

	s, err := h.u.Login(r.Context(), p.Email, p.Password)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	h.setAuthCookies(w, s)

	_ = render.Render(w, r, payload.User{
		FirstName:              s.user.FirstName,
		LastName:               s.user.LastName,
		Email:                  s.user.Email,
		Points:                 s.user.Points,
		PointsAwardedCount:     s.user.PointsAwardedCount,
		PointsAwardedResetTime: s.user.PointsAwardedResetTime,
		GodMode:                s.user.GodMode,
	})
}

func (h httpHandler) Register(w http.ResponseWriter, r *http.Request) {
	p := &payload.RegisterRequest{}
	if err := render.Decode(r, p); err != nil {
		_ = render.Render(w, r, apperror.ErrBadRequest(err))
		return
	}

	if v := p.Validate(); !v.Validate() {
		_ = render.Render(w, r, apperror.ErrValidation(v.Errors))
		return
	}

	s, err := h.u.Register(r.Context(), p.FirstName, p.LastName, p.Email, p.Password, p.ConfirmPassword)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	h.setAuthCookies(w, s)

	_ = render.Render(w, r, payload.User{
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

	_ = render.Render(w, r, payload.User{
		FirstName:              user.FirstName,
		LastName:               user.LastName,
		Email:                  user.Email,
		Points:                 user.Points,
		PointsAwardedCount:     user.PointsAwardedCount,
		PointsAwardedResetTime: user.PointsAwardedResetTime,
		GodMode:                user.GodMode,
	})
}

func (h httpHandler) setAuthCookies(w http.ResponseWriter, s *session) {
	// JWT cookie for the server to use
	http.SetCookie(w, &http.Cookie{
		Name:     h.c.AppJWTCookieName(),
		Domain:   h.c.AppJWTCookieDomain(),
		Value:    s.token,
		Path:     "/",
		Expires:  time.Now().Add(30 * time.Minute),
		MaxAge:   int(30 * time.Minute.Seconds()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	})

	b := "human"
	if s.user.GodMode {
		b = "god"
	}
	// Cookie for the client to know that it's authenticated
	// This must be the same as specified in https://github.com/np-inprove/app/blob/main/middleware
	http.SetCookie(w, &http.Cookie{
		Name:     "b",
		Domain:   h.c.AppJWTCookieDomain(),
		Value:    b,
		Path:     "/",
		Expires:  time.Now().Add(30 * time.Minute),
		MaxAge:   int(30 * time.Minute.Seconds()),
		Secure:   true,
		SameSite: http.SameSiteDefaultMode,
	})
}
