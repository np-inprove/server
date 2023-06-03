package auth

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/hash"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/np-inprove/server/internal/auth/mocks"
	"github.com/np-inprove/server/internal/config"
	"github.com/stretchr/testify/suite"
)

var (
	publicKey       jwk.Key
	privateKey      jwk.Key
	repoInternalErr = errors.New("repo err")
	repoNotFoundErr = new(ent.NotFoundError)

	password              = "example"
	encodedPasswordString = "{\"h\":\"15BrHhXrHhYM2yMFMrghE9gnY950dHNF1klveR5EdeY=\",\"s\":\"J7lr7qPnc/S4iGgdJQWSK2xVELHLE9N4+KIdcdBdHhM=\",\"t\":1,\"m\":65536,\"k\":32}\n"
	encodedPassword       hash.Encoded
)

type UseCaseTestSuite struct {
	suite.Suite
}

func TestUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UseCaseTestSuite))
}

func (suite *UseCaseTestSuite) SetupSuite() {
	// These should never fail
	raw, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	privateKey, _ = jwk.FromRaw(raw)
	publicKey, _ = privateKey.PublicKey()
	_ = json.Unmarshal([]byte(encodedPasswordString), &encodedPassword)
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
		want usecase
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
					return publicKey
				},
				privateKey: func() jwk.Key {
					return privateKey
				},
			},
			want: usecase{repo: new(mocks.MockRepository), cfg: new(config.Config), publicKey: publicKey, privateKey: privateKey},
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
			uc := NewUseCase(repo, nil, publicKey, privateKey)
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
		want      *entity.Session
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
			want: nil,
			err:  ErrUserNotFound,
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
			want: nil,
			err:  fmt.Errorf("failed to find user: %w", repoInternalErr),
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
					Return(&entity.User{Password: encodedPassword}, nil)
			},
			want: nil,
			err:  ErrInvalidPassword,
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			repo := new(mocks.MockRepository)
			uc := NewUseCase(repo, nil, publicKey, privateKey)
			tc.configure(repo)
			ret, err := uc.Login(tc.args.ctx, tc.args.email, tc.args.password)
			suite.Equal(err, tc.err)
			suite.Equal(tc.want, ret)
		})
	}
}
