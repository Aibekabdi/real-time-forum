package service

import (
	"context"
	"errors"
	"forum/internal/models"
	"forum/internal/repository"
	"strings"
	"unicode"
)

type CommentService struct {
	commentRepo repository.Comment
}

func newCommentService(commentRepo repository.Comment) *CommentService {
	return &CommentService{commentRepo: commentRepo}
}

func (s *CommentService) Create(ctx context.Context, comment models.Comments) (uint, error) {
	text := strings.TrimFunc(comment.Text, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	if len(text) <= 0 {
		return 0, errors.New("comment's text is null")
	}
	return s.commentRepo.Create(ctx, comment)
}

func (s *CommentService) Delete(ctx context.Context, commentID, userID uint) error {
	return s.commentRepo.Delete(ctx, commentID, userID)
}

func (s *CommentService) InsertorDelete(ctx context.Context, commentID, userID uint, likeType int) error {
	if likeType != -1 && likeType != 1 {
		return errors.New("invalid type of vote")
	}
	return s.commentRepo.InsertorDelete(ctx, commentID, userID, likeType)
}
