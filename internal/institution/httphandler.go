package institution

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
	service UseCase
	cfg     *config.Config
	jwt     *jwtauth.JWTAuth
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

		r.Get("/", h.ListInstitutions)
		r.Post("/", h.CreateInstitution)
		r.Delete("/{id}", h.DeleteInstitution)
		r.Put("/{id}", h.UpdateInstitution) 
	})

	return r
}

func (h httpHandler) ListInstitutions(w http.ResponseWriter, r *http.Request) {
	insts, err := h.service.ListInstitutions(r.Context())
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	l := make([]render.Renderer, len(insts))
	for i, inst := range insts {
		l[i] = payload.Institution{
			ID:        inst.ID,
			Name:      inst.Name,
			ShortName: inst.ShortName,
		}
	}

	_ = render.RenderList(w, r, l)
}

func (h httpHandler) CreateInstitution(w http.ResponseWriter, r *http.Request) {
	p := &payload.CreateInstitutionRequest{}
	if err := render.Decode(r, p); err != nil {
		_ = render.Render(w, r, apperror.ErrBadRequest(err))
		return
	}

	if v := p.Validate(); !v.Validate() {
		_ = render.Render(w, r, apperror.ErrValidation(v.Errors))
		return
	}

	inst, err := h.service.CreateInstitution(r.Context(), p.Name, p.ShortName, p.Description)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	_ = render.Render(w, r, payload.Institution{
		ID:          inst.ID,
		Name:        inst.Name,
		ShortName:   inst.ShortName,
		Description: inst.Description,
	})
}

func (h httpHandler) DeleteInstitution(w http.ResponseWriter, r *http.Request) {
	shortName := chi.URLParam(r, "id") //rename
	err := h.service.DeleteInstitution(r.Context(), shortName)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	render.Status(r, http.StatusNoContent)
	render.NoContent(w, r)
}

func (h httpHandler) UpdateInstitution(w http.ResponseWriter, r *http.Request) {
	p := &payload.UpdateInstitutionRequest{} 

	if v := p.Validate(); !v.Validate() {
		_ = render.Render(w, r, apperror.ErrValidation(v.Errors))
		return
	}

	inst, err := h.service.UpdateInstitution(r.Context(), p.ID, p.Name, p.ShortName, p.AdminDomain, p.StudentDomain)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	render.Status(r, http.StatusNoContent)
	_ = render.Render(w, r, payload.Institution{
		ID:            inst.ID,
		Name:          inst.Name,
		ShortName:     inst.ShortName,
		AdminDomain:   inst.AdminDomain,
		StudentDomain: inst.StudentDomain,
	})
}
