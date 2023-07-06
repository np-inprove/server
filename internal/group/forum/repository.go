package forum

import (
	"context"

	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/forum"
)

type Reader interface {
	FindForumsByGroupAndInstitution(ctx context.Context, principal string, instShortName string) ([]*entity.Forum, error)
	FindForumByGroupIDAndShortName(ctx context.Context, groupID int, shortName string) (*entity.Forum, error)
	FindForum(ctx context.Context, shortName string) (*entity.Forum, error)

	FindUserWithGroups(ctx context.Context, principal string) (*entity.User, error)
	FindGroupUserWithGroup(ctx context.Context, principal, shortName string) (*entity.GroupUser, error)
}

type Writer interface {
	CreateForum(ctx context.Context, groupID int, opts ...forum.Option) (*entity.Forum, error)
	UpdateForum(ctx context.Context, forumID int, name, shortName, description string) (*entity.Forum, error)
	DeleteForum(ctx context.Context, id int) error
}

type Repository interface {
	Reader
	Writer
}
