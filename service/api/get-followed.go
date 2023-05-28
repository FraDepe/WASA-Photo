package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) getFollowed(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id := ps.ByName("userid")
	userId, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user_id_followed := ps.ByName("useridfollowed")
	userIdFollowed, err := strconv.ParseUint(user_id_followed, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dbfollow, err := rt.db.GetFollowed(userIdFollowed, userId)

	if err != nil {
		ctx.Logger.WithError(err).Error("Can't get the follow")
		utils.ErrorTranslate(w, err)
		return
	}

	var follow Follow

	follow.FromDatabase(dbfollow)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(follow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
