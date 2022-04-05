package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *mutationResolver) CreateLike(ctx context.Context, input models.LikeInput) (*models.Like, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteLike(ctx context.Context, input models.LikeInput) (*models.Like, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeResolver) PostID(ctx context.Context, obj *models.Like) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeResolver) UID(ctx context.Context, obj *models.Like) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}
