package forum

import (
	"context"
	"fmt"

	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/forum"
	"github.com/np-inprove/server/internal/entity/group"
)

type UseCase interface {
	// groupShortName is the shortname used by groups, shortName us the shortname used by forums
	ListForums(ctx context.Context, principal string, groupShortName string) ([]*entity.Forum, error)

	// CreateGroup should be an educator + owner only function
	CreateForum(ctx context.Context, principal, groupShortName, name, shortName, description string) (*entity.Forum, error)

	// DeleteGroup should be an educator + owner only function
	DeleteForum(ctx context.Context, principal string, groupShortName string) error

	// UpdateForum should be an educator + owner only function
	UpdateForum(ctx context.Context, principal string, groupShortName, name, shortName, description string) (*entity.Forum, error)
}

type useCase struct {
	repo Repository
}

func NewUseCase(r Repository) UseCase {
	return useCase{repo: r}
}

func (u useCase) ListForums(ctx context.Context, principal string, groupShortName string) ([]*entity.Forum, error) {
	return u.repo.FindForumsByGroup(ctx, principal)
}

func (u useCase) CreateForum(ctx context.Context, principal, groupShortName, name, shortName, description string) (*entity.Forum, error) {
	usr, err := u.repo.FindGroupUser(ctx, principal, groupShortName)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if usr.Edges.Group == nil {
		return nil, fmt.Errorf("user edges not loaded")
	}

	if usr.Role != group.RoleEducator || usr.Role != group.RoleOwner {
		return nil, ErrUnauthorized
	}

	grp := usr.Edges.Group
	if _, err := u.repo.FindForumByGroupIDAndShortName(ctx, grp.ID, shortName); err == nil {
		return nil, ErrForumShortNameConflict
	}

	forum, err := u.repo.CreateForum(ctx, grp.ID, forum.Name(name), forum.ShortName(shortName), forum.Description(description))
	if err != nil {
		return nil, fmt.Errorf("failed to create forum: %w", err)
	}

	return forum, nil
}

func (u useCase) DeleteForum(ctx context.Context, principal string, groupShortName string) error {
	usr, err := u.repo.FindGroupUser(ctx, principal, groupShortName)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	if usr.Role != group.RoleEducator || usr.Role != group.RoleOwner {
		return ErrUnauthorized
	}

	if err := u.repo.DeleteForum(ctx, usr.GroupID); err != nil {
		return fmt.Errorf("failed to delete forum: %w", err)
	}

	return nil
}

func (u useCase) UpdateForum(ctx context.Context, principal string, groupShortName, name, shortName, description string) (*entity.Forum, error) {
	usr, err := u.repo.FindGroupUser(ctx, principal, groupShortName)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if usr.Edges.Group == nil {
		return nil, fmt.Errorf("user edges not loaded")
	}

	if usr.Role != group.RoleEducator || usr.Role != group.RoleOwner {
		return nil, ErrUnauthorized
	}

	forum, err := u.repo.FindForum(ctx, principal)
	if err != nil {
		return nil, err
	}

	forum, err = u.repo.UpdateForum(ctx, forum.ID, name, shortName, description)
	if err != nil {
		return nil, fmt.Errorf("failed to update institution: %w", err)
	}

	return forum, nil
}
