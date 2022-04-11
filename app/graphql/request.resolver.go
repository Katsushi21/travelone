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
	r.DB.Create(&request)

	return request, nil
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, input models.RequestInput) (*models.Request, error) {
	request := &models.Request{
		UserID:       input.UserID,
		TargetUserID: input.TargetUserID,
	}
	r.DB.First(&request)
	r.DB.Model(&request).Update(
		"status",
		&input.Status,
	)

	return request, nil
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, input models.RequestInput) (*models.Request, error) {
	request := &models.Request{
		UserID:       input.UserID,
		TargetUserID: input.TargetUserID,
	}
	r.DB.Clauses(clause.Returning{}).Delete(&request)

	return request, nil
}
