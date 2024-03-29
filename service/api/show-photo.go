package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) showPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// We need the id of logged user to check if he is banned
	logged_user_id := r.Header.Get("Authorization")
	loggedUserId, err := strconv.ParseUint(logged_user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photo_id := ps.ByName("photoid")
	photoid, err := strconv.ParseUint(photo_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbphoto, err := rt.db.ShowPhoto(photoid, loggedUserId)

	if err != nil {
		ctx.Logger.WithError(err).Error("Can't load the photo")
		utils.ErrorTranslate(w, err)
		return
	}

	var photo Photo

	photo.FromDatabase(dbphoto)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(photo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
