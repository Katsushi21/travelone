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

func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) Liked(ctx context.Context, obj *models.Post) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) Marker(ctx context.Context, obj *models.Post) (*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) Comment(ctx context.Context, obj *models.Post) ([]*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Post(ctx context.Context, obj *models.User) ([]*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
