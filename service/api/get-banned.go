package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) getBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id := ps.ByName("userid")
	userId, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user_id_banned := ps.ByName("useridbanned")
	userIdBanned, err := strconv.ParseUint(user_id_banned, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dbban, err := rt.db.GetBanned(userIdBanned, userId)

	if err != nil {
		ctx.Logger.WithError(err).Error("Can't get the ban")
		utils.ErrorTranslate(w, err)
		return
	}

	var ban Ban

	ban.FromDatabase(dbban)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ban)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
