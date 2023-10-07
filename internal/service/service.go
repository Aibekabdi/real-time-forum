package service

import (
	"context"
	"forum/internal/models"
	"forum/internal/repository"
)

type Auth interface {
	Create(ctx context.Context, user models.User) error
	SignIn(ctx context.Context, user models.SigningInput) (string, error)
	ParseToken(accessToken string) (models.UserToken, error)
}

type Post interface {
	Create(ctx context.Context, post models.Post) (uint, error)
	Delete(ctx context.Context, postID, userID uint) error
	GetALL(ctx context.Context) ([]models.Post, error)
	GetByID(ctx context.Context, postID uint) (models.Post, error)
	InsertorDelete(ctx context.Context, input models.PostVote) error
	GetALLByTag(ctx context.Context, tagName string) ([]models.Post, error)
	GetALLByUserID(ctx context.Context, userID uint) ([]models.Post, error)
}

type Comment interface {
	Create(ctx context.Context, comment models.Comments) (uint, error)
	Delete(ctx context.Context, commentID, userID uint) error
	InsertorDelete(ctx context.Context, input models.CommentVote) error
}
type User interface {
	UpdatePassword(ctx context.Context, updatePsw models.UpdatePassword, userID uint) error
	GetUserInfo(ctx context.Context, userID uint) (models.User, error)
}

type Chat interface {
	GetMessages(ctx context.Context, senderID, receiverID, lastMessageID uint) ([]models.Message, error)
	Create(ctx context.Context, senderID, receiverID uint, content string) (models.Message, error)
}
type Service struct {
	Auth
	Post
	Comment
	User
	Chat
}

func NewService(repo *repository.Repository, secretKey string) *Service {
	return &Service{
		Auth:    newAuthService(repo.User, secretKey),
		Post:    newPostService(repo.Post, repo.Tag, repo.Comment),
		Comment: newCommentService(repo.Comment),
		User:    newUserService(repo.User),
		Chat:    nil,
	}
}
