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

func (r *queryResolver) Accounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AccountOrder, where *ent.AccountWhereInput) (*ent.AccountConnection, error) {
	return r.client.Account.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithAccountOrder(orderBy),
		)
}

func (r *queryResolver) Comments(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CommentOrder, where *ent.CommentWhereInput) (*ent.CommentConnection, error) {
	return r.client.Comment.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithCommentOrder(orderBy),
		)
}

func (r *queryResolver) Friends(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.FriendOrder, where *ent.FriendWhereInput) (*ent.FriendConnection, error) {
	return r.client.Friend.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithFriendOrder(orderBy),
		)
}

func (r *queryResolver) Likes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.LikeOrder, where *ent.LikeWhereInput) (*ent.LikeConnection, error) {
	return r.client.Like.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithLikeOrder(orderBy),
		)
}

func (r *queryResolver) Markers(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MarkerOrder, where *ent.MarkerWhereInput) (*ent.MarkerConnection, error) {
	return r.client.Marker.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithMarkerOrder(orderBy),
		)
}

func (r *queryResolver) Mutes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MuteOrder, where *ent.MuteWhereInput) (*ent.MuteConnection, error) {
	return r.client.Mute.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithMuteOrder(orderBy),
		)
}

func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	return r.client.Post.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithPostOrder(orderBy),
		)
}

func (r *queryResolver) Requests(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.RequestOrder, where *ent.RequestWhereInput) (*ent.RequestConnection, error) {
	return r.client.Request.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithRequestOrder(orderBy),
		)
}

func (r *queryResolver) Sessions(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SessionOrder, where *ent.SessionWhereInput) (*ent.SessionConnection, error) {
	return r.client.Session.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithSessionOrder(orderBy),
		)
}

// Query returns travelone.QueryResolver implementation.
func (r *Resolver) Query() travelone.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
