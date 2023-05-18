package auth

import (
	"github.com/gookit/validate"
	"net/http"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required|email"`
	Password string `json:"password"`
}

func (l LoginRequest) Validate() *validate.Validation {
	return validate.Struct(l)
}

type WhoAmIResponse struct {
	FirstName              string    `json:"first_name,omitempty"`
	LastName               string    `json:"last_name,omitempty"`
	Email                  string    `json:"email,omitempty"`
	Points                 int       `json:"points,omitempty"`
	PointsAwardedCount     int       `json:"points_awarded_count,omitempty"`
	PointsAwardedResetTime time.Time `json:"points_awarded_reset_time,omitempty"`
	GodMode                bool      `json:"god_mode,omitempty"`
}

func (wai WhoAmIResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
