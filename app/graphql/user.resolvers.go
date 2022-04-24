package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	user := &models.User{
		Email:        input.Email,
		Password:     input.Password,
		Type:         input.Type,
		Name:         input.Name,
		Age:          input.Age,
		Gender:       input.Gender,
		Avatar:       input.Avatar,
		Introduction: input.Introduction,
	}
	err := r.DB.
		Debug().
		Create(&user).
		Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input models.UserInput) (*models.User, error) {
	user := &models.User{
		ID: id,
	}
	err := r.DB.
		Debug().
		Model(&user).
		Updates(
			&models.User{
				Email:        input.Email,
				Password:     input.Password,
				Type:         input.Type,
				Name:         input.Name,
				Age:          input.Age,
				Gender:       input.Gender,
				Avatar:       input.Avatar,
				Introduction: input.Introduction,
			},
		).
		Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{
		ID: id,
	}
	err := r.DB.
		Debug().
		Clauses(clause.Returning{}).
		Delete(&user).
		Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}
