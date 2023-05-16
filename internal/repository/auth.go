package repository

import (
	"context"
	"time"

	"github.com/np-inprove/server/internal/entity"
)

type AuthReader interface {
	FindUserByEmail(ctx context.Context, email string) (entity.User, error)
	FindJWTRevocationByJTI(ctx context.Context, jti string) (entity.JWTRevocation, error)
}

type AuthWriter interface {
	CreateUser(
		ctx context.Context,
		firstName string,
		lastName string,
		email string,
		passwordHash string,
		opts ...UserOption,
	) (entity.User, error)
	CreateJWTRevocation(ctx context.Context, jti string, expiry time.Time) (entity.JWTRevocation, error)
}

type Auth interface {
	AuthReader
	AuthWriter
}

type userOptions struct {
	points                 int
	pointsAwardedCount     int
	pointsAwardedResetTime time.Time
	godMode                bool
}

type UserOption func(*userOptions)

func Points(points int) UserOption {
	return func(opts *userOptions) {
		opts.points = points
	}
}

func PointsAwardedCount(c int) UserOption {
	return func(opts *userOptions) {
		opts.pointsAwardedCount = c
	}
}

func PointsAwardedResetTime(t time.Time) UserOption {
	return func(opts *userOptions) {
		opts.pointsAwardedResetTime = t
	}
}

func GodMode(gm bool) UserOption {
	return func(opts *userOptions) {
		opts.godMode = gm
	}
}
