package group

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/group"
	"github.com/np-inprove/server/internal/entity/institution"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

type UseCase interface {
	ListPrincipalGroups(ctx context.Context, principal string) ([]*entity.Group, error)

	// CreateGroup should be an educator and admin only feature
	CreateGroup(ctx context.Context, principal string, opts ...group.Option) (*entity.Group, error)

	// UpdateGroup should be an educator and admin only function
	UpdateGroup(ctx context.Context, principal string, shortName string, opts ...group.Option) (*entity.Group, error)

	// DeleteGroup should be an educator and admin only function
	DeleteGroup(ctx context.Context, principal string, shortName string) error

	// ListInviteLinks shows invite links for a group identified by shortName.
	ListInviteLinks(ctx context.Context, principal, shortName string) ([]*entity.GroupInviteLink, error)

	// CreateInviteLink should be an owner and educator only function
	CreateInviteLink(ctx context.Context, principal, shortName string, role group.Role) (*entity.GroupInviteLink, error)

	// DeleteInviteLink should be an owner and educator only function
	DeleteInviteLink(ctx context.Context, principal, shortName, code string) error
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

	if usr.Role != institution.RoleAdmin && usr.Role != institution.RoleEducator {
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

	grp, err := u.repo.WithTx(ctx, func(ctx context.Context) (interface{}, error) {
		grp, err := u.repo.CreateGroup(ctx, inst.ID, opts...)
		if err != nil {
			return nil, fmt.Errorf("failed to create group: %w", err)
		}

		_, err = u.repo.CreateGroupUser(ctx, usr.ID, grp.ID, group.RoleOwner)
		if err != nil {
			return nil, fmt.Errorf("failed to assign group owner: %w", err)
		}

		return grp, nil
	})

	if err != nil {
		return nil, err
	}

	return grp.(*entity.Group), nil
}

func (u useCase) UpdateGroup(ctx context.Context, principal string, shortName string, opts ...group.Option) (*entity.Group, error) {
	grpusr, err := u.repo.FindGroupUser(ctx, principal, shortName)
	if err != nil {
		return nil, fmt.Errorf("failed to find group user: %w", err)
	}

	if grpusr.Role != group.RoleOwner && grpusr.Role != group.RoleEducator {
		return nil, ErrUnauthorized
	}

	grp, err := u.repo.UpdateGroup(ctx, grpusr.GroupID, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to update group: %w", err)
	}

	return grp, nil
}

func (u useCase) DeleteGroup(ctx context.Context, principal string, shortName string) error {
	grpusr, err := u.repo.FindGroupUser(ctx, principal, shortName)
	if err != nil {
		return fmt.Errorf("failed to find group user: %w", err)
	}

	if grpusr.Role != group.RoleOwner && grpusr.Role != group.RoleEducator {
		return ErrUnauthorized
	}

	if err := u.repo.BulkDeleteGroupUsers(ctx, grpusr.GroupID); err != nil {
		return fmt.Errorf("failed to delete group users: %w", err)
	}

	if err := u.repo.DeleteGroup(ctx, grpusr.GroupID); err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}

	return nil
}

func (u useCase) authorizedForInvite(ctx context.Context, principal, shortName string) error {
	usr, err := u.repo.FindGroupUser(ctx, principal, shortName)
	if err != nil {
		return fmt.Errorf("failed to find group: %w", err)
	}

	if usr.Role != group.RoleOwner && usr.Role != group.RoleEducator {
		return ErrUnauthorized
	}

	return nil
}

func (u useCase) ListInviteLinks(ctx context.Context, principal, shortName string) ([]*entity.GroupInviteLink, error) {
	err := u.authorizedForInvite(ctx, principal, shortName)
	if err != nil {
		return nil, err
	}

	grp, err := u.repo.FindGroupWithInvites(ctx, shortName)
	if err != nil {
		return nil, err
	}

	return grp.Edges.Invites, nil
}

func (u useCase) CreateInviteLink(ctx context.Context, principal, shortName string, role group.Role) (*entity.GroupInviteLink, error) {
	err := u.authorizedForInvite(ctx, principal, shortName)
	if err != nil {
		return nil, err
	}

	grp, err := u.repo.FindGroup(ctx, shortName)
	if err != nil {
		return nil, fmt.Errorf("failed to find group")
	}

	code := make([]rune, 8)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}

	link, err := u.repo.CreateInviteLink(ctx, grp.ID, string(code), role)
	if err != nil {
		return nil, fmt.Errorf("failed to create invite link: %w", err)
	}

	return link, nil
}

func (u useCase) DeleteInviteLink(ctx context.Context, principal, shortName, code string) error {
	err := u.authorizedForInvite(ctx, principal, shortName)
	if err != nil {
		return err
	}

	link, err := u.repo.FindInviteWithGroup(ctx, code)
	if err != nil {
		return fmt.Errorf("failed to find invite link: %w", err)
	}

	if err := u.repo.DeleteInviteLink(ctx, link.ID); err != nil {
		return fmt.Errorf("failed to delete invite link: %w", err)
	}

	return nil
}
