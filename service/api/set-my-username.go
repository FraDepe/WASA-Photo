package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	logged_user_id := r.Header.Get("Authorization")

	loggedUserId, err := strconv.ParseUint(logged_user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user_id := ps.ByName("userid")
	userid, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if loggedUserId != userid {
		w.WriteHeader(http.StatusForbidden)
		return
	}

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

	err = rt.db.SetMyUsername(username, userid)
	if err != nil {
		utils.ErrorTranslate(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
