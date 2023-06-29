package forum

import (
	"context"
	"fmt"

	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/ent"
	entforum "github.com/np-inprove/server/internal/ent/forum"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entutils"
	"github.com/np-inprove/server/internal/logger"
)

type entRepository struct {
	log    logger.AppLogger
	client *ent.Client
}

func NewEntRepository(l logger.AppLogger, c *ent.Client) Repository {
	return entRepository{l, c}
}

func (e entRepository) FindForums(ctx context.Context) ([]*entity.Forum, error) {
	forum, err := e.client.Forum.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all forums: %w", err)
	}
	return forum, nil
}

func (e entRepository) FindForum(ctx context.Context, shortName string) (*entity.Forum, error) {
	inst, err := e.client.Forum.Query().Where(entforum.ShortName(shortName)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find forum: %w", err)
	}
	return inst, nil
}

func (e entRepository) CreateForum(ctx context.Context, name, shortName, description string) (*entity.Forum, error) {
	c := e.client
	if cc, ok := entutils.ExtractTx(ctx); ok {
		c = cc
	}

	forum, err := c.Forum.Create().
		SetName(name).
		SetShortName(shortName).
		SetDescription(description).
		Save(ctx)
	if err != nil {
		if apperror.IsConflict(err) {
			return nil, ErrForumShortNameConflict //create error.go file
		}
		return nil, fmt.Errorf("failed to save forum: %w", err)
	}

	return forum, nil
}

func (e entRepository) UpdateForum(ctx context.Context, id int, name, shortName, description string) (*entity.Forum, error) {
	c := e.client
	if cc, ok := entutils.ExtractTx(ctx); ok {
		c = cc
	}

	forum, err := c.Forum.UpdateOneID(id).
		SetName(name).
		SetShortName(shortName).
		SetDescription(description).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update forum: %w", err)
	}

	return forum, nil
}

func (e entRepository) DeleteForum(ctx context.Context, id int) error {
	err := e.client.Forum.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete forum: %w", err)
	}
	return err
}

func (e entRepository) FindUserWithForum()
