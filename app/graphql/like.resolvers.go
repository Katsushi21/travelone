package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateLike(ctx context.Context, input models.LikeInput) (*models.Like, error) {
	like := &models.Like{
		PostID: input.PostID,
		UserID: input.UserID,
	}
	err := r.DB.
		Debug().
		Create(&like).
		Error

	if err != nil {
		return nil, err
	}

	return like, nil
}

func (r *mutationResolver) DeleteLike(ctx context.Context, input models.LikeInput) (*models.Like, error) {
	like := &models.Like{
		PostID: input.PostID,
		UserID: input.UserID,
	}
	err := r.DB.
		Debug().
		Clauses(clause.Returning{}).
		Delete(&like).
		Error

	if err != nil {
		return nil, err
	}

	return like, nil
}
