package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input models.PostInput) (*models.Post, error) {
	post := models.Post{
		UserID: &input.UserID,
		Title:  &input.Title,
		Body:   &input.Body,
		Img:    &input.Img,
	}
	r.DB.Create(&post)

	return &post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input models.PostInput) (*models.Post, error) {
	post := models.Post{
		ID: id,
	}
	r.DB.First(&post)
	r.DB.Model(&post).Updates(
		&models.Post{
			Title: &input.Title,
			Body:  &input.Body,
			Img:   &input.Img,
		},
	)

	return &post, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id int) (*models.Post, error) {
	post := models.Post{
		ID: id,
	}
	r.DB.Clauses(clause.Returning{}).Delete(&post)

	return &post, nil
}
