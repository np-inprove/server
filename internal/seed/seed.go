package seed

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/ent"
	entinstitution "github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/entity/institution"
	"github.com/np-inprove/server/internal/hash"
	"github.com/np-inprove/server/internal/logger"
)

func Exec(ctx context.Context, log logger.AppLogger, cfg *config.Config, client *ent.Client) error {
	instID, err := client.Institution.Query().
		Where(entinstitution.ShortName(cfg.SeedRootInstitutionShortName())).
		FirstID(ctx)
	if err != nil && !apperror.IsNotFound(err) {
		return fmt.Errorf("failed to query for root institution: %w", err)
	}

	if apperror.IsNotFound(err) || cfg.SeedForceCreate() {
		log.Info("seeding root institution",
			logger.String("area", "seed"),
			logger.String("institution_name", cfg.SeedRootInstitutionName()),
			logger.String("institution_short_name", cfg.SeedRootInstitutionShortName()),
			logger.Bool("force_create", cfg.SeedForceCreate()),
		)
		instID, err = client.Institution.Create().
			SetName(cfg.SeedRootInstitutionName()).
			SetShortName(cfg.SeedRootInstitutionShortName()).
			SetDescription("").
			OnConflictColumns(entinstitution.FieldShortName).
			UpdateNewValues().
			ID(ctx)
		if err != nil {
			return fmt.Errorf("failed to upsert root institution: %w", err)
		}
	}

	exists, err := client.User.Query().Where(user.Email(cfg.SeedRootEmail())).Exist(ctx)
	if err != nil && !apperror.IsNotFound(err) {
		return fmt.Errorf("failed to query for root user: %w", err)
	}

	if !exists || cfg.SeedForceCreate() {
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
			logger.Bool("force_create", cfg.SeedForceCreate()),
		)

		err = client.User.Create().
			SetInstitutionID(instID).
			SetRole(institution.RoleAdmin).
			SetFirstName(firstName).
			SetLastName(lastName).
			SetEmail(cfg.SeedRootEmail()).
			SetPassword(p).
			SetGodMode(true).
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
