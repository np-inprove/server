package payload

import (
	"github.com/gookit/validate"
	"net/http"
)

type Institution struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	ShortName     string `json:"short_name,omitempty"`
	AdminDomain   string `json:"admin_domain,omitempty"`
	StudentDomain string `json:"student_domain,omitempty"`
}

func (i Institution) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type CreateInstitutionRequest struct {
	Name          string `json:"name" validate:"required"`
	ShortName     string `json:"short_name" validate:"required|alphaDash"`
	AdminDomain   string `json:"admin_domain,omitempty" validate:"required|dnsName"`
	StudentDomain string `json:"student_domain,omitempty" validate:"required|dnsName"`
}

func (c CreateInstitutionRequest) Validate() *validate.Validation {
	return validate.Struct(c)
}
