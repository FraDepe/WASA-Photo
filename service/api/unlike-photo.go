package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id := ps.ByName("userid")
	photo_id := ps.ByName("photoid")
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
			utils.ErrorTranslate(w, err)
			return
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
		return
	}

}
