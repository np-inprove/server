package auth

import (
	"context"
	"time"
)

type Reader interface {
	FindUser(ctx context.Context, opts ...UserOption) (User, error)
	FindJWTRevocation(ctx context.Context, jti string) (JWTRevocation, error)
}

type Writer interface {
	CreateUser(
		ctx context.Context,
		firstName string,
		lastName string,
		email string,
		passwordHash string,
		opts ...UserOption,
	) (User, error)
	CreateJWTRevocation(ctx context.Context, jti string, expiry time.Time) (JWTRevocation, error)
}

type Repository interface {
	Reader
	Writer
}

type userOptions struct {
	email                  string
	points                 int
	pointsAwardedCount     int
	pointsAwardedResetTime time.Time
	godMode                bool
}

type UserOption func(*userOptions)

func Email(email string) UserOption {
	return func(opts *userOptions) {
		opts.email = email
	}
}

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
