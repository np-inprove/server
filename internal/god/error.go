package god

import (
	"errors"
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/apperror"
	"net/http"
)

var (
	ErrShortNameConflict   = errors.New("short name conflicts with existing institution")
	ErrInstitutionNotFound = errors.New("institution does not exist")
)

func mapDomainErr(err error) *apperror.ErrResponse {
	if errors.Is(err, ErrShortNameConflict) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusConflict,
			AppErrCode:     http.StatusConflict,
			AppErrMessage:  "Short name already in use",
			Fields: validate.Errors{
				"short_name": map[string]string{
					"conflict": "Short name already in use",
				},
			},
		}
	}

	if errors.Is(err, ErrInstitutionNotFound) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusNotFound,
			AppErrCode:     http.StatusNotFound,
			AppErrMessage:  "Institution does not exist",
			Fields: validate.Errors{
				"short_name": map[string]string{
					"notFound": "Institution does not exist",
				},
			},
		}
	}

	return apperror.ErrInternal(err)
}
