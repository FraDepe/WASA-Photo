package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

	user_id_to_ban := ps.ByName("useridtoban")
	userIdToBan, err := strconv.ParseUint(user_id_to_ban, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dbuser, err := rt.db.BanUser(loggedUser, userIdToBan)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't ban the user")
		utils.ErrorTranslate(w, err)
		return
	}

	err = rt.db.UnfollowUser(userIdToBan, loggedUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't ban the user")
		utils.ErrorTranslate(w, err)
		return
	}

	var user User

	user.FromDatabase(dbuser)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
	return
}
