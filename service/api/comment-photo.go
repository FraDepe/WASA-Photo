package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/utils"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comment Comment
	var err error

	// Extracting id of logged user (who's commenting)
	user_id := r.Header.Get("Authorization")
	comment.UserId, err = strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photo_id := ps.ByName("photoid")
	photoId, err := strconv.ParseUint(photo_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment.PhotoId = photoId

	// Extracting the comment to insert into the database
	buffer, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comment.Text = string(buffer)

	// Entering databse function
	dbcomment, err := rt.db.CommentPhoto(comment.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't post the comment")
		utils.ErrorTranslate(w, err)
		return
	}

	comment.FromDatabase(dbcomment)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)

}
