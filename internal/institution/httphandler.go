package institution

import (
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/middleware"
	"github.com/np-inprove/server/internal/payload"
)

type httpHandler struct {
	service UseCase
	cfg     *config.Config
	jwt     *jwtauth.JWTAuth
}

func NewHTTPHandler(u UseCase, c *config.Config, j *jwtauth.JWTAuth) chi.Router {
	h := httpHandler{u, c, j}
	r := chi.NewRouter()

	r.Use(jwtauth.Verify(j, func(r *http.Request) string {
		c, err := r.Cookie(c.AppJWTCookieName())
		if err != nil {
			return ""
		}
		return c.Value
	}))

	// God-mode authenticated routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.GodAuthenticator)

		r.Get("/", h.ListInstitutions)
		r.Post("/", h.CreateInstitution)
		r.Delete("/{shortName}", h.DeleteInstitution)
		r.Put("/{shortName}", h.UpdateInstitution)
	})

	// Normal authenticated routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticator)

		r.Get("/{shortName}/invites", h.ListInviteLinks)
		r.Post("/{shortName}/invites", h.CreateInviteLink)
		r.Delete("/{shortName}/invites/{code}", h.DeleteInviteLink)
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
			ID:          inst.ID,
			Name:        inst.Name,
			ShortName:   inst.ShortName,
			Description: inst.Description,
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
	shortName := chi.URLParam(r, "shortName")
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
	if err := render.Decode(r, p); err != nil {
		_ = render.Render(w, r, apperror.ErrBadRequest(err))
		return
	}

	if v := p.Validate(); !v.Validate() {
		_ = render.Render(w, r, apperror.ErrValidation(v.Errors))
		return
	}

	shortName := chi.URLParam(r, "shortName")

	inst, err := h.service.UpdateInstitution(r.Context(), shortName, p.Name, p.ShortName, p.Description)
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

func (h httpHandler) ListInviteLinks(w http.ResponseWriter, r *http.Request) {
	shortName := chi.URLParam(r, "shortName")

	token := r.Context().Value(jwtauth.TokenCtxKey)
	email := token.(jwt.Token).Subject()

	links, err := h.service.ListInviteLinks(r.Context(), email, shortName)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	p := make([]render.Renderer, len(links))
	for i, v := range links {
		p[i] = payload.InstitutionInviteLink{
			ID:   v.ID,
			Code: v.Code,
			Role: v.Role,
		}
	}

	_ = render.RenderList(w, r, p)
}

func (h httpHandler) CreateInviteLink(w http.ResponseWriter, r *http.Request) {
	p := &payload.CreateInviteLinkRequest{}

	if err := render.Decode(r, p); err != nil {
		_ = render.Render(w, r, apperror.ErrBadRequest(err))
		return
	}

	if v := p.Validate(); !v.Validate() {
		_ = render.Render(w, r, apperror.ErrValidation(v.Errors))
		return
	}

	shortName := chi.URLParam(r, "shortName")
	token := r.Context().Value(jwtauth.TokenCtxKey)
	email := token.(jwt.Token).Subject()

	link, err := h.service.CreateInviteLink(r.Context(), email, shortName, p.Role)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	_ = render.Render(w, r, payload.InstitutionInviteLink{
		ID:   link.ID,
		Code: link.Code,
		Role: link.Role,
	})
}

func (h httpHandler) DeleteInviteLink(w http.ResponseWriter, r *http.Request) {
	shortName := chi.URLParam(r, "shortName")
	code := chi.URLParam(r, "code")

	token := r.Context().Value(jwtauth.TokenCtxKey)
	email := token.(jwt.Token).Subject()

	err := h.service.DeleteInviteLink(r.Context(), email, shortName, code)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
