package api

import (
	"net/http"
	"strconv"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id := ps.ByName("userId")
	photo_id := ps.ByName("photoId")
	logged_user := r.Header.Get("Authorization")
	userid, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photoid, err := strconv.ParseUint(photo_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	actual_user_id, err := strconv.ParseUint(logged_user, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userid == actual_user_id {
		err := rt.db.UnlikePhoto(userid, photoid)
		if err != nil {
			ctx.Logger.WithError(err).Error("Can't remove like")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}
