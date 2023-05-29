package auth

import (
	"github.com/np-inprove/server/internal/ent"
	"time"
)

type User = ent.User
type JWTRevocation = ent.JWTRevocation
type Institution = ent.Institution

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
