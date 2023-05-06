package database

import "net/http"

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
