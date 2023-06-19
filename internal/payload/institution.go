package payload

import (
	"github.com/gookit/validate"
	"net/http"
)

type Institution struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ShortName   string `json:"shortName,omitempty"`
	Description string `json:"description,omitempty"`
}

func (i Institution) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type CreateInstitutionRequest struct {
	Name        string `json:"name" validate:"required"`
	ShortName   string `json:"shortName" validate:"required|alphaDash"`
	Description string `json:"description"`
}

func (c CreateInstitutionRequest) Validate() *validate.Validation {
	return validate.Struct(c)
}

type UpdateInstitutionRequest struct {
	Name        string `json:"name" validate:"required"`
	ShortName   string `json:"shortName" validate:"required|alphaDash"`
	Description string `json:"description"`
}

func (u UpdateInstitutionRequest) Validate() *validate.Validation {
	return validate.Struct(u)
}
