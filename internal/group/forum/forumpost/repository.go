package forumpost

import (
	"context"

	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/forumpost"
)

type Reader interface {
	FindForumPostsByUser() ([]*entity.ForumPost, error)
	FindForumPost() (*entity.ForumPost, error)

	FindGroupUser(ctx context.Context, principal, shortName string) (*entity.GroupUser, error)
}

type Writer interface {
	CreateForumPost(ctx context.Context, forumID int, opts ...forumpost.Option) (*entity.ForumPost, error)
}

type Repository interface {
	Reader
	Writer
}
