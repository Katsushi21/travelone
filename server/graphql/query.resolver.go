package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *queryResolver) Post(ctx context.Context, id int) (*models.Post, error) {
	var posts []*models.Post
	r.DB.Preload("User").Preload("Marker").Preload("Comment").Find(&posts)
	return posts, nil
}

func (r *queryResolver) AllPosts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	r.DB.Preload("User").Preload("Marker").Preload("Comment").Find(&posts)
	return posts, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	var users []*models.User
	r.DB.Preload("Post").Preload("Comment").Find(&users)
	return users, nil
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	r.DB.Preload("Post").Preload("Comment").Find(&users)
	return users, nil
}

func (r *queryResolver) Comment(ctx context.Context, id int) (*models.Comment, error) {
	var comments []*models.Comment
	r.DB.Preload("Post").Preload("User").Find(&comments)
	return comments, nil
}

func (r *queryResolver) AllComments(ctx context.Context) ([]*models.Comment, error) {
	var comments []*models.Comment
	r.DB.Preload("Post").Preload("User").Find(&comments)
	return comments, nil
}

func (r *queryResolver) Marker(ctx context.Context, id int) (*models.Marker, error) {
	var markers []*models.Marker
	r.DB.Preload("Post").Find(&markers)
	return markers, nil
}

func (r *queryResolver) AllMarkers(ctx context.Context) ([]*models.Marker, error) {
	var markers []*models.Marker
	r.DB.Preload("Post").Find(&markers)
	return markers, nil
}

func (r *queryResolver) Request(ctx context.Context, id int) (*models.Request, error) {
	var requests []*models.Request
	r.DB.Preload("User").Find(&requests)
	return requests, nil
}

func (r *queryResolver) AllRequests(ctx context.Context) ([]*models.Request, error) {
	var requests []*models.Request
	r.DB.Preload("User").Find(&requests)
	return requests, nil
}

func (r *queryResolver) Like(ctx context.Context, id int) (*models.Like, error) {
	var likes []*models.Like
	r.DB.Preload("Post").Preload("User").Find(&likes)
	return likes, nil
}

func (r *queryResolver) AllLikes(ctx context.Context) ([]*models.Like, error) {
	var likes []*models.Like
	r.DB.Preload("Post").Preload("User").Find(&likes)
	return likes, nil
}

func (r *queryResolver) Friend(ctx context.Context, id int) (*models.Friend, error) {
	var friends []*models.Friend
	r.DB.Preload("User").Find(&friends)
	return friends, nil
}

func (r *queryResolver) AllFriends(ctx context.Context) ([]*models.Friend, error) {
	var friends []*models.Friend
	r.DB.Preload("User").Find(&friends)
	return friends, nil
}

func (r *queryResolver) Mute(ctx context.Context, id int) (*models.Mute, error) {
	var mutes []*models.Mute
	r.DB.Preload("User").Find(&mutes)
	return mutes, nil
}

func (r *queryResolver) AllMutes(ctx context.Context) ([]*models.Mute, error) {
	var mutes []*models.Mute
	r.DB.Preload("User").Find(&mutes)
	return mutes, nil
}

func (r *queryResolver) Session(ctx context.Context, id int) (*models.Session, error) {
	var session *models.Session
	r.DB.Find(&session)
	return session, nil
}
