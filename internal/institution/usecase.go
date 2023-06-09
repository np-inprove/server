package institution

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/entity"
	"strings"
)

type UseCase interface {
	ListInstitutions(ctx context.Context) ([]*entity.Institution, error)
	CreateInstitution(
		ctx context.Context,
		name string,
		shortName string,
		adminDomain string,
		studentDomain string,
	) (*entity.Institution, error)
	DeleteInstitution(ctx context.Context, shortName string) error

	// CreateDepartment should be an admin only function
	CreateDepartment(ctx context.Context, adminEmail, departmentName string, departmentShortName string) (*entity.Department, error)
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

func (u useCase) CreateInstitution(
	ctx context.Context,
	name string,
	shortName string,
	adminDomain string,
	studentDomain string,
) (*entity.Institution, error) {
	// TODO optimizations
	err := u.repo.WithTx(ctx, func(ctx context.Context) error {
		_, err := u.repo.CreateInstitution(ctx, name, shortName, adminDomain, studentDomain)
		if err != nil {
			if apperror.IsConflict(err) {
				return ErrInstitutionShortNameConflict
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

	var inst *entity.Institution
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

func (u useCase) CreateDepartment(ctx context.Context, adminEmail, departmentName string, departmentShortName string) (*entity.Department, error) {
	if ok, err := u.validateAdmin(ctx, adminEmail); err != nil || !ok {
		return nil, err
	}

	inst, err := u.getInstitution(ctx, adminEmail)
	if err != nil {
		return nil, err
	}

	dep, err := u.repo.CreateDepartment(ctx, departmentName, departmentShortName, inst.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create department: %w", err)
	}

	return dep, nil
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

func (u useCase) getInstitution(ctx context.Context, email string) (*entity.Institution, error) {
	domain := strings.Split(email, "@")[1] // This should not panic
	inst, err := u.repo.FindInstitutionByAdminDomain(ctx, domain)
	if err != nil {
		return nil, fmt.Errorf("failed to institution by admin domain: %w", err)
	}
	return inst, nil
}
