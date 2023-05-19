package seed

import (
	"context"
	"errors"
	"fmt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/auth"
	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/ent/department"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/user"
)

// TODO - can shift to config in the future
const (
	rootDepartmentName      = "Groot"
	rootDepartmentShortName = "groot"
)

func Exec(ctx context.Context, cfg *config.Config, client *ent.Client) error {
	instID, err := client.Institution.Query().
		Where(institution.ShortName(cfg.SeedRootInstitutionShortName())).
		FirstID(ctx)
	if err != nil && !errors.As(err, &apperror.ErrEntityNotFound) {
		return fmt.Errorf("failed to query for root institution: %v", err)
	}

	// TODO: this feels... inefficient
	if errors.As(err, &apperror.ErrEntityNotFound) || cfg.SeedForceCreate() {
		instID, err = client.Institution.Create().
			SetName(cfg.SeedRootInstitutionName()).
			SetShortName(cfg.SeedRootInstitutionShortName()).
			OnConflictColumns(institution.FieldShortName).
			UpdateNewValues().
			ID(ctx)
		if err != nil {
			return fmt.Errorf("failed to upsert root institution: %v", err)
		}
	}

	deptID, err := client.Department.Query().
		Where(
			department.ShortName(rootDepartmentShortName),
			department.HasInstitutionWith(
				institution.ID(instID),
			),
		).
		FirstID(ctx)
	if err != nil && !errors.As(err, &apperror.ErrEntityNotFound) {
		return fmt.Errorf("failed to query for root department: %v", err)
	}

	// Department will not be upserted as there is no way to determine
	// conflict within an institution
	if errors.As(err, &apperror.ErrEntityNotFound) {
		dept, err := client.Department.Create().
			SetName(rootDepartmentName).
			SetShortName(rootDepartmentShortName).
			SetInstitutionID(instID).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to upsert root department: %v", err)
		}

		deptID = dept.ID
	}
	//else if cfg.SeedForceCreate() {
	//	client.Department.Update().SetName(rootDepartmentName).SetShortName(rootDepartmentShortName).SetInstitutionID(instID).
	//}

	_, err = client.User.Query().Where(user.Email(cfg.SeedRootEmail())).FirstID(ctx)
	if err != nil && !errors.As(err, &apperror.ErrEntityNotFound) {
		return fmt.Errorf("failed to query for root user: %v", err)
	}

	if errors.As(err, &apperror.ErrEntityNotFound) || cfg.SeedForceCreate() {
		p, err := auth.HashPassword(cfg.SeedRootPassword())
		if err != nil {
			return fmt.Errorf("failed to hash root user password: %v", err)
		}

		err = client.User.Create().
			SetFirstName("iNProve").
			SetLastName("Administrator").
			SetEmail(cfg.SeedRootEmail()).
			SetPasswordHash(p).
			SetGodMode(true).
			SetDepartmentID(deptID).
			OnConflictColumns(user.FieldEmail).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to upsert root user: %v", err)
		}
	}

	return nil
}
