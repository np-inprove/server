package payload

import (
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/entity/institution"
	"golang.org/x/exp/slices"
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

type InstitutionInviteLink struct {
	ID          int              `json:"id,omitempty"`
	Code        string           `json:"code,omitempty"`
	Role        institution.Role `json:"role,omitempty"`
	Institution Institution      `json:"institution,omitempty"`
}

func (i InstitutionInviteLink) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type CreateInviteLinkRequest struct {
	Role institution.Role `json:"role,omitempty"`
}

func (c CreateInviteLinkRequest) Validate() *validate.Validation {
	v := validate.Struct(c)
	v.AddValidator("role", func(val interface{}) bool {
		return slices.Contains(institution.Roles, val.(institution.Role))
	})
	return v
}
