package group

import (
	"context"

	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/group"
)

type Reader interface {
	FindGroupsByUser(ctx context.Context, principal string) ([]*entity.Group, error)
	FindGroupByInstitutionIDAndShortName(ctx context.Context, institutionID int, shortName string) (*entity.Group, error)

	FindUserWithInstitution(ctx context.Context, principal string) (*entity.User, error)
	FindGroupUser(ctx context.Context, principal string, shortName string) (*entity.Group, error)
}

type Writer interface {
	CreateGroup(ctx context.Context, institutionID int, opts ...group.Option) (*entity.Group, error)
	UpdateGroup(ctx context.Context, id int, opts ...group.Option) (*entity.Group, error)
	DeleteGroup(ctx context.Context, id int) error
}

type Repository interface {
	Reader
	Writer
}
