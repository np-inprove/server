package apperror

import "github.com/np-inprove/server/internal/ent"

func IsNotFound(err error) bool {
	return ent.IsNotFound(err)
}

func IsConflict(err error) bool {
	return ent.IsConstraintError(err)
}
