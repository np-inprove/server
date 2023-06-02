package group

import (
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/entity"
)

func TypeValidator(gt entity.GroupType) error {
	return group.GroupTypeValidator(gt)
}

type Options struct {
	Name        string
	Path        string
	Description string
}

type Option func(*Options)

func Name(n string) Option {
	return func(opts *Options) {
		opts.Name = n
	}
}

func Path(p string) Option {
	return func(opts *Options) {
		opts.Path = p
	}
}

func Description(d string) Option {
	return func(opts *Options) {
		opts.Description = d
	}
}
