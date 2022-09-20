package http

import (
	"log"
	"net/http"
)

func (h *Handler) isSessionValid(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		_, err := r.Cookie("session")
		if err != nil {
			log.Println("session expired")
			jsonResponse(w, r, http.StatusUnauthorized, "session expired or not valid")
			return
		} else {
			f.ServeHTTP(w, r)
		}
	}
}
