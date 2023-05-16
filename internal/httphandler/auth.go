package httphandler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/np-inprove/server/internal/payload"
	"github.com/np-inprove/server/internal/usecase"
	"net/http"
)

type auth struct {
	u usecase.Auth
}

func NewAuth(u usecase.Auth) chi.Router {
	a := &auth{u}
	m := chi.NewRouter()
	m.Post("/login", a.Login)
	return m
}

func (a auth) Login(w http.ResponseWriter, r *http.Request) {
	p := &payload.LoginRequest{}
	if err := render.Decode(r, p); err != nil {
		_ = render.Render(w, r, payload.ErrBadRequest(err))
		return
	}

	fmt.Println()
}
