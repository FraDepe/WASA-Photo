package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id := ps.ByName("userId")
	photo_id := ps.ByName("photoId")
	userId, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	photoId, err := strconv.ParseUint(photo_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logged_user := r.Header.Get("Authorization")
	loggedUser, err := strconv.ParseUint(logged_user, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var like Like
	like.PhotoId = photoId

	if userId == loggedUser {
		like.UserId = userId
	} else {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	dblike, err := rt.db.LikePhoto(like.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't post the comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	like.FromDatabase(dblike)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(like)

}
