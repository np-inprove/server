package god

import (
	"errors"
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/apperror"
	"net/http"
)

var (
	ErrShortNameConflict = errors.New("short name conflicts with existing institution")
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

	return apperror.ErrInternal(err)
}
