package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	var err error

	photo.Picture, err = io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photo.Comments = 0
	photo.Likes = 0
	photo.Date_time = time.Now().Format("2017-07-21T17:32:28Z")

	dbphoto, err := rt.db.UploadPhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't upload the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photo.FromDatabase(dbphoto)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)
}
