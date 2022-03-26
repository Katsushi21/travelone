package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
	"github.com/Katsushi21/traveling_alone/postgres"
)

func (r *queryResolver) Post(ctx context.Context) ([]*models.Post, error) {
	db := postgres.Connect()
	var posts []*models.Post
	db.Find(&posts)
	return posts, nil
}

func (r *queryResolver) User(ctx context.Context) ([]*models.User, error) {
	db := postgres.Connect()
	var users []*models.User
	db.Find(&users)
	return users, nil
}

func (r *queryResolver) Comment(ctx context.Context) ([]*models.Comment, error) {
	db := postgres.Connect()
	var comments []*models.Comment
	db.Find(&comments)
	return comments, nil
}

func (r *queryResolver) Marker(ctx context.Context) ([]*models.Marker, error) {
	db := postgres.Connect()
	var markers []*models.Marker
	db.Find(&markers)
	return markers, nil
}

func (r *queryResolver) Request(ctx context.Context) ([]*models.Request, error) {
	db := postgres.Connect()
	var requests []*models.Request
	db.Find(&requests)
	return requests, nil
}
