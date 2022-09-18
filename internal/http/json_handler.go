package http

import "net/http"

func jsonResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.WriteHeader(status)
}
