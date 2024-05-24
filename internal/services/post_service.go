package services

import (
	"context"
	"errors"
	"github.com/Telekurysh/ozon_test/internal/models"
	"gorm.io/gorm"
)

type PostService struct {
	DB *gorm.DB
}

func (s *PostService) GetPosts(ctx context.Context) ([]models.Post, error) {
	var posts []models.Post
	result := s.DB.Find(&posts)
	return posts, result.Error
}

func (s *PostService) GetPost(ctx context.Context, id int) (*models.Post, error) {
	var post models.Post
	result := s.DB.First(&post, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &post, result.Error
}

func (s *PostService) CreatePost(ctx context.Context, title, content string, commentsDisabled bool) (*models.Post, error) {
	post := models.Post{Title: title, Content: content, CommentsDisabled: commentsDisabled}
	result := s.DB.Create(&post)
	return &post, result.Error
}
