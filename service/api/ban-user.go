package api

import (
	"net/http"
	"strconv"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id := ps.ByName("userId")
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

	_, err = rt.db.BanUser(loggedUser, userId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't ban the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
