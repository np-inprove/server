package seed

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/ent/department"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/hash"
	"github.com/np-inprove/server/internal/logger"
)

// TODO can shift to config in the future
const (
	rootDepartmentName      = "Groot"
	rootDepartmentShortName = "groot"
)

func Exec(ctx context.Context, log logger.AppLogger, cfg *config.Config, client *ent.Client) error {
	instID, err := client.Institution.Query().
		Where(institution.ShortName(cfg.SeedRootInstitutionShortName())).
		FirstID(ctx)
	if err != nil && !apperror.IsEntityNotFound(err) {
		return fmt.Errorf("failed to query for root institution: %w", err)
	}

	// TODO this feels... inefficient
	if apperror.IsEntityNotFound(err) || cfg.SeedForceCreate() {
		log.Info("seeding root institution",
			logger.String("area", "seed"),
			logger.String("institution_name", cfg.SeedRootInstitutionName()),
			logger.String("institution_short_name", cfg.SeedRootInstitutionShortName()),
			logger.Bool("force_create", cfg.SeedForceCreate()),
		)
		instID, err = client.Institution.Create().
			SetName(cfg.SeedRootInstitutionName()).
			SetShortName(cfg.SeedRootInstitutionShortName()).
			OnConflictColumns(institution.FieldShortName).
			UpdateNewValues().
			ID(ctx)
		if err != nil {
			return fmt.Errorf("failed to upsert root institution: %w", err)
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
	if err != nil && !apperror.IsEntityNotFound(err) {
		return fmt.Errorf("failed to query for root department: %w", err)
	}

	// TODO implement cfg.SeedForceCreate() by checking whether Department exists within Institution
	if apperror.IsEntityNotFound(err) {
		log.Info("seeding root department",
			logger.String("area", "seed"),
			logger.String("department_name", rootDepartmentName),
			logger.String("department_short_name", rootDepartmentShortName),
			logger.Int("institution_id", instID),
		)
		dept, err := client.Department.Create().
			SetName(rootDepartmentName).
			SetShortName(rootDepartmentShortName).
			SetInstitutionID(instID).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to upsert root department: %w", err)
		}

		deptID = dept.ID
	}

	_, err = client.User.Query().Where(user.Email(cfg.SeedRootEmail())).FirstID(ctx)
	if err != nil && !apperror.IsEntityNotFound(err) {
		return fmt.Errorf("failed to query for root user: %w", err)
	}

	if apperror.IsEntityNotFound(err) || cfg.SeedForceCreate() {
		p, err := hash.CreateEncoded(cfg.SeedRootPassword())
		if err != nil {
			return fmt.Errorf("failed to hash root user password: %w", err)
		}

		var (
			firstName = "iNProve"
			lastName  = "Administrator"
		)

		log.Info("seeding root user",
			logger.String("area", "seed"),
			logger.String("first_name", firstName),
			logger.String("last_name", lastName),
			logger.String("email", cfg.SeedRootEmail()),
			logger.Int("department_id", deptID),
			logger.Bool("force_create", cfg.SeedForceCreate()),
		)

		err = client.User.Create().
			SetFirstName(firstName).
			SetLastName(lastName).
			SetEmail(cfg.SeedRootEmail()).
			SetPassword(p).
			SetGodMode(true).
			SetDepartmentID(deptID).
			OnConflictColumns(user.FieldEmail).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to upsert root user: %w", err)
		}

		log.Info("root user created",
			logger.String("email", cfg.SeedRootEmail()),
			logger.String("password", cfg.SeedRootPassword()),
		)
	}

	return nil
}
