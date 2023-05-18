package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/config"
)

type usecase struct {
	r Repository
	c *config.Config

	publicKey  jwk.Key
	privateKey jwk.Key
}

type UseCase interface {
	WhoAmI(ctx context.Context, token jwt.Token) (User, error)
}

func NewUseCase(r Repository, c *config.Config, publicKey jwk.Key, privateKey jwk.Key) (UseCase, error) {
	return usecase{r, c, publicKey, privateKey}, nil
}

func (u usecase) WhoAmI(ctx context.Context, token jwt.Token) (User, error) {
	if err := u.tokenIsValid(ctx, token.JwtID()); err != nil {
		return User{}, fmt.Errorf("failed to find jwt revocation: %v", err)
	}

	user, err := u.r.FindUser(ctx, Email(token.Subject()))
	if err != nil {
		return User{}, fmt.Errorf("failed to find user: %v", err)
	}

	return user, nil
}

func (u usecase) tokenIsValid(ctx context.Context, jti string) error {
	_, err := u.r.FindJWTRevocation(ctx, jti)
	if err != nil {
		var e apperror.ErrEntityNotFound
		if errors.Is(err, &e) {
			return nil
		}
	}
	return err
}
