package models

type Session struct {
	ID       int    `json:"id"`
	UUID     string `json:"uuid"`
	UserID   int    `json:"userId"`
	DbCookie string `json:"dbCookie"`
}
