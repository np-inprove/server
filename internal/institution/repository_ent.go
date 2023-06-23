package institution

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/ent/institutioninvitelink"
	"github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/entity/institution"

	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/ent"
	entinstitution "github.com/np-inprove/server/internal/ent/institution"
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
	inst, err := e.client.Institution.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all institutions: %w", err)
	}
	return inst, nil
}

func (e entRepository) FindInstitution(ctx context.Context, shortName string) (*entity.Institution, error) {
	inst, err := e.client.Institution.Query().Where(entinstitution.ShortName(shortName)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find institution: %w", err)
	}
	return inst, nil
}

func (e entRepository) CreateInstitution(ctx context.Context, name, shortName, description string) (*entity.Institution, error) {
	c := e.client
	if cc, ok := entutils.ExtractTx(ctx); ok {
		c = cc
	}

	inst, err := c.Institution.Create().
		SetName(name).
		SetShortName(shortName).
		SetDescription(description).
		Save(ctx)
	if err != nil {
		if apperror.IsConflict(err) {
			return nil, ErrInstitutionShortNameConflict
		}
		return nil, fmt.Errorf("failed to save institution: %w", err)
	}

	return inst, nil
}

func (e entRepository) DeleteInstitution(ctx context.Context, id int) error {
	c := e.client
	if cc, ok := entutils.ExtractTx(ctx); ok {
		c = cc
	}

	err := c.Institution.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete institution: %w", err)
	}

	return nil
}

func (e entRepository) UpdateInstitution(ctx context.Context, id int, name, shortName, description string) (*entity.Institution, error) {
	c := e.client
	if cc, ok := entutils.ExtractTx(ctx); ok {
		c = cc
	}

	inst, err := c.Institution.UpdateOneID(id).
		SetName(name).
		SetShortName(shortName).
		SetDescription(description).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update institution: %w", err)
	}

	return inst, nil
}

func (e entRepository) WithTx(
	ctx context.Context,
	fn func(ctx context.Context) (interface{}, error),
) (interface{}, error) {
	tx, err := e.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start ent transaction: %w", err)
	}

	txc := tx.Client()
	ctx = context.WithValue(ctx, entutils.EntTxCtxKey, txc)

	ret, err := fn(ctx)
	if err != nil {
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
			return nil, fmt.Errorf("failed to rollback ent transaction: %w", err2)
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (e entRepository) FindInstitutionWithInvites(ctx context.Context, shortName string) (*entity.Institution, error) {
	inst, err := e.client.Institution.Query().Where(entinstitution.ShortName(shortName)).WithInvites().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find institution: %w", err)
	}
	return inst, nil
}

func (e entRepository) FindUserWithInstitution(ctx context.Context, principal string) (*entity.User, error) {
	usr, err := e.client.User.Query().
		Where(
			user.Email(principal),
		).
		WithInstitution().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find user with institution: %w", err)
	}

	return usr, nil
}

func (e entRepository) FindInviteLinks(ctx context.Context, id int) ([]*entity.InstitutionInviteLink, error) {
	links, err := e.client.Institution.Query().Where(entinstitution.ID(id)).QueryInvites().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find institution invite links: %w", err)
	}
	return links, nil
}

func (e entRepository) FindInviteLink(ctx context.Context, id int, code string) (*entity.InstitutionInviteLink, error) {
	link, err := e.client.Institution.Query().
		Where(
			entinstitution.ID(id),
		).
		QueryInvites().
		Where(
			institutioninvitelink.Code(code),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find institution invite link: %w", err)
	}
	return link, nil
}

func (e entRepository) CreateInviteLink(ctx context.Context, id int, code string, role institution.Role) (*entity.InstitutionInviteLink, error) {
	link, err := e.client.InstitutionInviteLink.Create().
		SetInstitutionID(id).
		SetCode(code).
		SetRole(role).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create institution invite link: %w", err)
	}
	return link, nil
}

func (e entRepository) DeleteInviteLink(ctx context.Context, id int) error {
	err := e.client.InstitutionInviteLink.DeleteOneID(id)
	if err != nil {
		return fmt.Errorf("failed to delete institution invite link: %w", err)
	}
	return nil
}
