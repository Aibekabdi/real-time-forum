package utils

import (
	"errors"
	"forum/internal/models"
	"strings"
	"unicode"
)

func PostValidation(post *models.Post) error {
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
