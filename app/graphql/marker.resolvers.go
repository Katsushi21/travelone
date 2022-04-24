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
	err := r.DB.
		Debug().
		Create(&marker).
		Error

	if err != nil {
		return nil, err
	}

	return marker, nil
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id int, input models.MarkerInput) (*models.Marker, error) {
	marker := &models.Marker{
		ID: id,
	}
	err := r.DB.
		Debug().
		Model(&marker).
		Updates(
			&models.Marker{
				PostID: input.PostID,
				Title:  input.Title,
				Lat:    input.Lat,
				Lng:    input.Lng,
			},
		).
		Error

	if err != nil {
		return nil, err
	}

	return marker, nil
}

func (r *mutationResolver) DeleteMarker(ctx context.Context, id int) (*models.Marker, error) {
	marker := &models.Marker{
		ID: id,
	}
	err := r.DB.
		Debug().
		Clauses(clause.Returning{}).
		Delete(&marker).
		Error

	if err != nil {
		return nil, err
	}

	return marker, nil
}
