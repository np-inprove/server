package dash

import (
	"errors"
	"github.com/np-inprove/server/internal/apperror"
	"net/http"
)

var (
	ErrInvalidGroupType    = errors.New("group type used is invalid")
	ErrInstitutionNotFound = errors.New("institution does not exist")
	ErrGroupPathConflict   = errors.New("a group with the desired path already exists")
	ErrUserNotAdmin        = errors.New("user tried to perform an admin-only action")
)

func mapDomainErr(err error) *apperror.ErrResponse {
	if errors.Is(err, ErrInvalidGroupType) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusBadGateway,
			AppErrCode:     40101,
			AppErrMessage:  "The group type used is invalid",
		}
	}
	if errors.Is(err, ErrInstitutionNotFound) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusNotFound,
			AppErrCode:     40401,
			AppErrMessage:  "The institution specified does not exist",
		}
	}

	if errors.Is(err, ErrGroupPathConflict) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusConflict,
			AppErrCode:     40901,
			AppErrMessage:  "The group path chosen is unavailable",
		}
	}

	if errors.Is(err, ErrUserNotAdmin) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusUnauthorized,
			AppErrCode:     40301,
			AppErrMessage:  "You must be an admin to perform this action",
		}
	}

	return apperror.ErrInternal(err)
}
