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
	r.DB.Create(&user)

	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input models.UserInput) (*models.User, error) {
	user := models.User{
		ID: id,
	}
	r.DB.First(&user)
	r.DB.Model(&user).Updates(
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
	)

	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*models.User, error) {
	user := models.User{
		ID: id,
	}
	r.DB.Clauses(clause.Returning{}).Delete(&user)

	return &user, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}
