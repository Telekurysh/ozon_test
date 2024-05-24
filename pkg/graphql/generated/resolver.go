package generated

import (
	"context"

	"github.com/Telekurysh/ozon_test/internal/models"
)

type Resolver struct{}

func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	// Реализация получения всех постов
}

func (r *queryResolver) Post(ctx context.Context, id int) (*models.Post, error) {
	// Реализация получения поста по ID
}

func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, commentsDisabled bool) (*models.Post, error) {
	// Реализация создания поста
}

func (r *mutationResolver) CreateComment(ctx context.Context, postId int, parentId *int, content string) (*models.Comment, error) {
	// Реализация создания комментария
}

func (r *subscriptionResolver) CommentAdded(ctx context.Context, postId int) (<-chan *models.Comment, error) {
	// Реализация подписки на комментарии
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
