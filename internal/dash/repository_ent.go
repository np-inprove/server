package dash

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/entity"
	group2 "github.com/np-inprove/server/internal/entity/group"
)

type entRepository struct {
	client *ent.Client
}

func NewEntRepository(e *ent.Client) Repository {
	return &entRepository{client: e}
}

func (e entRepository) FindInstitutionByAdminDomain(ctx context.Context, domain string) (*entity.Institution, error) {
	inst, err := e.client.Institution.Query().
		Where(institution.AdminDomainEQ(domain)).
		Only(ctx)
	if err != nil {
		if apperror.IsNotFound(err) {
			return nil, ErrInstitutionNotFound
		}
		return nil, err
	}
	return inst, nil
}

func (e entRepository) FindGroupsByUser(ctx context.Context, email string) ([]*entity.Group, error) {
	u, err := e.client.Group.Query().
		Where(
			group.HasUsersWith(
				user.Email(email),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find groups by user email: %w", err)
	}

	return u, nil
}

func (e entRepository) FindGroupTypes() ([]*entity.GroupType, error) {
	sig := group.GroupTypeSpecialInterestGroup
	mod := group.GroupTypeModuleGroup
	cca := group.GroupTypeCca
	mc := group.GroupTypeMentorClass
	return []*entity.GroupType{
		&sig, &mod, &cca, &mc,
	}, nil
}

func (e entRepository) CreateGroup(ctx context.Context, groupType entity.GroupType, opts ...group2.Option) (*entity.Group, error) {
	var options group2.Options
	for _, opt := range opts {
		opt(&options)
	}

	g, err := e.client.Group.
		Create().
		SetGroupType(groupType).
		SetName(options.Name).
		SetPath(options.Path).
		SetDescription(options.Description).
		Save(ctx)
	if err != nil {
		if apperror.IsConflict(err) {
			return nil, ErrGroupPathConflict
		}
		return nil, fmt.Errorf("failed to create group: %w", err)
	}

	return g, nil
}

func (e entRepository) DeleteGroup(ctx context.Context, path string) error {
	_, err := e.client.Group.Delete().Where(group.Path(path)).Exec(ctx)
	return err
}
