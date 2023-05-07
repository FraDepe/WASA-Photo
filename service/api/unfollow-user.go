package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id := ps.ByName("userid")
	userId, err := strconv.ParseUint(user_id, 10, 64)
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

	if userId != loggedUser {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	user_id_to_unfollow := ps.ByName("useridtofollow")
	userIdToUnfollow, err := strconv.ParseUint(user_id_to_unfollow, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UnfollowUser(loggedUser, userIdToUnfollow)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't follow the user")
		utils.ErrorTranslate(w, err)
		return
	}

}
