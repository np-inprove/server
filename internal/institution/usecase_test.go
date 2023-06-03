package institution

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UseCaseTestSuite struct {
	suite.Suite
}

func TestUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UseCaseTestSuite))
}
