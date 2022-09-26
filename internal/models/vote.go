package models

type PostRating struct {
	ID          int `json:"id"`
	Likes       int `json:"likes"`
	PostID      int `json:"postId"`
	LikedUserID int `json:"likedUserId"`
}

type CommentRating struct {
	ID          int `json:"id"`
	Likes       int `json:"likes"`
	CommentID   int `json:"commentId"`
	LikedUserID int `json:"likedUserId"`
}
