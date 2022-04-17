package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input models.PostInput) (*models.Post, error) {
	post := &models.Post{
		UserID: input.UserID,
		Title:  input.Title,
		Body:   input.Body,
		Img:    input.Img,
	}
	err := r.DB.Debug().Create(&post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input models.PostInput) (*models.Post, error) {
	post := &models.Post{
		ID: id,
	}
	r.DB.Debug().First(&post)
	err := r.DB.Debug().Model(&post).
		Updates(
			&models.Post{
				Title: input.Title,
				Body:  input.Body,
				Img:   input.Img,
			},
		).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id int) (*models.Post, error) {
	post := &models.Post{
		ID: id,
	}
	err := r.DB.Debug().Clauses(clause.Returning{}).Delete(&post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}
