package institution

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/entity"
)

type UseCase interface {
	ListInstitutions(ctx context.Context) ([]*entity.Institution, error)
	CreateInstitution(ctx context.Context, name string, shortName string, description string) (*entity.Institution, error)
	DeleteInstitution(ctx context.Context, shortName string) error
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

type T *entity.Institution

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
	inst, err := u.repo.FindInstitution(ctx, shortName)
	if err != nil {
		return fmt.Errorf("failed to find institution: %w", err)
	}

	err = u.repo.DeleteInstitution(ctx, inst.ID)
	if err != nil {
		return fmt.Errorf("failed to delete institution: %w", err)
	}

	return nil
}
