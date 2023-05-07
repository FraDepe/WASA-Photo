package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

	user_id_to_unban := ps.ByName("useridtoban")
	userIdToUnban, err := strconv.ParseUint(user_id_to_unban, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UnbanUser(loggedUser, userIdToUnban)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't unban the user")
		utils.ErrorTranslate(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
