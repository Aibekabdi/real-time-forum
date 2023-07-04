package models

type Comments struct {
	ID     uint   `json:"id"`
	Author User   `json:"author"`
	PostId uint   `json:"postid"`
	Vote   Vote   `json:"votes"`
	Text   string `json:"text"`
}
