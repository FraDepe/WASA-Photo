package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) listBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// We need the id of logged user
	logged_user_id := r.Header.Get("Authorization")
	loggedUserId, err := strconv.ParseUint(logged_user_id, 10, 64)
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

	if userId == loggedUserId {
		user_list, err := rt.db.ListBanned(loggedUserId)

		if err != nil {
			utils.ErrorTranslate(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(user_list)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
		return
	}

}
