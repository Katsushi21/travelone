package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateMarker(ctx context.Context, input models.MarkerInput) (*models.Marker, error) {
	marker := &models.Marker{
		PostID: input.PostID,
		Title:  input.Title,
		Lat:    input.Lat,
		Lng:    input.Lng,
	}
	r.DB.Create(&marker)
	return marker, nil
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id int, input models.MarkerInput) (*models.Marker, error) {
	marker := &models.Marker{
		ID: id,
	}
	r.DB.First(&marker)
	r.DB.Model(&marker).Updates(
		&models.Marker{
			PostID: input.PostID,
			Title:  input.Title,
			Lat:    input.Lat,
			Lng:    input.Lng,
		},
	)

	return marker, nil
}

func (r *mutationResolver) DeleteMarker(ctx context.Context, id int) (*models.Marker, error) {
	marker := &models.Marker{
		ID: id,
	}
	r.DB.Clauses(clause.Returning{}).Delete(&marker)

	return marker, nil
}
