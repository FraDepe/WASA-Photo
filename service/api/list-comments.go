package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) listComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// We need the id of logged user to check if he is banned
	logged_user_id := r.Header.Get("Authorization")
	loggedUserId, err := strconv.ParseUint(logged_user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photo_id := ps.ByName("photoid")
	photoId, err := strconv.ParseUint(photo_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stream_comments, err := rt.db.ListComments(photoId, loggedUserId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(stream_comments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
