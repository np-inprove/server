package fixture

import (
	"errors"
	"github.com/np-inprove/server/internal/ent"
)

var (
	RepoInternalErr = errors.New("repo internal err")
	RepoNotFoundErr = new(ent.NotFoundError)
	RepoConflictErr = new(ent.ConstraintError)
)
