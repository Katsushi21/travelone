package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *mutationResolver) CreateComment(ctx context.Context, input models.CommentInput) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id int, input models.CommentInput) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}
