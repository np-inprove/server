package forum

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/np-inprove/server/internal/apperror"
)

var (
	ErrGroupNotFound          = errors.New("group does not exist")
	ErrForumShortNameConflict = errors.New("a forum with the desired short name already exists")
	ErrUnauthorized           = errors.New("not authorized")
)

func mapDomainErr(err error) render.Renderer {
	if errors.Is(err, ErrGroupNotFound) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusNotFound,
			AppErrCode:     40402,
			AppErrMessage:  "The group specified does not exist",
		}
	}

	if errors.Is(err, ErrForumShortNameConflict) {
		return &apperror.ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusConflict,
			AppErrCode:     40902,
			AppErrMessage:  "This forum short name is unavailable",
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
