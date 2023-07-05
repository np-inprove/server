package group

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/entity/group"
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
	a := &httpHandler{u, c, j}
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

		r.Use(middleware.Authenticator)

		r.Get("/", a.ListGroups)
		r.Post("/", a.CreateGroup)
		r.Put("/{shortName}", a.UpdateGroup)
		r.Delete("/{shortName}", a.DeleteGroup)

		r.Get("/{shortName}/invites", a.ListInviteLinks)
		r.Post("/{shortName}/invites", a.CreateInviteLink)
		r.Delete("/{shortName}/invites/{code}", a.DeleteInviteLink)
	})

	return r
}

func (h httpHandler) ListGroups(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(jwtauth.TokenCtxKey)
	email := token.(jwt.Token).Subject()

	g, err := h.service.ListPrincipalGroups(r.Context(), email)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	res := make([]render.Renderer, len(g))
	for i, item := range g {
		res[i] = payload.Group{
			ID:          item.ID,
			ShortName:   item.ShortName,
			Name:        item.Name,
			Description: item.Description,
		}
	}

	_ = render.RenderList(w, r, res)
}

func (h httpHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	p := &payload.CreateGroupRequest{}
	if err := render.Decode(r, p); err != nil {
		_ = render.Render(w, r, apperror.ErrBadRequest(err))
		return
	}

	if v := p.Validate(); !v.Validate() {
		_ = render.Render(w, r, apperror.ErrValidation(v.Errors))
		return
	}

	token := r.Context().Value(jwtauth.TokenCtxKey)
	email := token.(jwt.Token).Subject()

	res, err := h.service.CreateGroup(r.Context(), email,
		group.Name(p.Name),
		group.ShortName(p.ShortName),
		group.Description(p.Description),
	)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	_ = render.Render(w, r, payload.Group{
		ID:          res.ID,
		ShortName:   res.ShortName,
		Name:        res.Name,
		Description: res.Description,
	})
}

func (h httpHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	p := &payload.UpdateGroupRequest{}
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

	res, err := h.service.UpdateGroup(r.Context(), email, shortName,
		group.Name(p.Name),
		group.ShortName(p.ShortName),
		group.Description(p.Description),
	)

	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	_ = render.Render(w, r, payload.Group{
		ID:          res.ID,
		ShortName:   res.ShortName,
		Name:        res.Name,
		Description: res.Description,
	})
}

func (h httpHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "shortName")
	token := r.Context().Value(jwtauth.TokenCtxKey)
	email := token.(jwt.Token).Subject()

	if err := h.service.DeleteGroup(r.Context(), email, path); err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	render.NoContent(w, r)
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
		p[i] = payload.GroupInviteLink{
			ID:   v.ID,
			Code: v.Code,
			Role: v.Role,
		}
	}

	_ = render.RenderList(w, r, p)
}

func (h httpHandler) CreateInviteLink(w http.ResponseWriter, r *http.Request) {
	p := &payload.CreateGroupInviteLinkRequest{}

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

	_ = render.Render(w, r, payload.GroupInviteLink{
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
