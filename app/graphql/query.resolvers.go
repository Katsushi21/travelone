package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *queryResolver) AccountPageInfo(ctx context.Context, id int) (*models.Account, error) {
	var account *models.Account
	err := r.DB.
		Debug().
		Preload("Post").
		Preload("Friend.Friend").
		First(&account, id).
		Error

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *queryResolver) MyPageInfo(ctx context.Context, id int) (*models.Account, error) {
	var account *models.Account
	err := r.DB.
		Debug().
		Preload("Post").
		Preload("Like.Post").
		Preload("Like.Account").
		Preload("Friend.Friend").
		Preload("Mute.Mute").
		First(&account, id).
		Error

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *queryResolver) LikesByPost(ctx context.Context, postID int) ([]*models.Like, error) {
	var likes []*models.Like
	err := r.DB.
		Debug().
		Joins("Account").
		Find(&likes).
		Error

	if err != nil {
		return nil, err
	}

	return likes, nil
}

func (r *queryResolver) AllMarkers(ctx context.Context) ([]*models.Marker, error) {
	var markers []*models.Marker
	err := r.DB.
		Debug().
		Joins("Post").
		Preload("Post.Account").
		Preload("Post.Like.Account").
		Find(&markers).
		Error

	if err != nil {
		return nil, err
	}

	return markers, nil
}

func (r *queryResolver) AllPosts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.DB.
		Debug().
		Joins("Account").
		Joins("Marker").
		Preload("Comment.Account").
		Preload("Like.Account").
		Find(&posts).
		Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *queryResolver) RequestsByAccountID(ctx context.Context, accountID int) ([]*models.Request, error) {
	var requests []*models.Request
	err := r.DB.
		Debug().
		Joins("TargetAccount").
		Where(&models.Request{AccountID: accountID}).
		Find(&requests).
		Error

	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *queryResolver) RequestsByTargetID(ctx context.Context, targetAccountID int) ([]*models.Request, error) {
	var requests []*models.Request
	err := r.DB.
		Debug().
		Joins("Account").
		Where(&models.Request{TargetAccountID: targetAccountID}).
		Find(&requests).
		Error

	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *queryResolver) SessionByAccountID(ctx context.Context, accountID int) (*models.Session, error) {
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
