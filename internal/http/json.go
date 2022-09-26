package http

import (
	"encoding/json"
	"forum/internal/models"
	"io/ioutil"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
}

func GetJsonData(w http.ResponseWriter, r *http.Request, class string) (*models.Comment, *models.Post, *models.User, *models.CommentRating, *models.PostRating, error) {
	var c models.Comment
	var p models.Post
	var u models.User
	var cr models.CommentRating
	var pr models.PostRating
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	switch class {
	case "comment":
		err = json.Unmarshal(body, &c)
	case "post":
		err = json.Unmarshal(body, &p)
	case "user":
		err = json.Unmarshal(body, &u)
	case "commentRating":
		err = json.Unmarshal(body, &cr)
	case "postRating":
		err = json.Unmarshal(body, &pr)
	}
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	return &c, &p, &u, &cr, &pr, nil
}
