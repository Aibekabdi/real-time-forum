package http

import (
	"net/http"
)

func (h *Handler) isSessionValid(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session")
		if err != nil {
			jsonResponse(w, r, http.StatusUnauthorized, "session expired or not valid")
			return
		} else {
			f.ServeHTTP(w, r)
		}
	}
}
