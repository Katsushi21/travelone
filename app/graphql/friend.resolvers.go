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
	err := r.DB.
		Debug().
		Create(&friend).
		Error
	if err != nil {
		return nil, err
	}

	return friend, nil
}

func (r *mutationResolver) DeleteFriend(ctx context.Context, input models.FriendInput) (*models.Friend, error) {
	friend := &models.Friend{
		UserID:   input.UserID,
		FriendID: input.FriendID,
	}
	r.DB.Debug().First(&friend)
	err := r.DB.Debug().Clauses(clause.Returning{}).Delete(&friend).Error
	if err != nil {
		return nil, err
	}

	return friend, nil
}
