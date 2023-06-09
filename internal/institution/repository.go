package institution

import (
	"context"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/transactor"
)

type Reader interface {
	FindInstitutions(ctx context.Context) ([]*entity.Institution, error)
	FindInstitutionByAdminDomain(ctx context.Context, domain string) (*entity.Institution, error)
}

type Writer interface {
	CreateInstitution(
		ctx context.Context,
		name string,
		shortName string,
		adminDomain string,
		studentDomain string,
	) (*entity.Institution, error)
	DeleteInstitution(ctx context.Context, shortName string) error

	CreateDepartment(
		ctx context.Context,
		name string,
		shortName string,
		institutionID int,
	) (*entity.Department, error)
}

type Repository interface {
	Reader
	Writer
	transactor.Transactor
}
