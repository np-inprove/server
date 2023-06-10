package institution

import (
	"context"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/transactor"
)

type Reader interface {
	FindInstitutions(ctx context.Context) ([]*entity.Institution, error)
	FindInstitution(ctx context.Context, shortName string) (*entity.Institution, error)
}

type Writer interface {
	CreateInstitution(ctx context.Context, name, shortName, description string) (*entity.Institution, error)
	DeleteInstitution(ctx context.Context, id int) error
	UpdateInstitution(
		ctx context.Context,
		id int,
		name string,
		shortName string,
		adminDomain string,
		studentDomain string,
	) (*entity.Institution, error)
}

type Repository interface {
	Reader
	Writer
	transactor.Transactor
}
