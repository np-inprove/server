package apperror

import "github.com/np-inprove/server/internal/ent"

func IsEntityNotFound(err error) bool {
	return ent.IsNotFound(err)
}
