package graphql

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *queryResolver) GetCommentByUserID(ctx context.Context, userID int) ([]*models.Comment, error) {
	var comment []*models.Comment
	err := r.DB.
		Debug().
		Joins("Post", r.DB.Where(&models.Post{UserID: userID})).
		Joins("User", r.DB.Where(&models.User{ID: userID})).
		Find(&comment, userID).
		Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *queryResolver) GetFriendsByUserID(ctx context.Context, userID int) ([]*models.Friend, error) {
	var friends []*models.Friend
	err := r.DB.
		Debug().
		Joins("User").
		Where(&models.Friend{UserID: userID}).
		Find(&friends).
		Error
	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (r *queryResolver) GetLikesByUserID(ctx context.Context, userID int) ([]*models.Like, error) {
	var likes []*models.Like
	err := r.DB.
		Debug().
		Joins("Post").
		Joins("User", r.DB.Where(&models.User{ID: userID})).
		Where(&models.Like{UserID: userID}).
		Find(&likes).
		Error
	if err != nil {
		return nil, err
	}

	return likes, nil
}

func (r *queryResolver) GetAllMarkers(ctx context.Context) ([]*models.Marker, error) {
	var markers []*models.Marker
	err := r.DB.
		Debug().
		Joins("Post").
		Find(&markers).
		Error
	if err != nil {
		return nil, err
	}

	return markers, nil
}

func (r *queryResolver) GetMutesByUserID(ctx context.Context, userID int) ([]*models.Mute, error) {
	var mutes []*models.Mute
	err := r.DB.
		Debug().
		Joins("User", r.DB.Where(&models.User{ID: userID})).
		Where(&models.Mute{UserID: userID}).
		First(&mutes).
		Error
	if err != nil {
		return nil, err
	}

	return mutes, nil
}

func (r *queryResolver) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.DB.
		Debug().
		Joins("User", r.DB.Where(&models.User{Type: "active"})).
		Joins("Marker").
		Preload("Comment").
		Find(&posts).
		Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *queryResolver) GetPostsByUserID(ctx context.Context, userID int) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.DB.
		Debug().
		Joins("User", r.DB.Where(&models.User{ID: userID})).
		Joins("Comment", r.DB.Where(&models.Comment{PostID: userID})).
		Preload("Marker").
		Where(&models.Post{UserID: userID}).
		Find(&posts).
		Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *queryResolver) GetRequestsByUserID(ctx context.Context, userID int) ([]*models.Request, error) {
	var requests []*models.Request
	err := r.DB.
		Debug().
		Joins("User", r.DB.Where(&models.User{ID: userID})).
		Where(&models.Request{UserID: userID}).
		Find(&requests).
		Error
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *queryResolver) GetRequestsByTargetID(ctx context.Context, targetID int) ([]*models.Request, error) {
	var requests []*models.Request
	err := r.DB.
		Debug().
		Joins("User", r.DB.Where(&models.User{ID: targetID})).
		Where(&models.Request{TargetUserID: targetID}).
		Find(&requests).
		Error
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *queryResolver) GetSessionByUserID(ctx context.Context, userID int) (*models.Session, error) {
	var session *models.Session
	err := r.DB.
		Where(&models.Session{UserID: userID}).
		First(&session).
		Error
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := r.DB.
		Debug().
		Preload("Post").
		Preload("Like").
		Preload("Comment").
		Preload("Friend").
		Preload("Mute").
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var user *models.User
	err := r.DB.
		Debug().
		Joins("Post", r.DB.Where(&models.Post{UserID: id})).
		Joins("Comment", r.DB.Where(&models.Comment{UserID: id})).
		Find(&user, id).
		Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
