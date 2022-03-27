package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
	"github.com/Katsushi21/traveling_alone/postgres"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateRequest(ctx context.Context, input models.RequestInput) (*models.Request, error) {
	db := postgres.Connect()
	request := &models.Request{
		RequestUID:   input.RequestUID,
		RequestedUID: input.RequestedUID,
		Status:       input.Status,
	}
	db.Create(&request)
	return request, nil
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, id int, input models.RequestInput) (*models.Request, error) {
	db := postgres.Connect()
	request := models.Request{
		ID: id,
	}
	db.First(&request)
	db.Model(&request).Where("id = ?", id).Updates( // Whereが必要か要検証
		&models.Request{
			RequestUID:   input.RequestUID,
			RequestedUID: input.RequestedUID,
			Status:       input.Status,
		},
	)

	return &request, nil
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, id int) (*models.Request, error) {
	db := postgres.Connect()
	request := models.Request{
		ID: id,
	}
	db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&models.Request{})

	return &request, nil
}

func (r *requestResolver) RequestUser(ctx context.Context, obj *models.Request) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestResolver) RequestedUser(ctx context.Context, obj *models.Request) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}
