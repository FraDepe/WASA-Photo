package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	var err error

	logged_user_id := r.Header.Get("Authorization")
	loggedUserId, err := strconv.ParseUint(logged_user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user_id := ps.ByName("userid")
	userId, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userId != loggedUserId {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	image, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if http.DetectContentType(image) == "image/png" || http.DetectContentType(image) == "image/jpg" || http.DetectContentType(image) == "image/jpeg" {
		photo.Picture = image
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photo.Comments = 0
	photo.Likes = 0
	photo.Date_time = time.Now().Format("2017-07-21T17:32:28Z")

	dbphoto, err := rt.db.UploadPhoto(photo.ToDatabase(), userId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't upload the photo")
		utils.ErrorTranslate(w, err)
		return
	}

	photo.FromDatabase(dbphoto)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)
}
