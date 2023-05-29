package payload

import (
	"github.com/gookit/validate"
	"net/http"
	"time"
)

type User struct {
	ID                     int       `json:"id,omitempty"`
	FirstName              string    `json:"first_name,omitempty"`
	LastName               string    `json:"last_name,omitempty"`
	Email                  string    `json:"email,omitempty"`
	Points                 int       `json:"points,omitempty"`
	PointsAwardedCount     int       `json:"points_awarded_count,omitempty"`
	PointsAwardedResetTime time.Time `json:"points_awarded_reset_time,omitempty"`
	GodMode                bool      `json:"god_mode,omitempty"`
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
	FirstName       string `json:"first_name,omitempty" validate:"required|min_len:1"`
	LastName        string `json:"last_name,omitempty" validate:"required|min_len:1"`
	Email           string `json:"email,omitempty" validate:"required|email"`
	Password        string `json:"password" validate:"required|min_len:6"`
	ConfirmPassword string `json:"confirm_password" validate:"required|min_len:6"`
}

func (r RegisterRequest) Validate() *validate.Validation {
	return validate.Struct(r)
}
