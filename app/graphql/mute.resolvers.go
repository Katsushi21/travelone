package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateMute(ctx context.Context, input models.MuteInput) (*models.Mute, error) {
	mute := &models.Mute{
		AccountID: input.AccountID,
		MuteID:    input.MuteID,
	}
	err := r.DB.Debug().Create(&mute).Error
	if err != nil {
		return nil, err
	}

	return mute, nil
}

func (r *mutationResolver) DeleteMute(ctx context.Context, input *models.MuteInput) (*models.Mute, error) {
	mute := &models.Mute{
		AccountID: input.AccountID,
		MuteID:    input.MuteID,
	}
	err := r.DB.
		Debug().
		Clauses(clause.Returning{}).
		Delete(&mute).
		Error

	if err != nil {
		return nil, err
	}

	return mute, nil
}