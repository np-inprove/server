package dash

import (
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/ent/group"
)

type User = ent.User
type Group = ent.Group
type GroupUser = ent.GroupUser
type Institution = ent.Institution
type GroupType = group.GroupType

func GroupTypeValidator(gt GroupType) error {
	return group.GroupTypeValidator(gt)
}

type CreateGroupOptions struct {
	name        string
	path        string
	description string
}

type CreateGroupOption func(*CreateGroupOptions)

func Name(n string) CreateGroupOption {
	return func(opts *CreateGroupOptions) {
		opts.name = n
	}
}

func Path(p string) CreateGroupOption {
	return func(opts *CreateGroupOptions) {
		opts.path = p
	}
}

func Description(d string) CreateGroupOption {
	return func(opts *CreateGroupOptions) {
		opts.description = d
	}
}
