package god

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/entutils"
	"github.com/np-inprove/server/internal/logger"
)

type entRepository struct {
	l logger.AppLogger
	c *ent.Client
}

func NewEntRepository(l logger.AppLogger, c *ent.Client) Repository {
	return entRepository{l, c}
}

func (e entRepository) FindInstitutions(ctx context.Context) ([]*Institution, error) {
	i, err := e.c.Institution.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all institutions: %w", err)
	}
	return i, nil
}

func (e entRepository) CreateInstitution(ctx context.Context, name string, shortName string, adminDomain string, studentDomain string) (*Institution, error) {
	client := e.c
	if cc, ok := entutils.ExtractTx(ctx); ok {
		*client = *cc
	}

	inst, err := client.Institution.Create().
		SetName(name).
		SetShortName(shortName).
		SetAdminDomain(adminDomain).
		SetStudentDomain(studentDomain).
		Save(ctx)
	if err != nil {
		if apperror.IsConflict(err) {
			return nil, ErrShortNameConflict
		}
		return nil, fmt.Errorf("failed to save institution: %w", err)
	}

	return inst, nil
}

func (e entRepository) DeleteInstitution(ctx context.Context, shortName string) error {
	client := e.c
	if cc, ok := entutils.ExtractTx(ctx); ok {
		*client = *cc
	}

	id, err := client.Institution.Query().Where(institution.ShortName(shortName)).OnlyID(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete institution as it does not exist: %w", err)
	}

	err = client.Institution.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete institution: %w", err)
	}

	return nil
}

func (e entRepository) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := e.c.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start ent transaction: %w", err)
	}

	txc := tx.Client()
	ctx = context.WithValue(ctx, entutils.EntTxCtxKey, txc)

	if err := fn(ctx); err != nil {
		e.l.Warn("failed database query during ent transaction",
			logger.String("err", err.Error()),
			logger.String("area", "god"),
		)
		if err2 := tx.Rollback(); err2 != nil {
			e.l.Error("failed ent transaction rollback",
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
