package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *mutationResolver) CreateMarker(ctx context.Context, input models.MarkerInput) (*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id string, input models.MarkerInput) (*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMarker(ctx context.Context, id string) (*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}
