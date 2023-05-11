package utils

import (
	"errors"
	"net/http"
	"regexp"
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

func ErrorTranslate(w http.ResponseWriter, err error) {

	if errors.Is(err, ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
	} else if errors.Is(err, ErrInternalServer) {
		w.WriteHeader(http.StatusInternalServerError)
	} else if errors.Is(err, ErrPhotoNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else if errors.Is(err, ErrLikeNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else if errors.Is(err, ErrUserAlreadyBanned) {
		w.WriteHeader(http.StatusForbidden)
	} else if errors.Is(err, ErrUserNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else if errors.Is(err, ErrBanned) {
		w.WriteHeader(http.StatusForbidden)
	} else if errors.Is(err, ErrMustFollow) {
		w.WriteHeader(http.StatusForbidden)
	} else if errors.Is(err, ErrPermissioneDenied) {
		w.WriteHeader(http.StatusForbidden)
	} else if errors.Is(err, ErrCommentNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func ValidUsername(username string) bool {
	isValid := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username)
	if isValid && len(username) >= 3 {
		return true
	}
	return false

}
