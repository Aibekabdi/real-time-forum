package service

import (
	"context"
	"errors"
	"fmt"
	"forum/internal/models"
	"forum/internal/repository"
	"strings"
	"unicode"
)

type PostService struct {
	postRepo repository.Post
	tagRepo  repository.Tag
}

func newPostService(postRepo repository.Post, tagRepo repository.Tag) *PostService {
	return &PostService{postRepo: postRepo, tagRepo: tagRepo}
}

func (s *PostService) Create(ctx context.Context, post models.Post) (uint, error) {
	if err := postValidation(&post); err != nil {
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

func postValidation(post *models.Post) error {
	if post.Title == "" || len(post.Title) > 60 {
		return errors.New("invalid title")
	}
	text := strings.TrimFunc(post.Text, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	if text == "" || len(text) > 500 {
		return errors.New("invalid text")
	}
	if len(post.Tags) == 0 {
		return errors.New("no chosen or created tag")
	}
	filteredTags, err := tagsValidation(post.Tags)
	if err != nil {
		return err
	}
	post.Tags = filteredTags
	return nil
}

func tagsValidation(tags []models.Tags) ([]models.Tags, error) {
	var filteredTags []models.Tags
	for _, tag := range tags {
		if tag.Name == "" || len(tag.Name) > 60 {
			continue
		}
		filteredTags = append(filteredTags, tag)
	}
	if len(filteredTags) == 0 {
		return nil, errors.New("invalid tags")
	}
	return filteredTags, nil
}
