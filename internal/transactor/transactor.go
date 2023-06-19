package transactor

import "context"

type Transactor interface {
	WithTx(
		context.Context,
		func(ctx context.Context) (interface{}, error),
	) (interface{}, error) // TODO possible to switch to generics
}
