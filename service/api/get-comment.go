package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) getComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	comment_id := ps.ByName("commentId")
	commentId, err := strconv.ParseUint(comment_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if commentId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbcomment, err := rt.db.GetComment(commentId)

	if err == nil {
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
