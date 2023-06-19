package models

type Comments struct {
	Id     uint   `json:"id"`
	Author User   `json:"author"`
	PostId uint   `json:"postid"`
	Title  string `json:"title"`
	Vote   Vote   `json:"votes"`
	Text   string `json:"text"`
}
