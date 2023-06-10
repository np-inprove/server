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
	ID            int    `json:"id" validate:"required"`
	Name          string `json:"name" validate:"required"`
	ShortName     string `json:"shortName" validate:"required|alphaDash"`
	AdminDomain   string `json:"adminDomain,omitempty" validate:"required|fqdn"`
	StudentDomain string `json:"studentDomain,omitempty" validate:"required|fqdn"`
}

func (c UpdateInstitutionRequest) Validate() *validate.Validation {
	v := validate.Struct(c)
	v.AddValidator("fqdn", func(val interface{}) bool {
		s, ok := val.(string)
		if !ok {
			return false
		}

		if s == "" {
			return false
		}

		return fqdnRegexRFC1123.MatchString(s)
	})
	v.AddMessages(map[string]string{
		"fqdn": "value is not a fully qualified domain name",
	})
	return v
}