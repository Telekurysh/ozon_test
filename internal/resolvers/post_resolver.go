package resolvers

import (
	"context"
	"github.com/Telekurysh/ozon_test/internal/models"
	"github.com/Telekurysh/ozon_test/internal/services"
)

type PostResolver struct {
	PostService *services.PostService
}

func (r *PostResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	posts, err := r.PostService.GetPosts(ctx)
	if err != nil {
		return nil, err
	}
	var result []*models.Post
	for _, post := range posts {
		result = append(result, &post)
	}
	return result, nil
}

func (r *PostResolver) Post(ctx context.Context, id int) (*models.Post, error) {
	return r.PostService.GetPost(ctx, id)
}
