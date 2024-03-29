package models

type Post struct {
	ID       uint       `json:"id"`
	Author   User       `json:"author"`
	Title    string     `json:"title"`
	Tags     []Tags     `json:"tags"`
	Comments []Comments `json:"comments,omitempty"`
	Vote     Vote       `json:"votes,omitempty"`
	Text     string     `json:"text"`
}

type Tags struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
