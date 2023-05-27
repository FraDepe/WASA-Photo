package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) getLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	photo_id := ps.ByName("photoid")
	photoId, err := strconv.ParseUint(photo_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user_id := ps.ByName("userid")
	userId, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dblike, err := rt.db.GetLike(photoId, userId)

	if err != nil {
		ctx.Logger.WithError(err).Error("Can't get the like")
		utils.ErrorTranslate(w, err)
		return
	}

	var like Like

	like.FromDatabase(dblike)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(like)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
