package forum

import (
	"context"

	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/forum"
)

type Reader interface {
	FindForumsByUser(ctx context.Context, principal string) ([]*entity.Forum, error)
	FindForumByGroupIDAndShortName(ctx context.Context, institutionID int, shortName string) (*entity.Forum, error)

	FindUserWithForum(ctx context.Context, principal string) (*entity.User, error)
}

type Writer interface {
	CreateForum(ctx context.Context, forumID int, opts ...forum.Options) (*entity.Forum, error)
	EditForum(ctx context.Context, id int, opts ...forum.Options) (*entity.Forum, error)
	DeleteForum(ctx context.Context, id int) error
}

type Repository interface {
	Reader
	Writer
}
