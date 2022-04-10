package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *queryResolver) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	r.DB.Preload("User").
		Preload("Marker").
		Preload("Comment").
		Find(&posts)
	return posts, nil
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	r.DB.Preload("Post").
		Preload("Comment").
		Find(&users)
	return users, nil
}

func (r *queryResolver) GetAllComments(ctx context.Context) ([]*models.Comment, error) {
	var comments []*models.Comment
	r.DB.Preload("Post").
		Preload("User").
		Find(&comments)
	return comments, nil
}

func (r *queryResolver) GetAllMarkers(ctx context.Context) ([]*models.Marker, error) {
	var markers []*models.Marker
	r.DB.Preload("Post").Find(&markers)
	return markers, nil
}

func (r *queryResolver) GetAllRequests(ctx context.Context) ([]*models.Request, error) {
	var requests []*models.Request
	r.DB.Preload("User").Find(&requests)
	return requests, nil
}

func (r *queryResolver) GetAllLikes(ctx context.Context) ([]*models.Like, error) {
	var likes []*models.Like
	r.DB.Preload("Post").
		Preload("User").
		Find(&likes)
	return likes, nil
}

func (r *queryResolver) GetAllFriends(ctx context.Context) ([]*models.Friend, error) {
	var friends []*models.Friend
	r.DB.Preload("User").Find(&friends)
	return friends, nil
}

func (r *queryResolver) GetAllMutes(ctx context.Context) ([]*models.Mute, error) {
	var mutes []*models.Mute
	r.DB.Preload("User").Find(&mutes)
	return mutes, nil
}

func (r *queryResolver) GetPost(ctx context.Context, id int) (*models.Post, error) {
	var post *models.Post
	r.DB.Preload("User").
		Preload("Marker").
		Preload("Comment").
		First(&post, id)
	return post, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id int) (*models.User, error) {
	var user *models.User
	r.DB.Preload("Post").
		Preload("Comment").
		First(&user, id)
	return user, nil
}

func (r *queryResolver) GetComment(ctx context.Context, id int) (*models.Comment, error) {
	var comment *models.Comment
	r.DB.Preload("Post").
		Preload("User").
		First(&comment, id)
	return comment, nil
}

func (r *queryResolver) GetMarker(ctx context.Context, id int) (*models.Marker, error) {
	var marker *models.Marker
	r.DB.Preload("Post").
		First(&marker, id)
	return marker, nil
}

func (r *queryResolver) GetPostPerUser(ctx context.Context, userID int) ([]*models.Post, error) {
	var posts []*models.Post
	r.DB.Preload("User").
		Preload("Marker").
		Preload("Comment").
		Where(&models.Post{UserID: userID}).
		First(&posts)

	return posts, nil
}

func (r *queryResolver) GetLikePerUser(ctx context.Context, userID int) ([]*models.Like, error) {
	var likes []*models.Like
	r.DB.Preload("Post").
		Preload("User").
		Where(&models.Like{UserID: userID}).
		First(&likes)

	return likes, nil
}

func (r *queryResolver) GetFriendPerUser(ctx context.Context, userID int) ([]*models.Friend, error) {
	var friends []*models.Friend
	r.DB.Preload("User").
		Where(&models.Friend{UserID: userID}).
		First(&friends)

	return friends, nil
}

func (r *queryResolver) GetMutePerUser(ctx context.Context, userID int) ([]*models.Mute, error) {
	var mutes []*models.Mute
	r.DB.Preload("User").
		Where(&models.Mute{UserID: userID}).
		First(&mutes)

	return mutes, nil
}

func (r *queryResolver) GetRequestPerUser(ctx context.Context, userID int) ([]*models.Request, error) {
	var requests []*models.Request
	r.DB.Preload("User").
		Where(&models.Request{UserID: userID}).
		First(&requests)

	return requests, nil
}

func (r *queryResolver) GetSessionPerUser(ctx context.Context, userID int) (*models.Session, error) {
	var session *models.Session
	r.DB.Where(&models.Session{UserID: userID}).
		First(&session)

	return session, nil
}
