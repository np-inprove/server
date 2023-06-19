package payload

import (
	"github.com/gookit/validate"
	"net/http"
)

type Group struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ShortName   string `json:"path,omitempty"`
	Description string `json:"description,omitempty"`
}

func (g Group) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type CreateGroupRequest struct {
	Name        string `json:"name,omitempty" validate:"required|minLen:3"`
	ShortName   string `json:"shortName,omitempty" validate:"required|alphaDash"`
	Description string `json:"description,omitempty"`
}

func (c CreateGroupRequest) Validate() *validate.Validation {
	return validate.Struct(c)
}
