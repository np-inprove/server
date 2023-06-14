package institution

import "errors"

type Role string

func (Role) Values() (kinds []string) {
	for _, s := range Roles {
		kinds = append(kinds, string(s))
	}
	return
}

func (ir Role) Validate() error {
	switch ir {
	case RoleAdmin, RoleEducator, RoleMember:
		return nil
	default:
		return errors.New("institution role is not valid")
	}
}

var (
	Roles = []Role{
		RoleAdmin, RoleEducator, RoleMember,
	}

	RoleAdmin    Role = "admin"
	RoleEducator Role = "educator"
	RoleMember   Role = "member"
)
