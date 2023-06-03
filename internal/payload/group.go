package payload

import (
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/ent/group"
	"net/http"
)

type Group struct {
	ID          int    `json:"id,omitempty"`
	Path        string `json:"path,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	GroupType   string `json:"groupType,omitempty"`
}

func (g Group) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type GroupType string

func (gt GroupType) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type CreateGroupRequest struct {
	Path        string `json:"path,omitempty" validate:"required|alphaDash"`
	Name        string `json:"name,omitempty" validate:"required|minLen:3"`
	Description string `json:"description,omitempty"`
	GroupType   string `json:"groupType,omitempty" validate:"customValidator"`
}

func (c CreateGroupRequest) CustomValidator(val string) bool {
	return group.GroupTypeValidator(group.GroupType(val)) == nil
}

func (c CreateGroupRequest) Validate() *validate.Validation {
	return validate.Struct(c)
}
