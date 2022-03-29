package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"github.com/Katsushi21/traveling_alone/postgres"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateMarker(ctx context.Context, input models.MarkerInput) (*models.Marker, error) {
	db := postgres.Connect()
	marker := &models.Marker{
		PostID: input.PostID,
		Title:  input.Title,
		Lat:    input.Lat,
		Lng:    input.Lng,
	}
	db.Create(&marker)
	return marker, nil
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id int, input models.MarkerInput) (*models.Marker, error) {
	db := postgres.Connect()
	marker := models.Marker{
		ID: id,
	}
	db.First(&marker)
	db.Model(&marker).Where("id = ?", id).Updates( // Whereが必要か要検証
		&models.Marker{
			PostID: input.PostID,
			Title:  input.Title,
			Lat:    input.Lat,
			Lng:    input.Lng,
		},
	)

	return &marker, nil
}

func (r *mutationResolver) DeleteMarker(ctx context.Context, id int) (*models.Marker, error) {
	db := postgres.Connect()
	marker := models.Marker{
		ID: id,
	}
	db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&models.Marker{})

	return &marker, nil
}
