package dash

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/user"
)

type entRepository struct {
	ent *ent.Client
}

func NewEntRepository(e *ent.Client) Repository {
	return &entRepository{ent: e}
}

func (e entRepository) FindInstitutionByAdminDomain(ctx context.Context, domain string) (*Institution, error) {
	inst, err := e.ent.Institution.Query().
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

func (e entRepository) FindGroupsByUser(ctx context.Context, email string) ([]*Group, error) {
	u, err := e.ent.Group.Query().
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

func (e entRepository) FindGroupTypes() ([]*GroupType, error) {
	sig := group.GroupTypeSpecialInterestGroup
	mod := group.GroupTypeModuleGroup
	cca := group.GroupTypeCca
	mc := group.GroupTypeMentorClass
	return []*GroupType{
		&sig, &mod, &cca, &mc,
	}, nil
}

func (e entRepository) CreateGroup(ctx context.Context, groupType GroupType, opts ...CreateGroupOption) (*Group, error) {
	var options createGroupOptions
	for _, opt := range opts {
		opt(&options)
	}

	g, err := e.ent.Group.
		Create().
		SetGroupType(groupType).
		SetName(options.name).
		SetPath(options.path).
		SetDescription(options.description).
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
	_, err := e.ent.Group.Delete().Where(group.Path(path)).Exec(ctx)
	return err
}