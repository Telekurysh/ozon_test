package services

import (
	"context"
	"github.com/Telekurysh/ozon_test/internal/models"
	"gorm.io/gorm"
)

type CommentService struct {
	DB *gorm.DB
}

func (s *CommentService) GetCommentsByPost(ctx context.Context, postId int, limit, offset int) ([]models.Comment, error) {
	var comments []models.Comment
	result := s.DB.Where("post_id = ? AND parent_id IS NULL", postId).Limit(limit).Offset(offset).Find(&comments)
	return comments, result.Error
}

func (s *CommentService) CreateComment(ctx context.Context, postId int, parentId *int, content string) (*models.Comment, error) {
	comment := models.Comment{PostID: postId, ParentID: parentId, Content: content}
	result := s.DB.Create(&comment)
	return &comment, result.Error
}
