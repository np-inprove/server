package institution

import (
	"context"
	"github.com/np-inprove/server/internal/entity/institution"

	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/transactor"
)

type Reader interface {
	FindInstitution(ctx context.Context, shortName string) (*entity.Institution, error)
	FindInstitutions(ctx context.Context) ([]*entity.Institution, error)
	FindInstitutionWithInvites(ctx context.Context, shortName string) (*entity.Institution, error)
	FindInviteLinkWithInstitution(ctx context.Context, code string) (*entity.InstitutionInviteLink, error)
	FindInviteLinks(ctx context.Context, id int) ([]*entity.InstitutionInviteLink, error)
	FindUser(ctx context.Context, principal string) (*entity.User, error)
}

type Writer interface {
	CreateInstitution(ctx context.Context, name, shortName, description string) (*entity.Institution, error)
	UpdateInstitution(ctx context.Context, id int, name, shortName, description string) (*entity.Institution, error)
	DeleteInstitution(ctx context.Context, id int) error
	CreateInviteLink(ctx context.Context, id int, code string, role institution.Role) (*entity.InstitutionInviteLink, error)
	DeleteInviteLink(ctx context.Context, id int) error
}

type Repository interface {
	Reader
	Writer
	transactor.Transactor
}
