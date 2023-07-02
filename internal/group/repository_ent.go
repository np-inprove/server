package group

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/ent"
	entgroup "github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/groupuser"
	entinstitution "github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/group"
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

func (e entRepository) FindGroupsByUser(ctx context.Context, principal string) ([]*entity.Group, error) {
	grps, err := e.client.Group.Query().
		Where(
			entgroup.HasUsersWith(
				user.Email(principal),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find groups by user email: %w", err)
	}

	return grps, nil
}

func (e entRepository) FindGroupByInstitutionIDAndShortName(ctx context.Context, institutionID int, shortName string) (*entity.Group, error) {
	grp, err := e.client.Group.Query().
		Where(
			entgroup.HasInstitutionWith(
				entinstitution.ID(institutionID),
			),
			entgroup.ShortName(shortName),
		).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find group by institution id and short name: %w", err)
	}

	return grp, nil
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

func (e entRepository) FindGroupUser(ctx context.Context, principal string, shortName string) (*entity.GroupUser, error) {
	grpusr, err := e.client.GroupUser.Query().
		Where(
			groupuser.HasUserWith(user.Email(principal)),
			groupuser.HasGroupWith(entgroup.ShortName(shortName)),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find group user: %w", err)
	}

	return grpusr, nil
}

func (e entRepository) CreateGroup(ctx context.Context, institutionID int, opts ...group.Option) (*entity.Group, error) {

	c := e.client
	if cc, ok := entutils.ExtractTx(ctx); ok {
		c = cc
	}

	var options group.Options
	for _, opt := range opts {
		opt(&options)
	}

	g, err := c.Group.
		Create().
		SetName(options.Name).
		SetShortName(options.ShortName).
		SetDescription(options.Description).
		SetInstitutionID(institutionID).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create group: %w", err)
	}

	return g, nil
}

func (e entRepository) CreateGroupUser(ctx context.Context, userID int, groupID int, role group.Role) (*entity.GroupUser, error) {

	c := e.client
	if cc, ok := entutils.ExtractTx(ctx); ok {
		c = cc
	}

	groupUser, err := c.GroupUser.
		Create().
		SetUserID(userID).
		SetGroupID(groupID).
		SetRole(role).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to add owner to group: %w", err)
	}

	return groupUser, nil
}

func (e entRepository) UpdateGroup(ctx context.Context, id int, opts ...group.Option) (*entity.Group, error) {

	c := e.client
	if cc, ok := entutils.ExtractTx(ctx); ok {
		c = cc
	}

	var options group.Options
	for _, opt := range opts {
		opt(&options)
	}

	query := c.Group.UpdateOneID(id)

	if options.Name != "" {
		query.SetName(options.Name)
	}

	if options.ShortName != "" {
		query.SetShortName(options.ShortName)
	}

	if options.Description != "" {
		query.SetDescription(options.Description)
	}

	grp, err := query.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update group: %w", err)
	}

	return grp, nil
}

func (e entRepository) DeleteGroup(ctx context.Context, id int) error {

	c := e.client
	if cc, ok := entutils.ExtractTx(ctx); ok {
		c = cc
	}

	err := c.Group.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}
	return err
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
