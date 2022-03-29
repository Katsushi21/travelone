package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
	"github.com/Katsushi21/traveling_alone/postgres"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	db := postgres.Connect()
	user := &models.User{
		Email:        &input.Email,
		Password:     &input.Password,
		Type:         &input.Type,
		Name:         &input.Name,
		Age:          &input.Age,
		Gender:       &input.Gender,
		Avatar:       &input.Avatar,
		Introduction: &input.Introduction,
	}
	db.Create(&user)
	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input models.UserInput) (*models.User, error) {
	db := postgres.Connect()
	user := models.User{
		ID: id,
	}
	db.First(&user)
	db.Model(&user).Where("id = ?", id).Updates( // Whereが必要か要検証
		&models.User{
			Email:        &input.Email,
			Password:     &input.Password,
			Type:         &input.Type,
			Name:         &input.Name,
			Age:          &input.Age,
			Gender:       &input.Gender,
			Avatar:       &input.Avatar,
			Introduction: &input.Introduction,
		},
	)

	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*models.User, error) {
	db := postgres.Connect()
	user := models.User{
		ID: id,
	}
	db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&models.User{})

	return &user, nil
}

func (r *mutationResolver) UpdateSession(ctx context.Context, id int, input models.SessionInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateFriend(ctx context.Context, id int, input models.FriendInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMute(ctx context.Context, id int, input *models.MuteInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}
