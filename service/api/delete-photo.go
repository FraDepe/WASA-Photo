package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id := r.Header.Get("Authorization")
	casted_user_id, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photo_id := ps.ByName("photoid")
	if photo_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	casted_photo_id, err := strconv.ParseUint(photo_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.DeletePhoto(casted_photo_id, casted_user_id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
