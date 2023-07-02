package forum

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
