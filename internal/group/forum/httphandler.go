package forum

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"

	/* "github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/np-inprove/server/internal/apperror" */
	"github.com/np-inprove/server/internal/config"
	/* "github.com/np-inprove/server/internal/entity/group" */
	"github.com/np-inprove/server/internal/middleware"
	/* "github.com/np-inprove/server/internal/payload" */
	"net/http"
)

type httpHandler struct {
	//service UseCase
	cfg *config.Config
	jwt *jwtauth.JWTAuth
}

func NewHTTPHandler( /* u UseCase, */ c *config.Config, j *jwtauth.JWTAuth) chi.Router {
	a := &httpHandler{c, j}
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

}

func (h httpHandler) CreateForum(w http.ResponseWriter, r *http.Request) {

}

func (h httpHandler) DeleteForum(w http.ResponseWriter, r *http.Request) {

}

func (h httpHandler) UpdateForum(w http.ResponseWriter, r *http.Request) {

}
