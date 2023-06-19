package auth

import (
	"context"
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/fixture"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/np-inprove/server/internal/auth/mocks"
	"github.com/np-inprove/server/internal/config"
	"github.com/stretchr/testify/suite"
)

type UseCaseTestSuite struct {
	suite.Suite
	publicKey  jwk.Key
	privateKey jwk.Key
}

func TestUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UseCaseTestSuite))
}

func (suite *UseCaseTestSuite) SetupSuite() {
	suite.privateKey, _ = jwk.ParseKey([]byte(fixture.AuthJWKString))
	suite.publicKey, _ = suite.privateKey.PublicKey()
}

func (suite *UseCaseTestSuite) TestNewUseCase() {
	type args struct {
		repository func() Repository
		cfg        func() *config.Config
		publicKey  func() jwk.Key
		privateKey func() jwk.Key
	}

	tests := []struct {
		name string
		args args
		want useCase
	}{
		{
			name: "New use case",
			args: args{
				repository: func() Repository {
					return new(mocks.MockRepository)
				},
				cfg: func() *config.Config {
					return new(config.Config)
				},
				publicKey: func() jwk.Key {
					return suite.publicKey
				},
				privateKey: func() jwk.Key {
					return suite.privateKey
				},
			},
			want: useCase{repo: new(mocks.MockRepository), cfg: new(config.Config), publicKey: suite.publicKey, privateKey: suite.privateKey},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			result := NewUseCase(tc.args.repository(), tc.args.cfg(), tc.args.publicKey(), tc.args.privateKey())
			suite.Equal(tc.want, result)
		})
	}
}

func (suite *UseCaseTestSuite) TestWhoAmI() {
	type args struct {
		ctx   context.Context
		token jwt.Token
	}

	fancyToken, _ := jwt.NewBuilder().Build()

	tests := []struct {
		name      string
		args      args
		configure func(repository *mocks.MockRepository)
		want      *entity.User
		err       error
	}{
		{
			name: "Token revocation repository error",
			args: args{
				ctx:   context.Background(),
				token: fancyToken,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindJWTRevocation", mock.Anything, mock.Anything).
					Return(&entity.JWTRevocation{}, fixture.RepoInternalErr)
			},
			want: nil,
			err:  fmt.Errorf("failed to find jwt revocation: %w", fixture.RepoInternalErr),
		},
		{
			name: "Token revoked",
			args: args{
				ctx:   context.Background(),
				token: fancyToken,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindJWTRevocation", mock.Anything, mock.Anything).
					Return(&entity.JWTRevocation{}, nil)
			},
			want: nil,
			err:  fmt.Errorf("failed to find jwt revocation: %w", ErrTokenRevoked),
		},
		{
			name: "User does not exist",
			args: args{
				ctx:   context.Background(),
				token: fancyToken,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindJWTRevocation", mock.Anything, mock.Anything).
					Return(&entity.JWTRevocation{}, fixture.RepoNotFoundErr)
				repository.On("FindUserByEmail", mock.Anything, mock.Anything).
					Return(nil, fixture.RepoNotFoundErr)
			},
			want: nil,
			err:  fmt.Errorf("failed to find user: %w", fixture.RepoNotFoundErr),
		},
		{
			name: "Success",
			args: args{
				ctx:   context.Background(),
				token: fancyToken,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindJWTRevocation", mock.Anything, mock.Anything).
					Return(&entity.JWTRevocation{}, fixture.RepoNotFoundErr)
				repository.On("FindUserByEmail", mock.Anything, mock.Anything).
					Return(&entity.User{}, nil)
			},
			want: &entity.User{},
			err:  nil,
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			repo := new(mocks.MockRepository)
			uc := NewUseCase(repo, nil, suite.publicKey, suite.privateKey)
			tc.configure(repo)
			ret, err := uc.WhoAmI(tc.args.ctx, tc.args.token)
			suite.Equal(err, tc.err)
			suite.Equal(tc.want, ret)
		})
	}
}

func (suite *UseCaseTestSuite) TestLogin() {
	type args struct {
		ctx      context.Context
		email    string
		password string
	}

	tests := []struct {
		name      string
		args      args
		configure func(repository *mocks.MockRepository)
		assert    func(ret *entity.Session)
		err       error
	}{
		{
			name: "User does not exist",
			args: args{
				ctx:      context.Background(),
				email:    "example@example.com",
				password: "example",
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindUserByEmail", mock.Anything, mock.Anything).
					Return(nil, fixture.RepoNotFoundErr)
			},
			assert: func(ret *entity.Session) {
				suite.Nil(ret)
			},
			err: ErrUserNotFound,
		},
		{
			name: "Find user internal repo error",
			args: args{
				ctx:      context.Background(),
				email:    fixture.UserJohn.Email,
				password: fixture.UserJohnPassword,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindUserByEmail", mock.Anything, mock.Anything).
					Return(nil, fixture.RepoInternalErr)
			},
			assert: func(ret *entity.Session) {
				suite.Nil(ret)
			},
			err: fmt.Errorf("failed to find user: %w", fixture.RepoInternalErr),
		},
		{
			name: "Hash verification failed",
			args: args{
				ctx:      context.Background(),
				email:    fixture.UserJohn.Email,
				password: "wrong password",
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindUserByEmail", mock.Anything, mock.Anything).
					Return(fixture.UserJohn, nil)
			},
			assert: func(ret *entity.Session) {
				suite.Nil(ret)
			},
			err: ErrInvalidPassword,
		},
		{
			name: "Success",
			args: args{
				ctx:      context.Background(),
				email:    fixture.UserJohn.Email,
				password: fixture.UserJohnPassword,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindUserByEmail", mock.Anything, mock.Anything).
					Return(fixture.UserJohn, nil)
			},
			assert: func(ret *entity.Session) {
				suite.Equal(ret.User, fixture.UserJohn)
				suite.NotEmpty(ret.Token)
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			repo := new(mocks.MockRepository)
			uc := NewUseCase(repo, nil, suite.publicKey, suite.privateKey)
			tc.configure(repo)
			ret, err := uc.Login(tc.args.ctx, tc.args.email, tc.args.password)
			suite.Equal(tc.err, err)
			tc.assert(ret)
		})
	}
}

func (suite *UseCaseTestSuite) TestRegister() {
	type args struct {
		ctx        context.Context
		inviteCode string
		firstName  string
		lastName   string
		email      string
		password   string
	}

	tests := []struct {
		name      string
		args      args
		configure func(repository *mocks.MockRepository)
		assert    func(repository *mocks.MockRepository)
		err       error
	}{
		{
			name: "Invalid invite",
			args: args{
				ctx:        context.Background(),
				inviteCode: "definitely not a valid invite code",
				firstName:  fixture.UserJohn.FirstName,
				lastName:   fixture.UserJohn.LastName,
				email:      fixture.UserJohn.Email,
				password:   fixture.UserJohnPassword,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindInstitutionInviteLinkWithInstitution", mock.Anything, mock.Anything).
					Return(nil, fixture.RepoNotFoundErr)
			},
			assert: func(repository *mocks.MockRepository) {
			},
			err: ErrInvalidInvite,
		},
		{
			name: "Repo internal error",
			args: args{
				ctx:        context.Background(),
				inviteCode: fixture.InstitutionInviteLinkNP.Code,
				firstName:  fixture.UserJohn.FirstName,
				lastName:   fixture.UserJohn.LastName,
				email:      fixture.UserJohn.Email,
				password:   fixture.UserJohnPassword,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindInstitutionInviteLinkWithInstitution", mock.Anything, mock.Anything).
					Return(nil, fixture.RepoInternalErr)
			},
			assert: func(repository *mocks.MockRepository) {
			},
			err: fixture.RepoInternalErr,
		},
		{
			name: "Success",
			args: args{
				ctx:        context.Background(),
				inviteCode: fixture.InstitutionInviteLinkNP.Code,
				firstName:  fixture.UserJohn.FirstName,
				lastName:   fixture.UserJohn.LastName,
				email:      fixture.UserJohn.Email,
				password:   fixture.UserJohnPassword,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindInstitutionInviteLinkWithInstitution", mock.Anything, mock.Anything).
					Return(fixture.InstitutionInviteLinkNP, nil)
				repository.On("CreateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(fixture.UserJohn, nil)
			},
			assert: func(repository *mocks.MockRepository) {
				repository.AssertCalled(suite.T(), "CreateUser",
					context.Background(),
					fixture.InstitutionInviteLinkNP.Edges.Institution.ID,
					fixture.InstitutionInviteLinkNP.Role,
					fixture.UserJohn.FirstName,
					fixture.UserJohn.LastName,
					fixture.UserJohn.Email,
					mock.Anything,
				)
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			repo := new(mocks.MockRepository)
			uc := NewUseCase(repo, nil, suite.publicKey, suite.privateKey)
			tc.configure(repo)
			_, err := uc.Register(tc.args.ctx, tc.args.inviteCode, tc.args.firstName, tc.args.lastName, tc.args.email, tc.args.password)
			suite.Equal(err, tc.err)
			tc.assert(repo)
		})
	}
}
