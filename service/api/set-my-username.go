package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id := ps.ByName("userId")
	userid, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var new_username string

	err = json.NewDecoder(r.Body).Decode(&new_username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.SetMyUsername(new_username, userid)
	if err != nil { // ???
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}