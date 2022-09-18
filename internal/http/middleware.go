package http

import (
	"log"
	"net/http"
)

func (h *Handler) isSessionValid(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session")
		if err != nil {
			log.Println("session expired")

		}
	}
}
