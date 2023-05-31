package auth

import (
	"github.com/np-inprove/server/internal/ent"
	"time"
)

type User = ent.User
type JWTRevocation = ent.JWTRevocation
type Institution = ent.Institution

type CreateUserOptions struct {
	deptID                 int
	points                 int
	pointsAwardedCount     int
	pointsAwardedResetTime time.Time
	godMode                bool
}

type CreateUserOption func(*CreateUserOptions)

func DeptID(id int) CreateUserOption {
	return func(opts *CreateUserOptions) {
		opts.deptID = id
	}
}

func Points(points int) CreateUserOption {
	return func(opts *CreateUserOptions) {
		opts.points = points
	}
}

func PointsAwardedCount(c int) CreateUserOption {
	return func(opts *CreateUserOptions) {
		opts.pointsAwardedCount = c
	}
}

func PointsAwardedResetTime(t time.Time) CreateUserOption {
	return func(opts *CreateUserOptions) {
		opts.pointsAwardedResetTime = t
	}
}

func GodMode(gm bool) CreateUserOption {
	return func(opts *CreateUserOptions) {
		opts.godMode = gm
	}
}
