package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *queryResolver) GetAllComments(ctx context.Context) ([]*models.Comment, error) {
	var comments []*models.Comment
	err := r.DB.
		Debug().
		Joins("Post").
		Joins("User").
		Find(&comments).
		Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *queryResolver) GetCommentByID(ctx context.Context, id int) (*models.Comment, error) {
	var comment *models.Comment
	err := r.DB.
		Debug().
		Joins("Post").
		Joins("User").
		First(&comment, id).
		Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *queryResolver) GetCommentByUserID(ctx context.Context, userID int) ([]*models.Comment, error) {
	var comment []*models.Comment
	err := r.DB.
		Debug().
		Joins("Post").
		Joins("User").
		Find(&comment, userID).
		Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *queryResolver) GetAllFriends(ctx context.Context) ([]*models.Friend, error) {
	var friends []*models.Friend
	err := r.DB.
		Debug().
		Preload("User").
		Find(&friends).
		Error
	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (r *queryResolver) GetFriendsByUserID(ctx context.Context, userID int) ([]*models.Friend, error) {
	var friends []*models.Friend
	err := r.DB.
		Debug().
		Joins("User").
		Where(&models.Friend{UserID: userID}).
		First(&friends).
		Error
	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (r *queryResolver) GetAllLikes(ctx context.Context) ([]*models.Like, error) {
	var likes []*models.Like
	err := r.DB.
		Debug().
		Preload("Post").
		Preload("User").
		Find(&likes).
		Error
	if err != nil {
		return nil, err
	}

	return likes, nil
}

func (r *queryResolver) GetLikesByUserID(ctx context.Context, userID int) ([]*models.Like, error) {
	var likes []*models.Like
	err := r.DB.
		Debug().
		Joins("Post").
		Joins("User").
		Where(&models.Like{UserID: userID}).
		First(&likes).
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
		Preload("Post").
		Find(&markers).
		Error
	if err != nil {
		return nil, err
	}

	return markers, nil
}

func (r *queryResolver) GetMarkerByID(ctx context.Context, id int) (*models.Marker, error) {
	var marker *models.Marker
	err := r.DB.
		Debug().
		Joins("Post").
		First(&marker, id).
		Error
	if err != nil {
		return nil, err
	}

	return marker, nil
}

func (r *queryResolver) GetAllMutes(ctx context.Context) ([]*models.Mute, error) {
	var mutes []*models.Mute
	err := r.DB.
		Debug().
		Preload("User").
		Find(&mutes).
		Error
	if err != nil {
		return nil, err
	}

	return mutes, nil
}

func (r *queryResolver) GetMutesByUserID(ctx context.Context, userID int) ([]*models.Mute, error) {
	var mutes []*models.Mute
	err := r.DB.
		Debug().
		Joins("User").
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

func (r *queryResolver) GetPostByID(ctx context.Context, id int) (*models.Post, error) {
	var post *models.Post
	err := r.DB.
		Debug().
		Preload("User").
		Preload("Marker").
		Preload("Comment").
		First(&post, id).
		Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *queryResolver) GetPostsByUserID(ctx context.Context, userID int) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.DB.
		Debug().
		Joins("User", r.DB.Where(&models.User{ID: userID})).
		Joins("Marker").
		Preload("Comment").
		Where(&models.Post{UserID: userID}).
		Find(&posts).
		Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *queryResolver) GetAllRequests(ctx context.Context) ([]*models.Request, error) {
	var requests []*models.Request
	err := r.DB.
		Debug().
		Preload("User").
		Find(&requests).
		Error
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *queryResolver) GetRequestsByUserID(ctx context.Context, userID int) ([]*models.Request, error) {
	var requests []*models.Request
	err := r.DB.
		Debug().
		Joins("User").
		Where(&models.Request{UserID: userID}).
		First(&requests).
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
		Preload("Post").
		Preload("Comment").
		First(&user, id).
		Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
