package group

import "errors"

type Role string

func (Role) Values() (kinds []string) {
	for _, s := range Roles {
		kinds = append(kinds, string(s))
	}
	return
}

func (gr Role) Validate() error {
	switch gr {
	case RoleOwner, RoleEducator, RoleMember:
		return nil
	default:
		return errors.New("institution role is not valid")
	}
}

var (
	Roles = []Role{RoleOwner, RoleEducator, RoleMember}

	RoleOwner    Role = "owner"
	RoleEducator Role = "educator"
	RoleMember   Role = "member"
)

type Options struct {
	Name        string
	ShortName   string
	Description string
}

type Option func(*Options)

func Name(n string) Option {
	return func(opts *Options) {
		opts.Name = n
	}
}

func ShortName(n string) Option {
	return func(opts *Options) {
		opts.ShortName = n
	}
}

func Description(d string) Option {
	return func(opts *Options) {
		opts.Description = d
	}
}
