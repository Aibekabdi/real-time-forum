package service

import (
	"context"
	"errors"
	"fmt"
	"forum/internal/models"
	"forum/internal/repository"
	"forum/pkg/utils"
	"strings"
	"unicode"
)

type PostService struct {
	postRepo    repository.Post
	tagRepo     repository.Tag
	commentRepo repository.Comment
}

func newPostService(postRepo repository.Post, tagRepo repository.Tag, commentRepo repository.Comment) *PostService {
	return &PostService{postRepo: postRepo, tagRepo: tagRepo, commentRepo: commentRepo}
}

func (s *PostService) Create(ctx context.Context, post models.Post) (uint, error) {
	if err := utils.PostValidation(&post); err != nil {
		return 0, fmt.Errorf("post service: Create: %w", err)
	}
	postID, err := s.postRepo.Create(ctx, post)
	if err != nil {
		return 0, fmt.Errorf("post service: Create: %w", err)
	}
	if err := s.tagRepo.Create(ctx, post.Tags, postID); err != nil {
		return 0, fmt.Errorf("post service: Create: %w", err)
	}
	return postID, nil
}

func (s *PostService) Delete(ctx context.Context, postID, userID uint) error {
	return s.postRepo.Delete(ctx, postID, userID)
}

func (s *PostService) GetALL(ctx context.Context) ([]models.Post, error) {
	return s.postRepo.GetALL(ctx)
}

func (s *PostService) GetByID(ctx context.Context, postID uint) (models.Post, error) {
	tags, err := s.tagRepo.GetByPostID(ctx, postID)
	if err != nil {
		return models.Post{}, fmt.Errorf("post service: get by id: %w", err)
	}
	comments, err := s.commentRepo.GetByPostID(ctx, postID)
	if err != nil {
		return models.Post{}, fmt.Errorf("post service: get by id: %w", err)
	}
	post, err := s.postRepo.GetByID(ctx, postID)
	if err != nil {
		return models.Post{}, fmt.Errorf("post service: get by id: %w", err)
	}
	post.Tags = tags
	post.Comments = comments
	return post, nil
}

func (s *PostService) InsertorDelete(ctx context.Context, input models.PostVote) error {
	if input.LikeType != -1 && input.LikeType != 1 {
		return errors.New("invalid type of vote")
	}
	return s.postRepo.InsertorDelete(ctx, input)
}

func (s *PostService) GetALLByTag(ctx context.Context, tagName string) ([]models.Post, error) {
	if len(strings.TrimFunc(tagName, func(r rune) bool {
		return unicode.IsSpace(r)
	})) == 0 {
		return nil, errors.New("invalid tag")
	}
	return s.postRepo.GetALLByTag(ctx, tagName)
}

func (s *PostService) GetALLByUserID(ctx context.Context, userID uint) ([]models.Post, error) {
	return s.postRepo.GetALLByUserID(ctx, userID)
}
