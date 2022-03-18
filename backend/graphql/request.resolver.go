package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *mutationResolver) CreateRequest(ctx context.Context, input models.RequestInput) (*models.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, id string, input models.RequestInput) (*models.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, id string) (*models.Request, error) {
	panic(fmt.Errorf("not implemented"))
}
