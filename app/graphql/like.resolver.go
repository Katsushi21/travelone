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
	r.DB.Create(&like)

	return like, nil
}

func (r *mutationResolver) DeleteLike(ctx context.Context, input models.LikeInput) (*models.Like, error) {
	like := &models.Like{
		PostID: input.PostID,
		UserID: input.UserID,
	}
	r.DB.Clauses(clause.Returning{}).Delete(&like)

	return like, nil
}
