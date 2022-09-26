package http

import "net/http"

func (h *Handler) posts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

	} else {
		jsonResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

	} else {
		jsonResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

	} else {
		jsonResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) votePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

	} else {
		jsonResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}
