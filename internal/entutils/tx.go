package entutils

import (
	"context"
	"github.com/np-inprove/server/internal/ent"
)

type CtxKey int

const (
	EntTxCtxKey CtxKey = iota
)

func ExtractTx(ctx context.Context) (*ent.Client, bool) {
	v := ctx.Value(EntTxCtxKey)

	if v == nil {
		return nil, false
	}

	if v, ok := v.(*ent.Client); ok {
		return v, true
	}

	return nil, false
}
