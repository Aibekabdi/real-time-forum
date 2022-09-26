package http

import "net/http"

func (h *Handler) profile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//todo signup

	} else {
		jsonResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}
