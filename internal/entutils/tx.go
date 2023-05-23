package entutils

import (
	"context"
	"github.com/np-inprove/server/internal/ent"
)

func ExtractTx(ctx context.Context, key string) (*ent.Client, bool) {
	v := ctx.Value(key)

	if v == nil {
		return nil, false
	}

	if v, ok := v.(*ent.Client); ok {
		return v, true
	}

	return nil, false
}
