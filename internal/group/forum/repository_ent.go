package forum

import (
	"context"
	"fmt"

	"github.com/np-inprove/server/internal/ent"
	entforum "github.com/np-inprove/server/internal/ent/forum"
	entgroup "github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/groupuser"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/forum"
)

type entRepository struct {
	client *ent.Client
}

func NewEntRepository(e *ent.Client) Repository {
	return &entRepository{client: e}
}

func (e entRepository) FindForumsByUser(ctx context.Context, principal string) ([]*entity.Forum, error) {
	forum, err := e.client.Forum.Query().
		Where(
			entforum.HasGroupWith(
				predicate.Group(user.Email(principal)),
			)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find forums by user email: %w", err)
	}

	return forum, nil
}

func (e entRepository) FindForumByGroupIDAndShortName(ctx context.Context, groupID int, shortName string) (*entity.Forum, error) {
	forum, err := e.client.Forum.Query().
		Where(
			entforum.HasGroupWith(
				entgroup.ID(groupID),
			),
			entforum.ShortName(shortName),
		).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find forums by user email: %w", err)
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

func (e entRepository) FindUserWithGroups(ctx context.Context, principal string) (*entity.User, error) {
	usr, err := e.client.User.Query().
		Where(
			user.Email(principal),
		).WithGroups().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find user with group: %w", err)
	}

	return usr, nil
}

func (e entRepository) FindGroupUser(ctx context.Context, principal string, shortName string) (*entity.GroupUser, error) {
	grpusr, err := e.client.GroupUser.Query().
		Where(
			groupuser.HasUserWith(user.Email(principal)),
			groupuser.HasGroupWith(entgroup.ShortName(shortName)),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find group user: %w", err)
	}

	return grpusr, nil
}

func (e entRepository) CreateForum(ctx context.Context, groupID int, opts ...forum.Option) (*entity.Forum, error) {
	var options forum.Options
	for _, opt := range opts {
		opt(&options)
	}

	forum, err := e.client.Forum.
		Create().
		SetName(options.Name).
		SetShortName(options.ShortName).
		SetDescription(options.Description).
		SetGroupID(groupID).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create group: %w", err)
	}
	return forum, nil
}

func (e entRepository) UpdateForum(ctx context.Context, id int, name, shortName, description string) (*entity.Forum, error) {
	forum, err := e.client.Forum.UpdateOneID(id).
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
