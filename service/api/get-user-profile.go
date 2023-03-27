package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	username := ps.ByName("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbuser, err := rt.db.GetUserProfile(username)

	if err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var user User

	user.FromDatabase(dbuser)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
