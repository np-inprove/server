package auth

import (
	"context"
	"github.com/np-inprove/server/internal/hash"
	"time"
)

type Reader interface {
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	FindJWTRevocation(ctx context.Context, jti string) (*JWTRevocation, error)
	FindInstitutionByDomains(ctx context.Context, domain string) (*Institution, error)
}

type Writer interface {
	CreateUser(ctx context.Context, firstName string, lastName string, email string, password hash.Encoded, opts ...CreateUserOption) (*User, error)
	CreateJWTRevocation(ctx context.Context, jti string, expiry time.Time) (*JWTRevocation, error)
}

type Repository interface {
	Reader
	Writer
}
