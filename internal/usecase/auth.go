package usecase

type auth struct {
}

type Auth interface {
}

func NewAuth() Auth {
	return auth{}
}
