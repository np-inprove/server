package auth

import (
	"context"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/user"
	"github.com/np-inprove/server/internal/hash"
	"time"
)

type Reader interface {
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	FindJWTRevocation(ctx context.Context, jti string) (*entity.JWTRevocation, error)
	FindInstitutionByDomains(ctx context.Context, domain string) (*entity.Institution, error)
}

type Writer interface {
	CreateUser(ctx context.Context, firstName string, lastName string, email string, password hash.Encoded, opts ...user.Option) (*entity.User, error)
	CreateJWTRevocation(ctx context.Context, jti string, expiry time.Time) (*entity.JWTRevocation, error)
}

type Repository interface {
	Reader
	Writer
}
