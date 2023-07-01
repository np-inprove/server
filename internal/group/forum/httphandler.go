package forum

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"

	"net/http"

	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/v2/jwt"
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

		r.Get("/", a.ListForums)
		r.Post("/", a.CreateForum)
		r.Delete("/{shortName}", a.DeleteForum)
		r.Put("/{shortName}", a.UpdateForum)
	})

	return r
}

func (h httpHandler) ListForums(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(jwtauth.TokenCtxKey)
	email := token.(jwt.Token).Subject()

	forums, err := h.service.ListPrincipalForums(r.Context(), email)
	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	res := make([]render.Renderer, len(forums))
	for i, item := range forums {
		res[i] = payload.Forum{
			ID:          item.ID,
			ShortName:   item.ShortName,
			Name:        item.Name,
			Description: item.Description,
		}
	}

	_ = render.RenderList(w, r, res)
}

func (h httpHandler) CreateForum(w http.ResponseWriter, r *http.Request) {
	p := &payload.CreateForumRequest{}
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

	res, err := h.service.CreateForum(r.Context(), email,
		p.Name,
		p.ShortName,
		p.Description,
	)

	if err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	_ = render.Render(w, r, payload.Forum{
		ID:          res.ID,
		ShortName:   res.ShortName,
		Name:        res.Name,
		Description: res.Description,
	})
}

func (h httpHandler) DeleteForum(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "shortName")
	token := r.Context().Value(jwtauth.TokenCtxKey)
	email := token.(jwt.Token).Subject()

	if err := h.service.DeleteForum(r.Context(), email, path); err != nil {
		_ = render.Render(w, r, mapDomainErr(err))
		return
	}

	render.NoContent(w, r)
}

func (h httpHandler) UpdateForum(w http.ResponseWriter, r *http.Request) {
	shortName := chi.URLParam(r, "shortName")
	p := payload.UpdateForumRequest{}
	if err := render.Decode(r, p); err != nil {
		_ = render.Render(w, r, apperror.ErrBadRequest(err))
		return
	}

	forum, err := h.service.UpdateForum(r.Context(), shortName,
		p.Name,
		p.ShortName,
		p.Description,
	)
	if err != nil {
		_ = render.Render(w, r, payload.Forum{
			ID:          forum.ID,
			Name:        forum.Name,
			ShortName:   forum.ShortName,
			Description: forum.Description,
		})
	}
}
