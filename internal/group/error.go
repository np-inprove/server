package group

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/gookit/validate"
	"github.com/np-inprove/server/internal/apperror"
	"net/http"
)

var (
	ErrGroupNotFound          = errors.New("group does not exist")
	ErrGroupShortNameConflict = errors.New("a group with the desired short name already exists")
	ErrUnauthorized           = errors.New("not authorized")
	ErrInviteLinkNotFound     = errors.New("invite link does not exist")
)

func mapDomainErr(err error) render.Renderer {
	if errors.Is(err, ErrGroupNotFound) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusNotFound,
			AppErrCode:     40401,
			AppErrMessage:  "The group specified does not exist",
			Fields: validate.Errors{
				"shortName": map[string]string{
					"notFound": "group does not exist",
				},
			},
		}
	}

	if errors.Is(err, ErrGroupShortNameConflict) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusConflict,
			AppErrCode:     40901,
			AppErrMessage:  "This group short name is unavailable",
			Fields: validate.Errors{
				"shortName": map[string]string{
					"conflict": "Short name already in use",
				},
			},
		}
	}

	if errors.Is(err, ErrUnauthorized) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusUnauthorized,
			AppErrCode:     40301,
			AppErrMessage:  "Unauthorized",
			Fields: validate.Errors{
				"user": map[string]string{
					"unauthorized": "You must be an owner or educator to perform this action",
				},
			},
		}
	}

	if errors.Is(err, ErrInviteLinkNotFound) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusNotFound,
			AppErrCode:     40402,
			AppErrMessage:  "The invite link does not exist",
			Fields: validate.Errors{
				"code": map[string]string{
					"notFound": "The invite link does not exist",
				},
			},
		}
	}

	return apperror.ErrInternal(err)
}
