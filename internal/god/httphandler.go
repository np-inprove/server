package god

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/middleware"
	"github.com/np-inprove/server/internal/payload"
	"net/http"
)

type httpHandler struct {
	u UseCase
	c *config.Config
	j *jwtauth.JWTAuth
}

func NewHTTPHandler(u UseCase, c *config.Config, j *jwtauth.JWTAuth) chi.Router {
	h := httpHandler{u, c, j}
	r := chi.NewRouter()

	// Authenticated routes
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verify(j, func(r *http.Request) string {
			c, err := r.Cookie(c.AppJWTCookieName())
			if err != nil {
				return ""
			}
			return c.Value
		}))

		r.Use(middleware.GodAuthenticator)

		r.Get("/institutions", h.ListInstitutions)
		r.Post("/institutions", h.CreateInstitution)
		r.Delete("/institutions/{institutionShortName}", h.DeleteInstitution)
	})

	return r
}

func (h httpHandler) ListInstitutions(w http.ResponseWriter, r *http.Request) {
	insts, err := h.u.ListInstitutions(r.Context())
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	l := make([]render.Renderer, len(insts))
	for _, inst := range insts {
		l = append(l, payload.Institution{
			ID:            inst.ID,
			Name:          inst.Name,
			ShortName:     inst.ShortName,
			AdminDomain:   inst.AdminDomain,
			StudentDomain: inst.StudentDomain,
		})
	}

	_ = render.RenderList(w, r, l)
}

func (h httpHandler) CreateInstitution(w http.ResponseWriter, r *http.Request) {
	p := payload.CreateInstitutionRequest{}
	if err := render.Decode(r, p); err != nil {
		_ = render.Render(w, r, apperror.ErrBadRequest(err))
		return
	}

	if v := p.Validate(); !v.Validate() {
		_ = render.Render(w, r, apperror.ErrValidation(v.Errors))
		return
	}

	inst, err := h.u.CreateInstitution(r.Context(), p.Name, p.ShortName, p.AdminDomain, p.StudentDomain)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	_ = render.Render(w, r, payload.Institution{
		ID:            inst.ID,
		Name:          inst.Name,
		ShortName:     inst.ShortName,
		AdminDomain:   inst.AdminDomain,
		StudentDomain: inst.StudentDomain,
	})
}

func (h httpHandler) DeleteInstitution(w http.ResponseWriter, r *http.Request) {
	shortName := chi.URLParam(r, "institutionShortName")
	err := h.u.DeleteInstitution(r.Context(), shortName)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	render.NoContent(w, r)
}
