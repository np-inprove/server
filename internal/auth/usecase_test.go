package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/hash"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/np-inprove/server/internal/auth/mocks"
	"github.com/np-inprove/server/internal/config"
	"github.com/stretchr/testify/suite"
)

var (
	repoInternalErr = errors.New("repo err")
	repoNotFoundErr = new(ent.NotFoundError)

	password              = "example"
	encodedPasswordString = "{\"h\":\"15BrHhXrHhYM2yMFMrghE9gnY950dHNF1klveR5EdeY=\",\"s\":\"J7lr7qPnc/S4iGgdJQWSK2xVELHLE9N4+KIdcdBdHhM=\",\"t\":1,\"m\":65536,\"k\":32}"
	jwkString             = "{\"kty\": \"EC\",\"d\": \"AFDDClQjKNoNfFGZB9iLP2gRmAJXBa-mp_XIdYMJjNKLIRELS03WArued737uRt9ayopD6AEQZOiRSLs8o9LrT5B\",\"use\": \"sig\",\"crv\": \"P-521\",\"x\": \"AZGgvUEwezY2RHU7l4m1aF5vp8Vj8CUo_nyqUzVVwoVjFz_mbZM4ZznXrAvuLKlAqagI2bHHj-97W7Blp3VaI0A5\",\"y\": \"AdmmlsSae6qoGTPsUYj7gSOGMJcaRdHpcyvQ3hdq5owJnyozbRib5mUYXwsV_voIKBTSkeztQSRagiILwqiEy5-y\",\"alg\": \"ES512\"}"
)

type UseCaseTestSuite struct {
	suite.Suite
	publicKey       jwk.Key
	privateKey      jwk.Key
	encodedPassword hash.Encoded
}

func TestUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UseCaseTestSuite))
}

func (suite *UseCaseTestSuite) SetupSuite() {
	suite.privateKey, _ = jwk.ParseKey([]byte(jwkString))
	suite.publicKey, _ = suite.privateKey.PublicKey()
	_ = json.Unmarshal([]byte(encodedPasswordString), &suite.encodedPassword)
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
			name: "New folder",
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
					Return(&entity.JWTRevocation{}, repoInternalErr)
			},
			want: nil,
			err:  fmt.Errorf("failed to find jwt revocation: %w", repoInternalErr),
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
					Return(&entity.JWTRevocation{}, repoNotFoundErr)
				repository.On("FindUserByEmail", mock.Anything, mock.Anything).
					Return(nil, repoNotFoundErr)
			},
			want: nil,
			err:  fmt.Errorf("failed to find user: %w", repoNotFoundErr),
		},
		{
			name: "Success",
			args: args{
				ctx:   context.Background(),
				token: fancyToken,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindJWTRevocation", mock.Anything, mock.Anything).
					Return(&entity.JWTRevocation{}, repoNotFoundErr)
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
					Return(nil, repoNotFoundErr)
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
				email:    "example@example.com",
				password: "example",
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindUserByEmail", mock.Anything, mock.Anything).
					Return(nil, repoInternalErr)
			},
			assert: func(ret *entity.Session) {
				suite.Nil(ret)
			},
			err: fmt.Errorf("failed to find user: %w", repoInternalErr),
		},
		{
			name: "Hash verification failed",
			args: args{
				ctx:      context.Background(),
				email:    "example@example.com",
				password: "wrong password",
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindUserByEmail", mock.Anything, mock.Anything).
					Return(&entity.User{Password: suite.encodedPassword}, nil)
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
				email:    "example@example.com",
				password: password,
			},
			configure: func(repository *mocks.MockRepository) {
				repository.On("FindUserByEmail", mock.Anything, mock.Anything).
					Return(&entity.User{
						Email:    "example@example.com",
						Password: suite.encodedPassword,
						GodMode:  false,
					}, nil)
			},
			assert: func(ret *entity.Session) {
				suite.Equal(ret.User, &entity.User{
					Email:    "example@example.com",
					Password: suite.encodedPassword,
					GodMode:  false,
				})
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
			suite.Equal(err, tc.err)
			tc.assert(ret)
		})
	}
}
