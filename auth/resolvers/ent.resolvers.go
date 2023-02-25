package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
	"github.com/google/uuid"
)

func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Accounts(ctx context.Context) ([]*ent.Account, error) {
	return r.client.Account.Query().All(ctx)
}

func (r *queryResolver) Comments(ctx context.Context) ([]*ent.Comment, error) {
	return r.client.Comment.Query().All(ctx)
}

func (r *queryResolver) Friends(ctx context.Context) ([]*ent.Friend, error) {
	return r.client.Friend.Query().All(ctx)
}

func (r *queryResolver) Likes(ctx context.Context) ([]*ent.Like, error) {
	return r.client.Like.Query().All(ctx)
}

func (r *queryResolver) Markers(ctx context.Context) ([]*ent.Marker, error) {
	return r.client.Marker.Query().All(ctx)
}

func (r *queryResolver) Mutes(ctx context.Context) ([]*ent.Mute, error) {
	return r.client.Mute.Query().All(ctx)
}

func (r *queryResolver) Posts(ctx context.Context) ([]*ent.Post, error) {
	return r.client.Post.Query().All(ctx)
}

func (r *queryResolver) Requests(ctx context.Context) ([]*ent.Request, error) {
	return r.client.Request.Query().All(ctx)
}

func (r *queryResolver) Sessions(ctx context.Context) ([]*ent.Session, error) {
	return r.client.Session.Query().All(ctx)
}

// Query returns travelone.QueryResolver implementation.
func (r *Resolver) Query() travelone.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
