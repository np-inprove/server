package dash

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/config"
	group2 "github.com/np-inprove/server/internal/entity/group"
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

		r.Get("/groups", a.ListGroups)
		r.Get("/grouptypes", a.ListGroupTypes)
		r.Post("/groups", a.CreateGroup)
		r.Delete("/groups/{path}", a.DeleteGroup)
	})

	return r
}

func (h httpHandler) ListGroups(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(jwtauth.TokenCtxKey)
	email := token.(jwt.Token).Subject()

	g, err := h.service.ListGroups(r.Context(), email)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	res := make([]render.Renderer, len(g))
	for i, item := range g {
		res[i] = payload.Group{
			ID:          item.ID,
			Path:        item.Path,
			Name:        item.Name,
			Description: item.Description,
			GroupType:   item.GroupType.String(),
		}
	}

	_ = render.RenderList(w, r, res)
}

func (h httpHandler) ListGroupTypes(w http.ResponseWriter, r *http.Request) {
	gt, err := h.service.ListGroupTypes()
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	res := make([]render.Renderer, len(gt))
	for i, item := range gt {
		res[i] = payload.GroupType(item.String())
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

	res, err := h.service.CreateGroup(r.Context(), email, p.GroupType,
		group2.Name(p.Name),
		group2.Path(p.Path),
		group2.Description(p.Description),
	)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	_ = render.Render(w, r, payload.Group{
		ID:          res.ID,
		Path:        res.Path,
		Name:        res.Name,
		Description: res.Description,
		GroupType:   res.GroupType.String(),
	})
}

func (h httpHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "path")
	token := r.Context().Value(jwtauth.TokenCtxKey)
	email := token.(jwt.Token).Subject()

	if err := h.service.DeleteGroup(r.Context(), email, path); err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	render.NoContent(w, r)
}
