package group

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/group"
	"github.com/np-inprove/server/internal/entity/institution"
)

type UseCase interface {
	ListPrincipalGroups(ctx context.Context, principal string) ([]*entity.Group, error)

	// CreateGroup should be an admin only function
	CreateGroup(ctx context.Context, principal string, opts ...group.Option) (*entity.Group, error)

	// DeleteGroup should be an admin only function
	DeleteGroup(ctx context.Context, principal string, shortName string) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(r Repository) UseCase {
	return useCase{repo: r}
}

func (u useCase) ListPrincipalGroups(ctx context.Context, principal string) ([]*entity.Group, error) {
	return u.repo.FindGroupsByUser(ctx, principal)
}

func (u useCase) CreateGroup(ctx context.Context, principal string, opts ...group.Option) (*entity.Group, error) {
	usr, err := u.repo.FindUserWithInstitution(ctx, principal)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if usr.Edges.Institution == nil {
		return nil, fmt.Errorf("user edges not loaded")
	}

	if usr.Role != institution.RoleAdmin {
		return nil, ErrUnauthorized
	}

	var options group.Options
	for _, opt := range opts {
		opt(&options)
	}

	inst := usr.Edges.Institution
	if _, err := u.repo.FindGroupByInstitutionIDAndShortName(ctx, inst.ID, options.ShortName); err == nil {
		return nil, ErrGroupShortNameConflict
	}

	grp, err := u.repo.CreateGroup(ctx, inst.ID, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create group: %w", err)
	}

	return grp, nil
}

func (u useCase) DeleteGroup(ctx context.Context, principal string, shortName string) error {
	grpusr, err := u.repo.FindGroupUser(ctx, principal, shortName)
	if err != nil {
		return fmt.Errorf("failed to find group user: %w", err)
	}

	if grpusr.Role != group.RoleOwner {
		return ErrUnauthorized
	}

	if err := u.repo.DeleteGroup(ctx, grpusr.GroupID); err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}

	return nil
}
