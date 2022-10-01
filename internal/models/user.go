package models

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Age       string `json:"age"`
	Genger    string `json:"gender"`
}
