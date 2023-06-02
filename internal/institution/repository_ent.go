package institution

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entutils"
	"github.com/np-inprove/server/internal/logger"
)

type entRepository struct {
	log    logger.AppLogger
	client *ent.Client
}

func NewEntRepository(l logger.AppLogger, c *ent.Client) Repository {
	return entRepository{l, c}
}

func (e entRepository) FindInstitutions(ctx context.Context) ([]*entity.Institution, error) {
	i, err := e.client.Institution.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all institutions: %w", err)
	}
	return i, nil
}

func (e entRepository) CreateInstitution(ctx context.Context, name string, shortName string, adminDomain string, studentDomain string) (*entity.Institution, error) {
	c := e.client
	if cc, ok := entutils.ExtractTx(ctx); ok {
		c = cc
	}

	inst, err := c.Institution.Create().
		SetName(name).
		SetShortName(shortName).
		SetAdminDomain(adminDomain).
		SetStudentDomain(studentDomain).
		Save(ctx)
	if err != nil {
		if apperror.IsConflict(err) {
			return nil, ErrInstitutionShortNameConflict
		}
		return nil, fmt.Errorf("failed to save institution: %w", err)
	}

	return inst, nil
}

func (e entRepository) DeleteInstitution(ctx context.Context, shortName string) error {
	c := e.client
	if cc, ok := entutils.ExtractTx(ctx); ok {
		c = cc
	}

	id, err := c.Institution.Query().Where(institution.ShortName(shortName)).OnlyID(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete institution as it does not exist: %w", err)
	}

	err = c.Institution.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete institution: %w", err)
	}

	return nil
}

func (e entRepository) WithTx(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := e.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start ent transaction: %w", err)
	}

	txc := tx.Client()
	ctx = context.WithValue(ctx, entutils.EntTxCtxKey, txc)

	if err := f(ctx); err != nil {
		e.log.Warn("failed database query during ent transaction",
			logger.String("err", err.Error()),
			logger.String("area", "god"),
		)
		if err2 := tx.Rollback(); err2 != nil {
			e.log.Error("failed ent transaction rollback",
				logger.String("err", err.Error()),
				logger.String("causer", err.Error()),
				logger.String("area", "god"),
			)
			return fmt.Errorf("failed to rollback ent transaction: %w", err2)
		}
		return err
	}

	return tx.Commit()
}
