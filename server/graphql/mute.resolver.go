package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateMute(ctx context.Context, input models.MuteInput) (*models.Mute, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMute(ctx context.Context, input *models.MuteInput) (*models.Mute, error) {
	mute := models.Mute{
		UID:       &input.UID,
		TargetUID: &input.TargetUID,
	}
	r.DB.Clauses(clause.Returning{}).Delete(&mute)

	return &mute, nil
}
