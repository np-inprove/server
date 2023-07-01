package forum

import (
	"context"
	"fmt"

	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/forum"
	"github.com/np-inprove/server/internal/entity/group"
)

type UseCase interface {
	ListPrincipalForums(ctx context.Context, principal string) ([]*entity.Forum, error)

	// CreateGroup should be an admin only function
	CreateForum(ctx context.Context, principal, name, shortName, description string) (*entity.Forum, error)

	// DeleteGroup should be an admin only function
	DeleteForum(ctx context.Context, principal string, shortName string) error

	// UpdateForum should be an admin only function
	UpdateForum(ctx context.Context, principal string, shortName string, opts ...forum.Option) (*entity.Forum, error)
}

type useCase struct {
	repo Repository
}

func NewUseCase(r Repository) UseCase {
	return useCase{repo: r}
}

func (u useCase) ListPrincipalForums(ctx context.Context, principal string) ([]*entity.Forum, error) {
	return u.repo.FindForumsByUser(ctx, principal)
}

func (u useCase) CreateForum(ctx context.Context, principal, name, shortName, description string) (*entity.Forum, error) {
	usr, err := u.repo.FindGroupUser(ctx, principal, shortName)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if usr.Edges.Group == nil {
		return nil, fmt.Errorf("user edges not loaded")
	}

	if usr.Role == group.RoleMember {
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

func (u useCase) DeleteForum(ctx context.Context, principal string, shortName string) error {
	usr, err := u.repo.FindGroupUser(ctx, principal, shortName)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	if usr.Role == group.RoleMember {
		return ErrUnauthorized
	}

	if err := u.repo.DeleteForum(ctx, usr.GroupID); err != nil {
		return fmt.Errorf("failed to delete forum: %w", err)
	}

	return nil
}

func (u useCase) UpdateForum(ctx context.Context, principal string, shortName string, opts ...forum.Option) (*entity.Forum, error) {
	usr, err := u.repo.FindGroupUser(ctx, principal, shortName)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if usr.Edges.Group == nil {
		return nil, fmt.Errorf("user edges not loaded")
	}

	if usr.Role == group.RoleMember {
		return nil, ErrUnauthorized
	}

	forum, err := u.repo.FindForum(ctx, principal)
	if err != nil {
		return nil, err
	}

	forum, err = u.repo.UpdateForum(ctx, forum.ID, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to update institution: %w", err)
	}

	return forum, nil
}
