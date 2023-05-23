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

type createUserOptions struct {
	deptID                 int
	points                 int
	pointsAwardedCount     int
	pointsAwardedResetTime time.Time
	godMode                bool
}

type CreateUserOption func(*createUserOptions)

func DeptID(id int) CreateUserOption {
	return func(opts *createUserOptions) {
		opts.deptID = id
	}
}

func Points(points int) CreateUserOption {
	return func(opts *createUserOptions) {
		opts.points = points
	}
}

func PointsAwardedCount(c int) CreateUserOption {
	return func(opts *createUserOptions) {
		opts.pointsAwardedCount = c
	}
}

func PointsAwardedResetTime(t time.Time) CreateUserOption {
	return func(opts *createUserOptions) {
		opts.pointsAwardedResetTime = t
	}
}

func GodMode(gm bool) CreateUserOption {
	return func(opts *createUserOptions) {
		opts.godMode = gm
	}
}
