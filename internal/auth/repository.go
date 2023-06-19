package auth

import (
	"context"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/institution"
	"github.com/np-inprove/server/internal/entity/user"
	"github.com/np-inprove/server/internal/hash"
	"time"
)

type Reader interface {
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	FindInstitutionInviteLinkWithInstitution(ctx context.Context, code string) (*entity.InstitutionInviteLink, error)
	FindJWTRevocation(ctx context.Context, jti string) (*entity.JWTRevocation, error)
}

type Writer interface {
	CreateUser(ctx context.Context, instID int, instRole institution.Role, firstName, lastName, email string, password hash.Encoded, opts ...user.Option) (*entity.User, error)
	CreateJWTRevocation(ctx context.Context, jti string, expiry time.Time) (*entity.JWTRevocation, error)
}

type Repository interface {
	Reader
	Writer
}
