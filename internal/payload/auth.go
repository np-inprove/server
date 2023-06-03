package payload

import (
	"github.com/gookit/validate"
	"net/http"
	"time"
)

type User struct {
	ID                     int       `json:"id,omitempty"`
	FirstName              string    `json:"firstName,omitempty"`
	LastName               string    `json:"lastName,omitempty"`
	Email                  string    `json:"email,omitempty"`
	Points                 int       `json:"points,omitempty"`
	PointsAwardedCount     int       `json:"pointsAwardedCount,omitempty"`
	PointsAwardedResetTime time.Time `json:"pointsAwardedResetTime,omitempty"`
	GodMode                bool      `json:"godMode,omitempty"`
}

func (u User) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required|email"`
	Password string `json:"password"`
}

func (l LoginRequest) Validate() *validate.Validation {
	return validate.Struct(l)
}

type RegisterRequest struct {
	FirstName string `json:"firstName,omitempty" validate:"required|min_len:1"`
	LastName  string `json:"lastName,omitempty" validate:"required|min_len:1"`
	Email     string `json:"email,omitempty" validate:"required|email"`
	Password  string `json:"password" validate:"required|min_len:6"`
}

func (r RegisterRequest) Validate() *validate.Validation {
	return validate.Struct(r)
}
