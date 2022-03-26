package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
	"github.com/Katsushi21/traveling_alone/postgres"
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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, id int) (*models.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestResolver) RequestUser(ctx context.Context, obj *models.Request) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestResolver) RequestedUser(ctx context.Context, obj *models.Request) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}
