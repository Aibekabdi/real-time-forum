package models

type Post struct {
	ID       uint       `json:"id"`
	Author   User       `json:"author"`
	Title    string     `json:"title"`
	Tags     []Tags     `json:"tags"`
	Comments []Comments `json:"comments"`
	Vote     Vote       `json:"votes"`
	Text     string     `json:"text"`
}

type Tags struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}
