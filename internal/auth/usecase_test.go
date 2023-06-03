package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/np-inprove/server/internal/auth/mocks"
	"github.com/np-inprove/server/internal/config"
	"github.com/stretchr/testify/suite"
)

var (
	publicKey  jwk.Key
	privateKey jwk.Key
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
}

func (suite *UseCaseTestSuite) TestNewFolder() {
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
			result, err := NewUseCase(tc.args.repository(), tc.args.cfg(), tc.args.publicKey(), tc.args.privateKey())
			if err != nil {
				panic(err)
			}
			suite.Equal(tc.want, result)
		})
	}
}
