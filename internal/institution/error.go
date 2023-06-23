package institution

import (
	"errors"
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/apperror"
	"net/http"
)

var (
	ErrInstitutionShortNameConflict = errors.New("short name conflicts with existing institution")
	ErrInstitutionNotFound          = errors.New("institution does not exist")
	ErrUnauthorized                 = errors.New("unauthorized")
)

func mapDomainErr(err error) *apperror.ErrResponse {
	if errors.Is(err, ErrInstitutionShortNameConflict) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusConflict,
			AppErrCode:     http.StatusConflict,
			AppErrMessage:  "Short name already in use",
			Fields: validate.Errors{
				"shortName": map[string]string{
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
				"shortName": map[string]string{
					"notFound": "Institution does not exist",
				},
			},
		}
	}

	if errors.Is(err, ErrUnauthorized) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusUnauthorized,
			AppErrCode:     40302,
			AppErrMessage:  "You must be an admin to perform this action",
			Fields: validate.Errors{
				"user": map[string]string{
					"unauthorized": "You must be an admin to perform this action",
				},
			},
		}
	}

	return apperror.ErrInternal(err)
}
