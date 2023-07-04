package group

import (
	"context"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/group"
	"github.com/np-inprove/server/internal/transactor"
)

type Reader interface {
	FindGroupsByUser(ctx context.Context, principal string) ([]*entity.Group, error)
	FindGroupByInstitutionIDAndShortName(ctx context.Context, institutionID int, shortName string) (*entity.Group, error)

	FindUserWithInstitution(ctx context.Context, principal string) (*entity.User, error)
	FindGroupUser(ctx context.Context, principal string, shortName string) (*entity.GroupUser, error)
}

type Writer interface {
	CreateGroup(ctx context.Context, institutionID int, opts ...group.Option) (*entity.Group, error)
	UpdateGroup(ctx context.Context, id int, opts ...group.Option) (*entity.Group, error)
	DeleteGroup(ctx context.Context, id int) error
	CreateGroupUser(ctx context.Context, userID int, groupID int, role group.Role) (*entity.GroupUser, error)
	BulkDeleteGroupUsers(ctx context.Context, groupID int) error
}

type Repository interface {
	Reader
	Writer
	transactor.Transactor
}
