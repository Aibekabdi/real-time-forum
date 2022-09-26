package http

import "net/http"

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//todo signup

	} else {
		jsonResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//todo signin

	} else {
		jsonResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) signout(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		//todo signup

	} else {
		jsonResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}
