package security

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/ent"
	entgroup "github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/entity/group"
	"github.com/np-inprove/server/internal/entity/institution"
)

type RBAC interface {
	UserInInstitution(ctx context.Context, principal string, id int) (*institution.Role, error)
	UserInGroup(ctx context.Context, principal string, id int) (*group.Role, error)
	UserHasGroupRole(role group.Role, allowedRoles ...group.Role) bool
	UserHasInstitutionRole(role institution.Role, allowedRoles ...institution.Role) bool
}

type rbac struct {
	client *ent.Client
}

func NewRBAC(client *ent.Client) RBAC {
	return &rbac{client}
}

func (s rbac) UserInInstitution(ctx context.Context, principal string, id int) (*institution.Role, error) {
	usr, err := s.client.User.Query().Where(user.Email(principal)).WithInstitution().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query user: %w", err)
	}

	if usr.Edges.Institution.ID != id {
		return nil, fmt.Errorf("insufficient permissions: %w", err)
	}

	return &usr.Role, nil
}

func (s rbac) UserInGroup(ctx context.Context, principal string, id int) (*group.Role, error) {
	grpusr, err := s.client.User.Query().Where(user.Email(principal), user.HasGroupsWith(entgroup.ID(id))).QueryGroupUsers().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query user: %w", err)
	}

	return &grpusr.Role, nil
}

func (s rbac) UserHasGroupRole(role group.Role, allowedRoles ...group.Role) bool {
	res := false
	for _, allowed := range allowedRoles {
		if role == allowed {
			res = true
			break
		}
	}
	return res
}

func (s rbac) UserHasInstitutionRole(role institution.Role, allowedRoles ...institution.Role) bool {
	res := false
	for _, allowed := range allowedRoles {
		if role == allowed {
			res = true
			break
		}
	}
	return res
}
