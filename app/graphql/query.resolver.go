package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *queryResolver) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	r.DB.Debug().
		Joins("User").
		Joins("Marker").
		Joins("Comment").
		Find(&posts)
	return posts, nil
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	r.DB.Debug().
		Preload("Post").
		Joins("Like").
		Joins("Comment").
		Joins("Friend").
		Joins("Mute").
		Find(&users)
	return users, nil
}

func (r *queryResolver) GetAllComments(ctx context.Context) ([]*models.Comment, error) {
	var comments []*models.Comment
	r.DB.Debug().
		Joins("Post").
		Joins("User").
		Find(&comments)
	return comments, nil
}

func (r *queryResolver) GetAllMarkers(ctx context.Context) ([]*models.Marker, error) {
	var markers []*models.Marker
	r.DB.Debug().
		Joins("Post").
		Find(&markers)
	return markers, nil
}

func (r *queryResolver) GetAllRequests(ctx context.Context) ([]*models.Request, error) {
	var requests []*models.Request
	r.DB.Debug().
		Joins("User").
		Find(&requests)
	return requests, nil
}

func (r *queryResolver) GetAllLikes(ctx context.Context) ([]*models.Like, error) {
	var likes []*models.Like
	r.DB.Debug().
		Joins("Post").
		Joins("User").
		Find(&likes)
	return likes, nil
}

func (r *queryResolver) GetAllFriends(ctx context.Context) ([]*models.Friend, error) {
	var friends []*models.Friend
	r.DB.Debug().
		Joins("User").
		Find(&friends)
	return friends, nil
}

func (r *queryResolver) GetAllMutes(ctx context.Context) ([]*models.Mute, error) {
	var mutes []*models.Mute
	r.DB.Debug().
		Joins("User").
		Find(&mutes)
	return mutes, nil
}

func (r *queryResolver) GetPostByID(ctx context.Context, id int) (*models.Post, error) {
	var post *models.Post
	r.DB.Debug().
		Joins("User").
		Joins("Marker").
		Joins("Comment").
		First(&post, id)
	return post, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var user *models.User
	r.DB.Debug().
		Preload("Post").
		Preload("Comment").
		First(&user, id)
	return user, nil
}

func (r *queryResolver) GetCommentByID(ctx context.Context, id int) (*models.Comment, error) {
	var comment *models.Comment
	r.DB.Debug().
		Joins("Post").
		Joins("User").
		First(&comment, id)
	return comment, nil
}

func (r *queryResolver) GetMarkerByID(ctx context.Context, id int) (*models.Marker, error) {
	var marker *models.Marker
	r.DB.Debug().
		Joins("Post").
		First(&marker, id)
	return marker, nil
}

func (r *queryResolver) GetPostByUserID(ctx context.Context, userID int) ([]*models.Post, error) {
	var posts []*models.Post
	r.DB.Debug().
		Joins("User").
		Joins("Marker").
		Joins("Comment").
		Where(&models.Post{UserID: userID}).
		First(&posts)

	return posts, nil
}

func (r *queryResolver) GetLikeByUserID(ctx context.Context, userID int) ([]*models.Like, error) {
	var likes []*models.Like
	r.DB.Debug().
		Joins("Post").
		Joins("User").
		Where(&models.Like{UserID: userID}).
		First(&likes)

	return likes, nil
}

func (r *queryResolver) GetFriendByUserID(ctx context.Context, userID int) ([]*models.Friend, error) {
	var friends []*models.Friend
	r.DB.Debug().
		Joins("User").
		Where(&models.Friend{UserID: userID}).
		First(&friends)

	return friends, nil
}

func (r *queryResolver) GetMuteByUserID(ctx context.Context, userID int) ([]*models.Mute, error) {
	var mutes []*models.Mute
	r.DB.Debug().
		Joins("User").
		Where(&models.Mute{UserID: userID}).
		First(&mutes)

	return mutes, nil
}

func (r *queryResolver) GetRequestByUserID(ctx context.Context, userID int) ([]*models.Request, error) {
	var requests []*models.Request
	r.DB.Debug().
		Joins("User").
		Where(&models.Request{UserID: userID}).
		First(&requests)

	return requests, nil
}

func (r *queryResolver) GetSessionByUserID(ctx context.Context, userID int) (*models.Session, error) {
	var session *models.Session
	r.DB.Where(&models.Session{UserID: userID}).
		First(&session)

	return session, nil
}
