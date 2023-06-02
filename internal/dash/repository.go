package dash

import (
	"context"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/group"
)

type Reader interface {
	FindInstitutionByAdminDomain(ctx context.Context, domain string) (*entity.Institution, error)
	FindGroupsByUser(ctx context.Context, email string) ([]*entity.Group, error)
	FindGroupTypes() ([]*entity.GroupType, error)
}

type Writer interface {
	CreateGroup(ctx context.Context, groupType entity.GroupType, opts ...group.Option) (*entity.Group, error)
	DeleteGroup(ctx context.Context, path string) error
}

type Repository interface {
	Reader
	Writer
}
