package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateComment(ctx context.Context, input models.CommentInput) (*models.Comment, error) {
	comment := &models.Comment{
		PostID: input.PostID,
		UserID: input.UserID,
		Body:   input.Body,
	}
	err := r.DB.Debug().Create(&comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (*models.Comment, error) {
	comment := &models.Comment{
		ID: id,
	}
	err := r.DB.Debug().Clauses(clause.Returning{}).Delete(&comment).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}
