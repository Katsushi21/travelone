package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *queryResolver) Post(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	r.DB.Find(&posts)
	return posts, nil
}

func (r *queryResolver) User(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	r.DB.Find(&users)
	return users, nil
}

func (r *queryResolver) Comment(ctx context.Context) ([]*models.Comment, error) {
	var comments []*models.Comment
	r.DB.Find(&comments)
	return comments, nil
}

func (r *queryResolver) Marker(ctx context.Context) ([]*models.Marker, error) {
	var markers []*models.Marker
	r.DB.Find(&markers)
	return markers, nil
}

func (r *queryResolver) Request(ctx context.Context) ([]*models.Request, error) {
	var requests []*models.Request
	r.DB.Find(&requests)
	return requests, nil
}
