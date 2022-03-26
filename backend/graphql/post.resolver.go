package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
	"github.com/Katsushi21/traveling_alone/postgres"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input models.PostInput) (*models.Post, error) {
	db := postgres.Connect()
	post := &models.Post{
		UID:   *input.UID,
		Title: input.Title,
		Body:  input.Body,
		Img:   input.Img,
	}
	db.Create(&post)
	return post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input models.PostInput) (*models.Post, error) {
	db := postgres.Connect()
	post := models.Post{ID: *input.ID}
	db.First(&post)
	post = input.Text
	db.Model(&models.Post{}).Update(&post)

	return &post, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id int) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateLiked(ctx context.Context, id int, input models.LikedInput) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) Liked(ctx context.Context, obj *models.Post) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) Marker(ctx context.Context, obj *models.Post) (*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) Comment(ctx context.Context, obj *models.Post) ([]*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}
