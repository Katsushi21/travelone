package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"github.com/Katsushi21/traveling_alone/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input models.PostInput) (*models.Post, error) {
	db := postgres.Connect()
	post := models.Post{
		UID:   input.UID,
		Title: input.Title,
		Body:  input.Body,
		Img:   input.Img,
	}
	db.Create(&post)

	return &post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input models.PostInput) (*models.Post, error) {
	db := postgres.Connect()
	post := models.Post{
		ID: id,
	}
	db.First(&post)
	db.Model(&post).Where("id = ?", id).Updates( // Whereが必要か要検証
		&models.Post{
			Title: input.Title,
			Body:  input.Body,
			Img:   input.Img,
		},
	)

	return &post, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id int) (*models.Post, error) {
	db := postgres.Connect()
	post := models.Post{
		ID: id,
	}
	db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&models.Post{})

	return &post, nil
}

func (r *mutationResolver) UpdateLiked(ctx context.Context, id int, input models.LikedInput) (*models.Post, error) {
	db := postgres.Connect()
	post := models.Post{
		ID: id,
	}
	db.First(&post)
	db.Model(&post).Where("id = ?", id).Update( // Whereが必要か要検証
		"liked", gorm.Expr("array_append(?, ?)", "liked", input.UID),
	)

	return &post, nil
}
