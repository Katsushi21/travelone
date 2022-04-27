package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input models.AccountInput) (*models.Account, error) {
	account := &models.Account{
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
		Create(&account).
		Error

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, id int, input models.AccountInput) (*models.Account, error) {
	account := &models.Account{
		ID: id,
	}
	err := r.DB.
		Debug().
		Model(&account).
		Updates(
			&models.Account{
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

	return account, nil
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, id int) (*models.Account, error) {
	account := &models.Account{
		ID: id,
	}
	err := r.DB.
		Debug().
		Clauses(clause.Returning{}).
		Delete(&account).
		Error

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.Account, error) {
	panic(fmt.Errorf("not implemented"))
}
