// Query returns QueryResolver implementation.
package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *queryResolver) Post(ctx context.Context) ([]*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comment(ctx context.Context) ([]*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Marker(ctx context.Context) ([]*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Request(ctx context.Context) ([]*models.Request, error) {
	panic(fmt.Errorf("not implemented"))
}
