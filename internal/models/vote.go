package models

type Vote struct {
	Likes    uint64 `json:"likes"`
	Dislikes uint64 `json:"dislikes"`
}

type PostVote struct {
	ID       int
	PostID   int
	UserID   int
	LikeType int
}

type CommentVote struct {
	ID        int
	CommentID int
	UserID    int
	LikeType  int
}

var LikeTypes = struct {
	Like    int
	Dislike int
}{
	Like:    1,
	Dislike: 2,
}
