package auth

import "errors"

var (
	ErrInvalidPassword = errors.New("password provided did not match hash")
)
