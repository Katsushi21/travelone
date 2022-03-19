package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input models.PostInput) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input models.PostInput) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePost(ctx context.Context, id int) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateLiked(ctx context.Context, id int, input models.LikedInput) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
