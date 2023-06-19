package institution

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/fixture"
	"github.com/np-inprove/server/internal/institution/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
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
					Return(nil, fixture.RepoInternalErr)
			},
			want: nil,
			err:  fmt.Errorf("failed to find institutions: %w", fixture.RepoInternalErr),
		},
		{
			name: "Success",
			args: args{ctx: context.Background()},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindInstitutions", mock.Anything).
					Return(fixture.Institutions, nil)
			},
			want: fixture.Institutions,
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
		want      *entity.Institution
	}{
		{
			name: "Wraps repo in tx",
			args: args{
				ctx:         context.Background(),
				name:        fixture.InstitutionNP.Name,
				shortName:   fixture.InstitutionNP.ShortName,
				description: fixture.InstitutionNP.Description,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("WithTx", mock.Anything, mock.Anything).Return(nil, fixture.RepoInternalErr)
			},
			assert: func(ret *entity.Institution, repository *mocks.MockRepository) {
				repository.AssertCalled(suite.T(), "WithTx", mock.Anything, mock.Anything)
			},
			err:  fixture.RepoInternalErr,
			want: nil,
		},
		{
			name: "Tx internal error",
			args: args{
				ctx:         context.Background(),
				name:        fixture.InstitutionNP.Name,
				shortName:   fixture.InstitutionNP.ShortName,
				description: fixture.InstitutionNP.Description,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("CreateInstitution", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, fixture.RepoInternalErr)
				repository.On("WithTx", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					fn := args.Get(1).(func(ctx context.Context) (interface{}, error))
					_, _ = fn(context.Background())
				}).Return(nil, fixture.RepoInternalErr)
			},
			assert: func(ret *entity.Institution, repository *mocks.MockRepository) {
				repository.AssertCalled(suite.T(), "CreateInstitution", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			},
			err:  fixture.RepoInternalErr,
			want: nil,
		},
		{
			name: "Short name conflict",
			args: args{
				ctx:         context.Background(),
				name:        fixture.InstitutionNP.Name,
				shortName:   fixture.InstitutionNP.ShortName,
				description: fixture.InstitutionNP.Description,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("CreateInstitution", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, fixture.RepoConflictErr)
				repository.On("WithTx", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					fn := args.Get(1).(func(ctx context.Context) (interface{}, error))
					_, _ = fn(context.Background())
				}).Return(nil, ErrInstitutionShortNameConflict)
			},
			assert: func(ret *entity.Institution, repository *mocks.MockRepository) {
				repository.AssertCalled(suite.T(), "CreateInstitution", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			},
			err:  ErrInstitutionShortNameConflict,
			want: nil,
		},
		{
			name: "Success",
			args: args{
				ctx:         context.Background(),
				name:        fixture.InstitutionNP.Name,
				shortName:   fixture.InstitutionNP.ShortName,
				description: fixture.InstitutionNP.Description,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("CreateInstitution", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(fixture.InstitutionNP, nil)
				repository.On("WithTx", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					fn := args.Get(1).(func(ctx context.Context) (interface{}, error))
					_, _ = fn(context.Background())
				}).Return(fixture.InstitutionNP, nil)
			},
			assert: func(ret *entity.Institution, repository *mocks.MockRepository) {
				repository.AssertCalled(suite.T(), "CreateInstitution", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			},
			err:  nil,
			want: fixture.InstitutionNP,
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
			suite.Equal(tc.want, ret)
		})
	}
}

func (suite *UseCaseTestSuite) TestDeleteInstitution() {
	type args struct {
		ctx       context.Context
		shortName string
	}

	tests := []struct {
		name      string
		args      args
		configure func(repository *mocks.MockRepository)
		err       error
	}{
		{
			name: "Institution does not exist",
			args: args{
				ctx:       context.Background(),
				shortName: "does not exist",
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindInstitution", mock.Anything, mock.Anything).
					Return(nil, fixture.RepoNotFoundErr)
			},
			err: fmt.Errorf("failed to find institution: %w", fixture.RepoNotFoundErr),
		},
		{
			name: "Internal repo error",
			args: args{
				ctx:       context.Background(),
				shortName: fixture.InstitutionNP.ShortName,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindInstitution", mock.Anything, mock.Anything).
					Return(fixture.InstitutionNP, nil)
				repository.On("DeleteInstitution", mock.Anything, mock.Anything).
					Return(fixture.RepoInternalErr)
			},
			err: fmt.Errorf("failed to delete institution: %w", fixture.RepoInternalErr),
		},
		{
			name: "Success",
			args: args{
				ctx:       context.Background(),
				shortName: fixture.InstitutionNP.ShortName,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindInstitution", mock.Anything, mock.Anything).
					Return(fixture.InstitutionNP, nil)
				repository.On("DeleteInstitution", mock.Anything, mock.Anything).
					Return(nil)
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			repo := new(mocks.MockRepository)
			tc.configure(repo)
			u := NewUseCase(repo)
			err := u.DeleteInstitution(tc.args.ctx, tc.args.shortName)
			suite.Equal(tc.err, err)
		})
	}
}
