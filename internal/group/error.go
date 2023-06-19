package group

import (
	"errors"
	"github.com/np-inprove/server/internal/apperror"
	"net/http"
)

var (
	ErrInstitutionNotFound    = errors.New("institution does not exist")
	ErrGroupShortNameConflict = errors.New("a group with the desired short name already exists")
	ErrUnauthorized           = errors.New("not authorized")
)

func mapDomainErr(err error) *apperror.ErrResponse {
	if errors.Is(err, ErrInstitutionNotFound) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusNotFound,
			AppErrCode:     40401,
			AppErrMessage:  "The institution specified does not exist",
		}
	}

	if errors.Is(err, ErrGroupShortNameConflict) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusConflict,
			AppErrCode:     40901,
			AppErrMessage:  "This group short name is unavailable",
		}
	}

	if errors.Is(err, ErrUnauthorized) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusUnauthorized,
			AppErrCode:     40301,
			AppErrMessage:  "Unauthorized",
		}
	}

	return apperror.ErrInternal(err)
}
