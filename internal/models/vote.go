package models

type Vote struct {
	Likes    uint `json:"likes"`
	Dislikes uint `json:"dislikes"`
}

type PostVote struct {
	ID       uint `json:"id"`
	PostID   uint `json:"postID"`
	UserID   uint `json:"userID"`
	LikeType int  `json:"likeType"`
}

type CommentVote struct {
	ID        uint `json:"id"`
	CommentID uint `json:"commentID"`
	UserID    uint `json:"userID"`
	LikeType  int  `json:"likeType"`
}
