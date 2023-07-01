package payload

import (
	"net/http"

	"github.com/gookit/validate"
)

type Forum struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ShortName   string `json:"path,omitempty"`
	Description string `json:"description,omitempty"`
}

func (f Forum) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type CreateForumRequest struct {
	Name        string `json:"name,omitempty" validate:"required|minLen:3"`
	ShortName   string `json:"shortName,omitempty" validate:"required|alphaDash"`
	Description string `json:"description,omitempty"`
}

func (c CreateForumRequest) Validate() *validate.Validation {
	return validate.Struct(c)
}

type UpdateForumRequest struct {
	Name        string `json:"name,omitempty" validate:"required|minLen:3"`
	ShortName   string `json:"shortName,omitempty" validate:"required|alphaDash"`
	Description string `json:"description,omitempty"`
}

func (u UpdateForumRequest) Validate() *validate.Validation {
	return validate.Struct(u)
}
