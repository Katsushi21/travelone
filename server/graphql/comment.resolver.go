package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateComment(ctx context.Context, input models.CommentInput) (*models.Comment, error) {
	comment := &models.Comment{
		PostID: &input.PostID,
		UID:    &input.UID,
		Body:   input.Body,
	}
	r.DB.Create(&comment)
	return comment, nil
}
func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (*models.Comment, error) {
	comment := models.Comment{
		ID: id,
	}
	r.DB.Clauses(clause.Returning{}).Delete(&models.Comment{})

	return &comment, nil
}
