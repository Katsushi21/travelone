package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
	"github.com/Katsushi21/travelone/ent/account"
	"github.com/Katsushi21/travelone/ent/like"
	"github.com/Katsushi21/travelone/ent/request"
	"github.com/Katsushi21/travelone/model"
	"github.com/google/uuid"
)

func (r *queryResolver) QueryAccountByID(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	return r.client.Account.Get(ctx, id)
}

func (r *queryResolver) QueryAccountByMyID(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	return r.client.Account.Get(ctx, id)
}

func (r *queryResolver) QueryAccountByLoginInput(ctx context.Context, input model.LoginInput) (*ent.Account, error) {
	return r.client.Account.
		Query().
		Where(
			account.Email(input.Email),
			account.Password(input.Password),
		).
		Only(ctx)
}

func (r *queryResolver) QueryLikesByPostID(ctx context.Context, postID uuid.UUID) ([]*ent.Like, error) {
	return r.client.Like.
		Query().
		Where(like.PostID(postID)).
		All(ctx)
}

func (r *queryResolver) QueryMarkers(ctx context.Context) ([]*ent.Marker, error) {
	return r.client.Marker.
		Query().
		All(ctx)
}

func (r *queryResolver) QueryPosts(ctx context.Context) ([]*ent.Post, error) {
	return r.client.Post.
		Query().
		All(ctx)
}

func (r *queryResolver) QueryRequestsByAccountID(ctx context.Context, accountID uuid.UUID) ([]*ent.Request, error) {
	return r.client.Request.
		Query().
		Where(request.AccountID(accountID)).
		All(ctx)
}

func (r *queryResolver) QueryRequestsByRequestID(ctx context.Context, requestID uuid.UUID) ([]*ent.Request, error) {
	return r.client.Request.
		Query().
		Where(request.RequestID(requestID)).
		All(ctx)
}

func (r *queryResolver) QuerySessionByAccountID(ctx context.Context, id uuid.UUID) (*ent.Session, error) {
	return r.client.Session.Get(ctx, id)
}

// Query returns travelone.QueryResolver implementation.
func (r *Resolver) Query() travelone.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
