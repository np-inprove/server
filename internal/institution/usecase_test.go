package institution

import (
	"context"
	"errors"
	"fmt"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/institution/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	repoInternalErr = errors.New("funny repo error")
)

type UseCaseTestSuite struct {
	suite.Suite
}

func TestUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UseCaseTestSuite))
}

func (suite *UseCaseTestSuite) TestNewUseCase() {
	type args struct {
		repo Repository
	}

	tests := []struct {
		name string
		args args
		want *useCase
	}{
		{
			name: "Success",
			args: args{repo: new(mocks.MockRepository)},
			want: &useCase{repo: new(mocks.MockRepository)},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			u := NewUseCase(tc.args.repo)
			suite.Equal(tc.want, u)
		})
	}
}

func (suite *UseCaseTestSuite) TestListInstitutions() {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name      string
		args      args
		configure func(repository *mocks.MockRepository)
		want      []*entity.Institution
		err       error
	}{
		{
			name: "Internal repo error",
			args: args{ctx: context.Background()},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindInstitutions", mock.Anything).
					Return(nil, repoInternalErr)
			},
			want: nil,
			err:  fmt.Errorf("failed to find institutions: %w", repoInternalErr),
		},
		{
			name: "Success",
			args: args{ctx: context.Background()},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindInstitutions", mock.Anything).
					Return([]*entity.Institution{}, nil)
			},
			want: []*entity.Institution{},
			err:  nil,
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			repo := new(mocks.MockRepository)
			tc.configure(repo)
			u := NewUseCase(repo)
			ret, err := u.ListInstitutions(tc.args.ctx)
			suite.Equal(tc.err, err)
			suite.Equal(tc.want, ret)
		})
	}
}

func (suite *UseCaseTestSuite) TestCreateInstitution() {
	type args struct {
		ctx         context.Context
		name        string
		shortName   string
		description string
	}

	tests := []struct {
		name      string
		args      args
		configure func(repository *mocks.MockRepository)
		assert    func(ret *entity.Institution, repository *mocks.MockRepository)
		err       error
	}{
		{
			name: "Wraps repo in tx",
			args: args{
				ctx:         context.Background(),
				name:        "Funny institution",
				shortName:   "funny-institution",
				description: "",
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("WithTx", mock.Anything, mock.Anything).Return(repoInternalErr)
			},
			assert: func(ret *entity.Institution, repository *mocks.MockRepository) {
				repository.AssertCalled(suite.T(), "WithTx", mock.Anything, mock.Anything)
			},
			err: repoInternalErr,
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			repo := new(mocks.MockRepository)
			tc.configure(repo)
			u := NewUseCase(repo)
			ret, err := u.CreateInstitution(tc.args.ctx, tc.args.name, tc.args.shortName, tc.args.description)
			suite.Equal(tc.err, err)
			tc.assert(ret, repo)
		})
	}
}
