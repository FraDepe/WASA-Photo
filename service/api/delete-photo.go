package api

import (
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// VA FATTO UN CONTROLLO DELL'AUTHENTICATION

	photo_id := ps.ByName("photoId")
	if photo_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.db.DeletePhoto(photo_id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
