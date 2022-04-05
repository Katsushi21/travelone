package graphql

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *queryResolver) Post(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	r.DB.Preload("User").Preload("Marker").Preload("Comment").Find(&posts)
	return posts, nil
}

func (r *queryResolver) User(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	r.DB.Preload("Post").Preload("Comment").Find(&users)
	return users, nil
}

func (r *queryResolver) Comment(ctx context.Context) ([]*models.Comment, error) {
	var comments []*models.Comment
	r.DB.Preload("Post").Preload("User").Find(&comments)
	return comments, nil
}

func (r *queryResolver) Marker(ctx context.Context) ([]*models.Marker, error) {
	var markers []*models.Marker
	r.DB.Preload("Post").Find(&markers)
	return markers, nil
}

func (r *queryResolver) Request(ctx context.Context) ([]*models.Request, error) {
	var requests []*models.Request
	r.DB.Preload("User").Find(&requests)
	return requests, nil
}

func (r *queryResolver) Like(ctx context.Context) ([]*models.Like, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Friend(ctx context.Context) ([]*models.Friend, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Mute(ctx context.Context) ([]*models.Mute, error) {
	panic(fmt.Errorf("not implemented"))
}
