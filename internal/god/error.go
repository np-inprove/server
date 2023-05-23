package god

import (
	"errors"
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
			AppErrMessage:  http.StatusText(http.StatusContinue),
		}
	}

	return apperror.ErrInternal(err)
}
