package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateRequest(ctx context.Context, input models.RequestInput) (*models.Request, error) {
	request := &models.Request{
		UserID:       input.UserID,
		TargetUserID: input.TargetUserID,
		Status:       input.Status,
	}
	err := r.DB.Debug().Create(&request).Error
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, input models.RequestInput) (*models.Request, error) {
	request := &models.Request{
		UserID:       input.UserID,
		TargetUserID: input.TargetUserID,
	}
	r.DB.Debug().First(&request)
	err := r.DB.Debug().Model(&request).
		Update(
			"status",
			&input.Status,
		).
		Error
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, input models.RequestInput) (*models.Request, error) {
	request := &models.Request{
		UserID:       input.UserID,
		TargetUserID: input.TargetUserID,
	}
	err := r.DB.Debug().Clauses(clause.Returning{}).
		Delete(&request).
		Error
	if err != nil {
		return nil, err
	}

	return request, nil
}
