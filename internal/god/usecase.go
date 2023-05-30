package god

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/apperror"
)

type UseCase interface {
	ListInstitutions(ctx context.Context) ([]*Institution, error)
	CreateInstitution(
		ctx context.Context,
		name string,
		shortName string,
		adminDomain string,
		studentDomain string,
	) (*Institution, error)
	DeleteInstitution(ctx context.Context, shortName string) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(r Repository) UseCase {
	return &useCase{repo: r}
}

func (u useCase) ListInstitutions(ctx context.Context) ([]*Institution, error) {
	insts, err := u.repo.FindInstitutions(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find institutions: %w", err)
	}
	return insts, nil
}

func (u useCase) CreateInstitution(
	ctx context.Context,
	name string,
	shortName string,
	adminDomain string,
	studentDomain string,
) (*Institution, error) {
	// TODO optimizations
	err := u.repo.WithTx(ctx, func(ctx context.Context) error {
		_, err := u.repo.CreateInstitution(ctx, name, shortName, adminDomain, studentDomain)
		if err != nil {
			if apperror.IsConflict(err) {
				return ErrShortNameConflict
			}
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	insts, err := u.repo.FindInstitutions(ctx)
	if err != nil {
		return nil, err
	}

	var inst *Institution
	for _, i := range insts {
		if i.ShortName == shortName {
			inst = i
			break
		}
	}

	return inst, nil
}

func (u useCase) DeleteInstitution(ctx context.Context, shortName string) error {
	return u.repo.WithTx(ctx, func(ctx context.Context) error {
		return u.repo.DeleteInstitution(ctx, shortName)
	})
}
