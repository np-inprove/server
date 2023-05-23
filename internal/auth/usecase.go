package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/hash"
)

// TODO move to config
var (
	jwtIssuer   = "np-inprove-server"
	jwtAudience = []string{"np-inprove.qinguan.me"}
)

// session is used to represent a successful call to usecase.Login.
// It contains the user information retrieved, as well as the token issued.
// The handlers should distribute the issued token to the client.
type session struct {
	user  *User
	token string
}

type usecase struct {
	r Repository
	c *config.Config

	publicKey  jwk.Key
	privateKey jwk.Key
}

type UseCase interface {
	WhoAmI(ctx context.Context, token jwt.Token) (*User, error)
	Login(ctx context.Context, email string, password string) (*session, error)
	Register(ctx context.Context, firstName string, lastName string, email string, password string) (*session, error)
}

func NewUseCase(r Repository, c *config.Config, publicKey jwk.Key, privateKey jwk.Key) (UseCase, error) {
	return usecase{r, c, publicKey, privateKey}, nil
}

func (u usecase) WhoAmI(ctx context.Context, token jwt.Token) (*User, error) {
	if err := u.tokenIsValid(ctx, token.JwtID()); err != nil {
		return nil, fmt.Errorf("failed to find jwt revocation: %w", err)
	}

	user, err := u.r.FindUserByEmail(ctx, token.Subject())
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return user, nil
}

func (u usecase) Login(ctx context.Context, email string, password string) (*session, error) {
	user, err := u.r.FindUserByEmail(ctx, email)
	if err != nil {
		if apperror.IsEntityNotFound(err) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	ok, err := hash.VerifyEncoded(user.Password, password)
	if err != nil {
		return nil, fmt.Errorf("failed to verify encoded password: %w", err)
	}

	if !ok {
		return nil, ErrInvalidPassword
	}

	j, err := u.createJWT(user.Email, user.GodMode)
	if err != nil {
		return nil, fmt.Errorf("failed to create jwt: %w", err)
	}

	return &session{
		user:  user,
		token: string(j),
	}, nil
}

func (u usecase) Register(ctx context.Context, firstName string, lastName string, email string, password string) (*session, error) {
	domain := strings.Split(email, "@")[1] // This should not panic
	if _, err := u.r.FindInstitutionByDomains(ctx, domain); err != nil {
		if apperror.IsEntityNotFound(err) {
			return nil, ErrDomainNotFound
		}
		return nil, err
	}

	h, err := hash.CreateEncoded(password)
	if err != nil {
		return nil, fmt.Errorf("failed to create encoded password: %w", err)
	}

	user, err := u.r.CreateUser(ctx, firstName, lastName, email, h)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	j, err := u.createJWT(user.Email, user.GodMode)
	if err != nil {
		return nil, fmt.Errorf("failed to create jwt: %w", err)
	}

	return &session{
		user:  user,
		token: string(j),
	}, nil
}

func (u usecase) createJWT(email string, godMode bool) ([]byte, error) {
	jti, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate jti: %w", err)
	}

	t, err := jwt.NewBuilder().
		Subject(email).
		Audience(jwtAudience).
		Issuer(jwtIssuer).
		JwtID(jti.String()).
		IssuedAt(time.Now()).
		NotBefore(time.Now()).
		Expiration(time.Now().Add(30*time.Minute)).
		// The Claim for god mode which is checked when calling god mode endpoints.
		// Ensure it is the same as used in middleware/middleware.go
		Claim("god", godMode).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build jwt: %w", err)
	}

	j, err := jwt.Sign(t, jwt.WithKey(u.privateKey.Algorithm(), u.privateKey))
	if err != nil {
		return nil, fmt.Errorf("failed to sign jwt: %w", err)
	}

	return j, nil
}

func (u usecase) tokenIsValid(ctx context.Context, jti string) error {
	_, err := u.r.FindJWTRevocation(ctx, jti)
	if err != nil && apperror.IsEntityNotFound(err) {
		return nil
	}
	return err
}
