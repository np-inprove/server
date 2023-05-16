package usecase

import "github.com/np-inprove/server/internal/repository"

type auth struct {
	r repository.Auth
}

type Auth interface {
}

func NewAuth() Auth {
	return auth{}
}
