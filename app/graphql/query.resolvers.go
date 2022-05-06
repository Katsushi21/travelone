package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *queryResolver) GetAccountByID(ctx context.Context, id int) (*models.Account, error) {
	var account *models.Account
	err := r.DB.
		Debug().
		Joins("Post", r.DB.Where(&models.Post{AccountID: id})).
		Joins("Comment", r.DB.Where(&models.Comment{AccountID: id})).
		First(&account, id).
		Error

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *queryResolver) GetCommentByAccountID(ctx context.Context, accountID int) ([]*models.Comment, error) {
	var comment []*models.Comment
	err := r.DB.
		Debug().
		Joins("Post", r.DB.Where(&models.Post{AccountID: accountID})).
		Joins("Account", r.DB.Where(&models.Account{ID: accountID})).
		Find(&comment, accountID).
		Error

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *queryResolver) GetFriendsByAccountID(ctx context.Context, accountID int) ([]*models.Friend, error) {
	var friends []*models.Friend
	err := r.DB.
		Debug().
		Joins("Friend", r.DB.Where(&models.Account{Type: "active"})).
		Where(&models.Friend{AccountID: accountID}).
		Find(&friends).
		Error

	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (r *queryResolver) GetLikesByAccountID(ctx context.Context, accountID int) ([]*models.Like, error) {
	var likes []*models.Like
	err := r.DB.
		Debug().
		Joins("Post").
		Joins("Account", r.DB.Where(&models.Account{ID: accountID})).
		Where(&models.Like{AccountID: accountID}).
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

func (r *queryResolver) GetMutesByAccountID(ctx context.Context, accountID int) ([]*models.Mute, error) {
	var mutes []*models.Mute
	err := r.DB.
		Debug().
		Joins("Mute", r.DB.Where(&models.Account{Type: "active"})).
		Where(&models.Mute{AccountID: accountID}).
		Find(&mutes).
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
		Joins("Account", r.DB.Where(&models.Account{Type: "active"})).
		Joins("Marker").
		Preload("Comment").
		Preload("Like").
		Find(&posts).
		Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *queryResolver) GetPostsByAccountID(ctx context.Context, accountID int) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.DB.
		Debug().
		Joins("Marker").
		Preload("Comment").
		Preload("Like").
		Where(&models.Post{AccountID: accountID}).
		Find(&posts).
		Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *queryResolver) GetRequestsByAccountID(ctx context.Context, accountID int) ([]*models.Request, error) {
	var requests []*models.Request
	err := r.DB.
		Debug().
		Joins("TargetAccount", r.DB.Where(&models.Account{Type: "active"})).
		Where(&models.Request{AccountID: accountID}).
		Find(&requests).
		Error

	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *queryResolver) GetRequestsByTargetID(ctx context.Context, targetAccountID int) ([]*models.Request, error) {
	var requests []*models.Request
	err := r.DB.
		Debug().
		Joins("Account", r.DB.Where(&models.Account{Type: "active"})).
		Where(&models.Request{TargetAccountID: targetAccountID}).
		Find(&requests).
		Error

	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *queryResolver) GetSessionByAccountID(ctx context.Context, accountID int) (*models.Session, error) {
	var session *models.Session
	err := r.DB.
		Where(&models.Session{AccountID: accountID}).
		First(&session).
		Error

	if err != nil {
		return nil, err
	}

	return session, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
