package institution

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/institution"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

type UseCase interface {
	ListInstitutions(ctx context.Context) ([]*entity.Institution, error)
	CreateInstitution(ctx context.Context, name, shortName, description string) (*entity.Institution, error)

	// UpdateInstitution modifies an institution which has a short name identified by the principal argument.
	UpdateInstitution(ctx context.Context, principal, name, shortName, description string) (*entity.Institution, error)
	DeleteInstitution(ctx context.Context, shortName string) error

	// ListInviteLinks shows invite links for an institution identified by shortName.
	// If the user (identified by principal), has God mode enabled, then all links will be accessible.
	// Else, only users with a role of Admin can list links for the institution they are associated with.
	ListInviteLinks(ctx context.Context, principal, shortName string) ([]*entity.InstitutionInviteLink, error)
	// CreateInviteLink creates a link an institution identified by shortName.
	// If the user (identified by principal) has God mode enabled, any shortName is accepted.
	// ELse, only users with a role of Admin can create links for the institution they are associated with.
	CreateInviteLink(ctx context.Context, principal, shortName string, role institution.Role) (*entity.InstitutionInviteLink, error)
	// DeleteInviteLink deletes a link.
	// If the user (identified by principal) has God mode enabled, any shortName is accepted.
	// ELse, only users with a role of Admin can delete links for the institution they are associated with.
	DeleteInviteLink(ctx context.Context, principal, shortName, code string) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(r Repository) UseCase {
	return &useCase{repo: r}
}

func (u useCase) ListInstitutions(ctx context.Context) ([]*entity.Institution, error) {
	insts, err := u.repo.FindInstitutions(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find institutions: %w", err)
	}
	return insts, nil
}

func (u useCase) CreateInstitution(ctx context.Context, name, shortName, description string) (*entity.Institution, error) {
	inst, err := u.repo.WithTx(ctx, func(ctx context.Context) (interface{}, error) {
		inst, err := u.repo.CreateInstitution(ctx, name, shortName, description)
		if err != nil {
			if apperror.IsConflict(err) {
				return nil, ErrInstitutionShortNameConflict
			}
			return nil, err
		}

		return inst, nil
	})
	if err != nil {
		return nil, err
	}

	return inst.(*entity.Institution), nil
}

func (u useCase) DeleteInstitution(ctx context.Context, shortName string) error {
	inst, err := u.findInstitution(ctx, shortName)
	if err != nil {
		return err
	}

	err = u.repo.DeleteInstitution(ctx, inst.ID)
	if err != nil {
		return fmt.Errorf("failed to delete institution: %w", err)
	}

	return nil
}

func (u useCase) UpdateInstitution(ctx context.Context, principal, name, shortName, description string) (*entity.Institution, error) {
	inst, err := u.findInstitution(ctx, principal)
	if err != nil {
		return nil, err
	}

	inst, err = u.repo.UpdateInstitution(ctx, inst.ID, name, shortName, description)
	if err != nil {
		return nil, fmt.Errorf("failed to update institution %w", err)
	}

	return inst, nil
}

func (u useCase) ListInviteLinks(ctx context.Context, principal, shortName string) ([]*entity.InstitutionInviteLink, error) {
	err := u.authorizedForInvite(ctx, principal, shortName)
	if err != nil {
		return nil, err
	}

	inst, err := u.repo.FindInstitutionWithInvites(ctx, shortName)
	if err != nil {
		return nil, err
	}

	return inst.Edges.Invites, nil
}

func (u useCase) CreateInviteLink(ctx context.Context, principal, shortName string, role institution.Role) (*entity.InstitutionInviteLink, error) {
	err := u.authorizedForInvite(ctx, principal, shortName)
	if err != nil {
		return nil, err
	}

	inst, err := u.repo.FindInstitution(ctx, shortName)
	if err != nil {
		return nil, fmt.Errorf("failed to find institution")
	}

	code := make([]rune, 8)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}

	link, err := u.repo.CreateInviteLink(ctx, inst.ID, string(code), role)
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

	link, err := u.repo.FindInviteLinkWithInstitution(ctx, code)
	if err != nil {
		return fmt.Errorf("failed to find invite link: %w", err)
	}

	if err := u.repo.DeleteInviteLink(ctx, link.ID); err != nil {
		return fmt.Errorf("failed to delete invite link: %w", err)
	}

	return nil
}

// findInstitution is a helper method to return the correct errors
func (u useCase) findInstitution(ctx context.Context, shortName string) (*entity.Institution, error) {
	inst, err := u.repo.FindInstitution(ctx, shortName)
	if err != nil {
		if apperror.IsNotFound(err) {
			return nil, ErrInstitutionNotFound
		}
		return nil, fmt.Errorf("failed to find institution: %w", err)
	}
	return inst, err
}

func (u useCase) authorizedForInvite(ctx context.Context, principal, shortName string) error {
	usr, err := u.repo.FindUser(ctx, principal)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	if usr.Edges.Institution == nil {
		return fmt.Errorf("invariant")
	}

	if usr.Role != institution.RoleAdmin {
		return ErrUnauthorized
	}

	if !usr.GodMode && usr.Edges.Institution.ShortName != shortName { // If user does not have God mode, check for short name
		return ErrUnauthorized
	}

	return nil
}
