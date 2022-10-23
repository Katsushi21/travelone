package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Katsushi21/travelone/ent"
	"github.com/Katsushi21/travelone/ent/like"
	"github.com/Katsushi21/travelone/ent/request"
	"github.com/google/uuid"
)

func (r *queryResolver) AccountByID(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	return r.client.Account.Get(ctx, id)
}

func (r *queryResolver) AccountBySelfID(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	return r.client.Account.Get(ctx, id)
}

func (r *queryResolver) LikesByPostID(ctx context.Context, postID uuid.UUID) ([]*ent.Like, error) {
	like, err := r.client.Like.
		Query().
		Where(
			like.PostID(postID),
		).
		All(ctx)

	return like, err
}

func (r *queryResolver) RequestsByAccountID(ctx context.Context, accountID uuid.UUID) ([]*ent.Request, error) {
	request, err := r.client.Request.
		Query().
		Where(
			request.AccountID(accountID),
		).
		All(ctx)

	return request, err
}

func (r *queryResolver) RequestsByRequestID(ctx context.Context, requestID uuid.UUID) ([]*ent.Request, error) {
	request, err := r.client.Request.
		Query().
		Where(
			request.AccountID(requestID),
		).
		All(ctx)

	return request, err
}

func (r *queryResolver) SessionByID(ctx context.Context, id uuid.UUID) (*ent.Session, error) {
	return r.client.Session.Get(ctx, id)
}
