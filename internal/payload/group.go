package payload

import (
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/entity/group"
	"golang.org/x/exp/slices"
	"net/http"
)

type Group struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ShortName   string `json:"shortName,omitempty"`
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

type UpdateGroupRequest struct {
	Name        string `json:"name,omitempty" validate:"required|minLen:3"`
	ShortName   string `json:"shortName,omitempty" validate:"required|alphaDash"`
	Description string `json:"description,omitempty"`
}

func (u UpdateGroupRequest) Validate() *validate.Validation {
	return validate.Struct(u)
}

type GroupInviteLink struct {
	ID    int        `json:"id,omitempty"`
	Code  string     `json:"code,omitempty"`
	Role  group.Role `json:"role,omitempty"`
	Group Group      `json:"group,omitempty"`
}

func (g GroupInviteLink) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type CreateGroupInviteLinkRequest struct {
	Role group.Role `json:"role,omitempty"`
}

func (c CreateGroupInviteLinkRequest) Validate() *validate.Validation {
	v := validate.Struct(c)
	v.AddValidator("role", func(val interface{}) bool {
		return slices.Contains(group.Roles, val.(group.Role))
	})
	return v
}
