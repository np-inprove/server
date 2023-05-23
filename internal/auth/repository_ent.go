package auth

import (
	"context"
	"fmt"
	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/jwtrevocation"
	"github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/hash"
	"time"
)

type entRepository struct {
	ent *ent.Client
}

func NewEntRepository(ent *ent.Client) Repository {
	return entRepository{ent}
}

func (e entRepository) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	u, err := e.ent.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find user with email: %w", err)
	}

	return u, nil
}

func (e entRepository) FindJWTRevocation(ctx context.Context, jti string) (*JWTRevocation, error) {
	r, err := e.ent.JWTRevocation.Query().Where(jwtrevocation.Jti(jti)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find jwt revocation with jti: %w", err)
	}

	return r, nil
}

func (e entRepository) FindInstitutionByDomains(ctx context.Context, domain string) (*Institution, error) {
	i, err := e.ent.Institution.Query().Where(institution.Or(institution.StudentDomain(domain), institution.AdminDomain(domain))).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find institution by domains: %w", err)
	}

	return i, nil
}

func (e entRepository) CreateUser(ctx context.Context, firstName string, lastName string, email string, password hash.Encoded, opts ...CreateUserOption) (*User, error) {
	var options createUserOptions
	for _, opt := range opts {
		opt(&options)
	}

	u, err := e.ent.User.Create().
		SetFirstName(firstName).
		SetLastName(lastName).
		SetEmail(email).
		SetPassword(password).
		SetPoints(options.points).
		SetPointsAwardedCount(options.pointsAwardedCount).
		SetPointsAwardedResetTime(options.pointsAwardedResetTime).
		SetGodMode(options.godMode).
		Save(ctx)
	if err != nil {
		if apperror.IsConflict(err) {
			return nil, ErrUserConflict
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return u, nil
}

func (e entRepository) CreateJWTRevocation(ctx context.Context, jti string, expiry time.Time) (*JWTRevocation, error) {
	j, err := e.ent.JWTRevocation.Create().SetJti(jti).SetExpiry(expiry).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create jwt revocation: %w", err)
	}

	return j, nil
}
