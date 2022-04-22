package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateFriend(ctx context.Context, input models.FriendInput) (*models.Friend, error) {
	friend := &models.Friend{
		OwnID:  input.OwnID,
		UserID: input.UserID,
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
		OwnID:  input.OwnID,
		UserID: input.UserID,
	}
	r.DB.Debug().First(&friend)
	err := r.DB.Debug().Clauses(clause.Returning{}).Delete(&friend).Error
	if err != nil {
		return nil, err
	}

	return friend, nil
}
