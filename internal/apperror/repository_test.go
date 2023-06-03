package apperror

import (
	"github.com/np-inprove/server/internal/ent"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RepositoryTestSuite struct {
	suite.Suite
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestIsNotFound() {
	err := &ent.NotFoundError{}
	suite.True(IsNotFound(err))
}

func (suite *RepositoryTestSuite) TestIsConflict() {
	err := &ent.ConstraintError{}
	suite.True(IsConflict(err))
}
