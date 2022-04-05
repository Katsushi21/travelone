package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateRequest(ctx context.Context, input models.RequestInput) (*models.Request, error) {
	request := &models.Request{
		RequestUID:   &input.RequestUID,
		RequestedUID: &input.RequestedUID,
		Status:       input.Status,
	}
	r.DB.Create(&request)
	return request, nil
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, input models.RequestInput) (*models.Request, error) {
	request := models.Request{
		RequestUID:   &input.RequestUID,
		RequestedUID: &input.RequestedUID,
	}
	r.DB.First(&request)
	r.DB.Model(&request).Updates( // Whereが必要か要検証
		&models.Request{
			RequestUID:   &input.RequestUID,
			RequestedUID: &input.RequestedUID,
			Status:       input.Status,
		},
	)

	return &request, nil
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, input models.RequestInput) (*models.Request, error) {
	request := models.Request{
		RequestUID:   &input.RequestUID,
		RequestedUID: &input.RequestedUID,
	}
	r.DB.Clauses(clause.Returning{}).Delete(&models.Request{})

	return &request, nil
}
