package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"github.com/Katsushi21/traveling_alone/postgres"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateComment(ctx context.Context, input models.CommentInput) (*models.Comment, error) {
	db := postgres.Connect()
	comment := &models.Comment{
		PostID: input.PostID,
		UID:    input.UID,
		Body:   input.Body,
	}
	db.Create(&comment)
	return comment, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (*models.Comment, error) {
	db := postgres.Connect()
	comment := models.Comment{
		ID: id,
	}
	db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&models.Comment{})

	return &comment, nil
}
