package models

type Comment struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	PostID      int    `json:"postId"`
	CommenterID int    `json:"commenterId"`
}
