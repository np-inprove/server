package user

import "time"

type Options struct {
	DeptID                 int
	Points                 int
	PointsAwardedCount     int
	PointsAwardedResetTime time.Time
	GodMode                bool
}

type Option func(*Options)

func DeptID(id int) Option {
	return func(opts *Options) {
		opts.DeptID = id
	}
}

func Points(points int) Option {
	return func(opts *Options) {
		opts.Points = points
	}
}

func PointsAwardedCount(c int) Option {
	return func(opts *Options) {
		opts.PointsAwardedCount = c
	}
}

func PointsAwardedResetTime(t time.Time) Option {
	return func(opts *Options) {
		opts.PointsAwardedResetTime = t
	}
}

func GodMode(gm bool) Option {
	return func(opts *Options) {
		opts.GodMode = gm
	}
}
