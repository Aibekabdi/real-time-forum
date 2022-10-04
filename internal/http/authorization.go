package http

import (
	"log"
	"net/http"
)

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		_, _, user, _, _, err := GetJsonData(w, r, "user")
		log.Println(user)
		if err != nil {
			jsonResponse(w, r, http.StatusBadRequest, err.Error())
			return
		}
		if err := h.services.Authorization.Signup(user); err != nil {
			jsonResponse(w, r, http.StatusBadRequest, err.Error())
			return
		}
		jsonResponse(w, r, http.StatusCreated, nil)
	} else {
		jsonResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//todo signin
		//if login succeded , redirect to main page
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
