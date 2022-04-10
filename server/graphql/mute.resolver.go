package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateMute(ctx context.Context, input models.MuteInput) (*models.Mute, error) {
	mute := &models.Mute{
		UserID: input.UserID,
		MuteID: input.MuteID,
	}
	r.DB.Create(&mute)

	return mute, nil
}

func (r *mutationResolver) DeleteMute(ctx context.Context, input *models.MuteInput) (*models.Mute, error) {
	mute := &models.Mute{
		UserID: input.UserID,
		MuteID: input.MuteID,
	}
	r.DB.Clauses(clause.Returning{}).Delete(&mute)

	return mute, nil
}
