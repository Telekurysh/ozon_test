package resolvers

import (
	"context"
	"github.com/Telekurysh/ozon_test/internal/models"
	"github.com/Telekurysh/ozon_test/internal/services"
)

type CommentResolver struct {
	CommentService *services.CommentService
}

func (r *CommentResolver) CreateComment(ctx context.Context, postId int, parentId *int, content string) (*models.Comment, error) {
	return r.CommentService.CreateComment(ctx, postId, parentId, content)
}
