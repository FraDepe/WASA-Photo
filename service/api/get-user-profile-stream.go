package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) getUserProfileStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	username := ps.ByName("username")

	if !utils.ValidUsername(username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user_id := ps.ByName("userid")
	loggedUser, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stream, err := rt.db.GetUserProfileStream(username, loggedUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't get the stream")
		utils.ErrorTranslate(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(stream)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()
}
