package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	username_buf, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	username := string(username_buf)
	if !utils.ValidUsername(username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user User
	user.Username = username

	dbuser, err := rt.db.DoLogin(username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't find or create the user")
		utils.ErrorTranslate(w, err)
		return
	}

	user.FromDatabase(dbuser)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
