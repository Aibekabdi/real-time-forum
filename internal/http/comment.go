package http

import "net/http"

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

	} else {
		jsonResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) voteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

	} else {
		jsonResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}
