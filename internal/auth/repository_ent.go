package auth

import (
	"context"
	"github.com/np-inprove/server/internal/ent"
	"time"
)

type entRepository struct {
	ent *ent.Client
}

func NewEntRepository(ent *ent.Client) Repository {
	return entRepository{ent}
}

func (e entRepository) FindUser(ctx context.Context, opts ...UserOption) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (e entRepository) FindJWTRevocation(ctx context.Context, jti string) (JWTRevocation, error) {
	//TODO implement me
	panic("implement me")
}

func (e entRepository) CreateUser(ctx context.Context, firstName string, lastName string, email string, passwordHash string, opts ...UserOption) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (e entRepository) CreateJWTRevocation(ctx context.Context, jti string, expiry time.Time) (JWTRevocation, error) {
	//TODO implement me
	panic("implement me")
}
