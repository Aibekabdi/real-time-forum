package service

import (
	"context"
	"fmt"
	"forum/internal/models"
	"forum/internal/repository"
	"forum/pkg/utils"
)

type PostService struct {
	postRepo repository.Post
	tagRepo  repository.Tag
}

func newPostService(postRepo repository.Post, tagRepo repository.Tag) *PostService {
	return &PostService{postRepo: postRepo, tagRepo: tagRepo}
}

func (s *PostService) Create(ctx context.Context, post models.Post) (uint, error) {
	if err := utils.PostValidation(&post); err != nil {
		return 0, fmt.Errorf("post service: sign in: %w", err)
	}
	postID, err := s.postRepo.Create(ctx, post)
	if err != nil {
		return 0, fmt.Errorf("post service: sign in: %w", err)
	}
	if err := s.tagRepo.Create(ctx, post.Tags, postID); err != nil {
		return 0, fmt.Errorf("post service: sign in: %w", err)
	}
	return postID, nil
}

func (s *PostService) Delete(ctx context.Context, postID, userID uint) error {
	return s.postRepo.Delete(ctx, postID, userID)
}

func (s *PostService) GetALL(ctx context.Context) ([]models.Post, error) {
	return s.postRepo.GetALL(ctx)
}
