package payload

import (
	"net/http"

	"github.com/gookit/validate"
)

type ForumPost struct {
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Content  string `json:"content,omitempty"`
	TimeDate string `json:"description,omitempty"`
}

func (fp ForumPost) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type CreateForumPostRequest struct {
	Name        string `json:"name,omitempty" validate:"required|minLen:3"`
	ShortName   string `json:"shortName,omitempty" validate:"required|alphaDash"`
	Description string `json:"description,omitempty"`
}

func (c CreateForumPostRequest) Validate() *validate.Validation {
	return validate.Struct(c)
}

type UpdateForumPostRequest struct {
	Name        string `json:"name,omitempty" validate:"required|minLen:3"`
	ShortName   string `json:"shortName,omitempty" validate:"required|alphaDash"`
	Description string `json:"description,omitempty"`
}

func (u UpdateForumPostRequest) Validate() *validate.Validation {
	return validate.Struct(u)
}
