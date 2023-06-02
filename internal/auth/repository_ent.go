package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/np-inprove/server/internal/apperror"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/jwtrevocation"
	entuser "github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/user"
	"github.com/np-inprove/server/internal/hash"
)

type entRepository struct {
	client *ent.Client
}

func NewEntRepository(ent *ent.Client) Repository {
	return entRepository{ent}
}

func (e entRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u, err := e.client.User.Query().Where(entuser.Email(email)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find user with email: %w", err)
	}

	return u, nil
}

func (e entRepository) FindJWTRevocation(ctx context.Context, jti string) (*entity.JWTRevocation, error) {
	r, err := e.client.JWTRevocation.Query().Where(jwtrevocation.Jti(jti)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find jwt revocation with jti: %w", err)
	}

	return r, nil
}

func (e entRepository) FindInstitutionByDomains(ctx context.Context, domain string) (*entity.Institution, error) {
	i, err := e.client.Institution.Query().Where(institution.Or(institution.StudentDomain(domain), institution.AdminDomain(domain))).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find institution by domains: %w", err)
	}

	return i, nil
}

func (e entRepository) CreateUser(ctx context.Context, firstName string, lastName string, email string, password hash.Encoded, opts ...user.Option) (*entity.User, error) {
	var options user.Options
	for _, opt := range opts {
		opt(&options)
	}

	u, err := e.client.User.Create().
		SetFirstName(firstName).
		SetLastName(lastName).
		SetEmail(email).
		SetPassword(password).
		SetPoints(options.Points).
		SetPointsAwardedCount(options.PointsAwardedCount).
		SetPointsAwardedResetTime(options.PointsAwardedResetTime).
		SetGodMode(options.GodMode).
		Save(ctx)
	if err != nil {
		if apperror.IsConflict(err) {
			return nil, ErrUserConflict
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return u, nil
}

func (e entRepository) CreateJWTRevocation(ctx context.Context, jti string, expiry time.Time) (*entity.JWTRevocation, error) {
	j, err := e.client.JWTRevocation.Create().SetJti(jti).SetExpiry(expiry).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create jwt revocation: %w", err)
	}

	return j, nil
}
