package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateFriend(ctx context.Context, input models.FriendInput) (*models.Friend, error) {
	friend := &models.Friend{
		UserID:   input.UserID,
		FriendID: input.FriendID,
	}
	r.DB.Create(&friend)

	return friend, nil
}

func (r *mutationResolver) DeleteFriend(ctx context.Context, input models.FriendInput) (*models.Friend, error) {
	friend := &models.Friend{
		UserID:   input.UserID,
		FriendID: input.FriendID,
	}
	r.DB.First(&friend)
	r.DB.Clauses(clause.Returning{}).Delete(&friend)

	return friend, nil
}
