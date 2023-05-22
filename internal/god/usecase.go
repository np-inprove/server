package god

import "context"

type UseCase interface {
	ListInstitutions(ctx context.Context) ([]Institution, error)
	CreateInstitution(ctx context.Context, name string, shortName string) (Institution, error)
	DeleteInstitution(ctx context.Context, id int) error

	AddInstitutionAdmin(ctx context.Context, id int)
	DeleteInstitutionAdmin(ctx context.Context, id int) (Institution, error)
}
