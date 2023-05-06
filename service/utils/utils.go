package utils

import (
	"errors"
	"net/http"
)

// implementa globalmente
var ErrUserDoesNotExist = errors.New("user doesn't exist")
var ErrInternalServer = errors.New("internal server error")
var ErrPhotoNotFound = errors.New("photo not found")
var ErrLikeNotFound = errors.New("like not found")
var ErrUserAlreadyBanned = errors.New("user already banned")
var ErrUserNotFound = errors.New("user not found")
var ErrBanned = errors.New("you can't do it")
var ErrMustFollow = errors.New("you must follow before")
var ErrPermissioneDenied = errors.New("permission denied")
var ErrCommentNotFound = errors.New("comment not found")

func errorTranslate(w http.ResponseWriter, err error) {
	switch err {
	case ErrUserDoesNotExist:
		w.WriteHeader(http.StatusNotFound)
	case ErrInternalServer:
		w.WriteHeader(http.StatusInternalServerError)
	case ErrPhotoNotFound:
		w.WriteHeader(http.StatusNotFound)
	case ErrLikeNotFound:
		w.WriteHeader(http.StatusNotFound)
	case ErrUserAlreadyBanned:
		w.WriteHeader(http.StatusForbidden)
	case ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
	case ErrBanned:
		w.WriteHeader(http.StatusForbidden)
	case ErrMustFollow:
		w.WriteHeader(http.StatusForbidden)
	case ErrPermissioneDenied:
		w.WriteHeader(http.StatusForbidden)
	case ErrCommentNotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
