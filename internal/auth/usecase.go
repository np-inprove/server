package auth

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/entity"
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

type useCase struct {
	repo Repository
	cfg  *config.Config

	publicKey  jwk.Key
	privateKey jwk.Key
}

type UseCase interface {
	WhoAmI(ctx context.Context, token jwt.Token) (*entity.User, error)
	Login(ctx context.Context, email, password string) (*entity.Session, error)
	Register(ctx context.Context, inviteCode, firstName, lastName, email, password string) (*entity.Session, error)

	// CheckInviteCode retrieves an invite code and the associated institution from the repository
	CheckInviteCode(ctx context.Context, inviteCode string) (*entity.InstitutionInviteLink, error)

	// GetUserInstitution retrieves the institution associated with the user
	GetUserInstitution(ctx context.Context, token jwt.Token) (*entity.Institution, error)
}

func NewUseCase(r Repository, c *config.Config, publicKey jwk.Key, privateKey jwk.Key) UseCase {
	return useCase{r, c, publicKey, privateKey}
}

func (u useCase) WhoAmI(ctx context.Context, token jwt.Token) (*entity.User, error) {
	if err := u.tokenIsValid(ctx, token.JwtID()); err != nil {
		return nil, fmt.Errorf("failed to find jwt revocation: %w", err)
	}

	user, err := u.repo.FindUserByEmail(ctx, token.Subject())
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return user, nil
}

func (u useCase) Login(ctx context.Context, email string, password string) (*entity.Session, error) {
	user, err := u.repo.FindUserByEmail(ctx, email)
	if err != nil {
		if apperror.IsNotFound(err) {
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

	return &entity.Session{
		User:  user,
		Token: string(j),
	}, nil
}

func (u useCase) Register(ctx context.Context, inviteCode string, firstName string, lastName string, email string, password string) (*entity.Session, error) {
	invite, err := u.repo.FindInstitutionInviteLinkWithInstitution(ctx, inviteCode)
	if err != nil {
		if apperror.IsNotFound(err) {
			return nil, ErrInvalidInvite
		}
		return nil, err
	}

	if invite.Edges.Institution == nil {
		return nil, fmt.Errorf("invite edges not loaded")
	}

	h, err := hash.CreateEncoded(password)
	if err != nil {
		return nil, fmt.Errorf("failed to create encoded password: %w", err)
	}

	inst := invite.Edges.Institution
	user, err := u.repo.CreateUser(ctx, inst.ID, invite.Role, firstName, lastName, email, h)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	j, err := u.createJWT(user.Email, user.GodMode)
	if err != nil {
		return nil, fmt.Errorf("failed to create jwt: %w", err)
	}

	return &entity.Session{
		User:  user,
		Token: string(j),
	}, nil
}

func (u useCase) CheckInviteCode(ctx context.Context, inviteCode string) (*entity.InstitutionInviteLink, error) {
	link, err := u.repo.FindInstitutionInviteLinkWithInstitution(ctx, inviteCode)
	if err != nil {
		if apperror.IsNotFound(err) {
			return nil, ErrInvalidInvite
		}
		return nil, fmt.Errorf("failed to get invite link: %w", err)
	}

	return link, nil
}

func (u useCase) GetUserInstitution(ctx context.Context, token jwt.Token) (*entity.Institution, error) {
	if err := u.tokenIsValid(ctx, token.JwtID()); err != nil {
		return nil, fmt.Errorf("failed to find jwt revocation: %w", err)
	}

	user, err := u.repo.FindUserByEmailWithInstitute(ctx, token.Subject())
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return user.Edges.Institution, nil
}

func (u useCase) createJWT(email string, godMode bool) ([]byte, error) {
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

func (u useCase) tokenIsValid(ctx context.Context, jti string) error {
	_, err := u.repo.FindJWTRevocation(ctx, jti)
	if err == nil {
		return ErrTokenRevoked
	}

	if apperror.IsNotFound(err) {
		return nil
	}

	return err
}
