package payload

import (
	"github.com/gookit/validate"
	"net/http"
	"regexp"
)

const (
	fqdnRegexStringRFC1123 = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$` // From go-playground/validator, same as hostnameRegexStringRFC1123 but must contain a non numerical TLD (possibly ending with '.')
)

var (
	fqdnRegexRFC1123 = regexp.MustCompile(fqdnRegexStringRFC1123)
)

type Institution struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	ShortName     string `json:"shortName,omitempty"`
	AdminDomain   string `json:"adminDomain,omitempty"`
	StudentDomain string `json:"studentDomain,omitempty"`
}

func (i Institution) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type CreateInstitutionRequest struct {
	Name          string `json:"name" validate:"required"`
	ShortName     string `json:"shortName" validate:"required|alphaDash"`
	AdminDomain   string `json:"adminDomain,omitempty" validate:"required|fqdn"`
	StudentDomain string `json:"studentDomain,omitempty" validate:"required|fqdn"`
}

func (c CreateInstitutionRequest) Validate() *validate.Validation {
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