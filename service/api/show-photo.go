package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) showPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	photo_id := ps.ByName("photoId")
	if photo_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbphoto, err := rt.db.ShowPhoto(photo_id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
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
