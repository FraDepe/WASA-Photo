package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) getComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	logged_user_id := r.Header.Get("Authorization")
	loggedUserId, err := strconv.ParseUint(logged_user_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment_id := ps.ByName("commentid")
	commentId, err := strconv.ParseUint(comment_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dbcomment, err := rt.db.GetComment(commentId, loggedUserId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var comment Comment

	comment.FromDatabase(dbcomment)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
