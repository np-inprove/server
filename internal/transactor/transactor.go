package transactor

import "context"

type Transactor interface {
	WithTx(context.Context, func(ctx context.Context) error) error
}
