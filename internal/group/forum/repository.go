package forum

import (
	"context"

	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/forum"
)

type Reader interface {
	FindForumsByGroup(ctx context.Context, principal string) ([]*entity.Forum, error)
	FindForumByGroupIDAndShortName(ctx context.Context, groupID int, shortName string) (*entity.Forum, error)
	FindForum(ctx context.Context, shortName string) (*entity.Forum, error)

	FindGroupUser(ctx context.Context, principal, shortName string) (*entity.GroupUser, error)
}

type Writer interface {
	CreateForum(ctx context.Context, forumID int, opts ...forum.Option) (*entity.Forum, error)
	UpdateForum(ctx context.Context, id int, opts ...forum.Option) (*entity.Forum, error)
	DeleteForum(ctx context.Context, id int) error
}

type Repository interface {
	Reader
	Writer
}
