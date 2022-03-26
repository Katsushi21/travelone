package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
	"github.com/Katsushi21/traveling_alone/postgres"
)

func (r *mutationResolver) CreateMarker(ctx context.Context, input models.MarkerInput) (*models.Marker, error) {
	db := postgres.Connect()
	marker := &models.Marker{
		PostID: *input.PostID,
		Title:  input.Title,
		Lat:    input.Lat,
		Lng:    input.Lng,
	}
	db.Create(&marker)
	return marker, nil
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id int, input models.MarkerInput) (*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMarker(ctx context.Context, id int) (*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerResolver) Post(ctx context.Context, obj *models.Marker) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
