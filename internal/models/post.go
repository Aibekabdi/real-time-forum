package models

type Post struct {
	ID        int    `json:"id"`
	CreatorID int    `json:"creatorId"`
	Title     string `json:"title"`
	Tags      string `json:"tags"`
	Content   string `json:"content"`
}
