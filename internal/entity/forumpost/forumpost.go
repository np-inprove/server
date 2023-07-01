package forumpost

import "errors"

type Role string

func (Role) Values() (kinds []string) {
	for _, s := range Roles {
		kinds = append(kinds, string(s))
	}
	return
}

func (fr Role) Validate() error {
	switch fr {
	case RoleOwner, RoleEducator, RoleMember:
		return nil
	default:
		return errors.New("group role is not valid")
	}
}

var (
	Roles = []Role{RoleOwner, RoleEducator, RoleMember}

	RoleOwner    Role = "owner"
	RoleEducator Role = "educator"
	RoleMember   Role = "member"
)

type Options struct {
	Title   string
	Content string
}

type Option func(*Options)

func Name(n string) Option {
	return func(opts *Options) {
		opts.Title = n
	}
}

func ShortName(n string) Option {
	return func(opts *Options) {
		opts.Content = n
	}
}
