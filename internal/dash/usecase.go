package dash

import (
	"context"
	"fmt"
	"strings"
)

type UseCase interface {
	ListGroups(ctx context.Context, email string) ([]*Group, error)
	ListGroupTypes() ([]*GroupType, error)

	// CreateGroup should be an admin only function
	CreateGroup(ctx context.Context, adminEmail, groupType string, opts ...CreateGroupOption) (*Group, error)
	// DeleteGroup should be an admin only function
	DeleteGroup(ctx context.Context, adminEmail string, path string) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(r Repository) UseCase {
	return useCase{repo: r}
}

func (u useCase) ListGroups(ctx context.Context, email string) ([]*Group, error) {
	return u.repo.FindGroupsByUser(ctx, email)
}

func (u useCase) ListGroupTypes() ([]*GroupType, error) {
	return u.repo.FindGroupTypes()
}

func (u useCase) CreateGroup(ctx context.Context, adminEmail, groupType string, opts ...CreateGroupOption) (*Group, error) {
	if err := GroupTypeValidator(GroupType(groupType)); err != nil {
		return nil, ErrInvalidGroupType
	}

	if ok, err := u.validateAdmin(ctx, adminEmail); err != nil || !ok {
		return nil, err
	}

	var options createGroupOptions
	for _, opt := range opts {
		opt(&options)
	}

	grp, err := u.repo.CreateGroup(ctx, GroupType(groupType), opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create group: %w", err)
	}

	return grp, nil
}

func (u useCase) DeleteGroup(ctx context.Context, adminEmail string, path string) error {
	if ok, err := u.validateAdmin(ctx, adminEmail); err != nil || !ok {
		return err
	}

	if err := u.repo.DeleteGroup(ctx, path); err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}

	return nil
}

func (u useCase) validateAdmin(ctx context.Context, email string) (bool, error) {
	domain := strings.Split(email, "@")[1] // This should not panic
	inst, err := u.repo.FindInstitutionByAdminDomain(ctx, email)
	if err != nil {
		return false, fmt.Errorf("failed to institution by admin domain: %w", err)
	}
	if domain != inst.AdminDomain {
		return false, ErrUserNotAdmin

	}
	return true, nil
}
