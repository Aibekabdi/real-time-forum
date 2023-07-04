package service

import (
	"context"
	"forum/internal/models"
	"forum/internal/repository"
)

type CommentService struct {
	commentRepo repository.Comment
}

func newCommentService(commentRepo repository.Comment) *CommentService {
	return &CommentService{commentRepo: commentRepo}
}

func (s *CommentService) Create(ctx context.Context, comment models.Comments) (uint, error) {
	// TODO chechk is comment valid
	return s.commentRepo.Create(ctx, comment)
}
func (s *CommentService) Delete(ctx context.Context, commentID, userID uint) error {
	return s.commentRepo.Delete(ctx, commentID, userID)
}
