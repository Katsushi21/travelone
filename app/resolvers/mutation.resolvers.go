package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input ent.CreateAccountInput) (*ent.Account, error) {
	return r.client.Account.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, id uuid.UUID, input ent.UpdateAccountInput) (*ent.Account, error) {
	return r.client.Account.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	account, err := r.client.Account.Get(ctx, id)
	if err != nil {
		return account, err
	}

	err = r.client.Account.DeleteOneID(id).Exec(ctx)
	return account, err
}

func (r *mutationResolver) CreateComment(ctx context.Context, input ent.CreateCommentInput) (*ent.Comment, error) {
	return r.client.Comment.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id uuid.UUID, input ent.UpdateCommentInput) (*ent.Comment, error) {
	return r.client.Comment.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id uuid.UUID) (*ent.Comment, error) {
	comment, err := r.client.Comment.Get(ctx, id)
	if err != nil {
		return comment, err
	}

	err = r.client.Comment.DeleteOneID(id).Exec(ctx)
	return comment, err
}

func (r *mutationResolver) CreateFriend(ctx context.Context, input ent.CreateFriendInput) (*ent.Friend, error) {
	return r.client.Friend.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteFriend(ctx context.Context, id uuid.UUID) (*ent.Friend, error) {
	friend, err := r.client.Friend.Get(ctx, id)
	if err != nil {
		return friend, err
	}

	err = r.client.Friend.DeleteOneID(id).Exec(ctx)
	return friend, err
}

func (r *mutationResolver) CreateLike(ctx context.Context, input ent.CreateLikeInput) (*ent.Like, error) {
	return r.client.Like.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteLike(ctx context.Context, id uuid.UUID) (*ent.Like, error) {
	like, err := r.client.Like.Get(ctx, id)
	if err != nil {
		return like, err
	}

	err = r.client.Like.DeleteOneID(id).Exec(ctx)
	return like, err
}

func (r *mutationResolver) CreateMarker(ctx context.Context, input ent.CreateMarkerInput) (*ent.Marker, error) {
	return r.client.Marker.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id uuid.UUID, input ent.UpdateMarkerInput) (*ent.Marker, error) {
	return r.client.Marker.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteMarker(ctx context.Context, id uuid.UUID) (*ent.Marker, error) {
	marker, err := r.client.Marker.Get(ctx, id)
	if err != nil {
		return marker, err
	}

	err = r.client.Marker.DeleteOneID(id).Exec(ctx)
	return marker, err
}

func (r *mutationResolver) CreateMute(ctx context.Context, input ent.CreateMuteInput) (*ent.Mute, error) {
	return r.client.Mute.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteMute(ctx context.Context, id uuid.UUID) (*ent.Mute, error) {
	mute, err := r.client.Mute.Get(ctx, id)
	if err != nil {
		return mute, err
	}

	err = r.client.Mute.DeleteOneID(id).Exec(ctx)
	return mute, err
}

func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput) (*ent.Post, error) {
	return r.client.Post.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id uuid.UUID, input ent.UpdatePostInput) (*ent.Post, error) {
	return r.client.Post.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeletePost(ctx context.Context, id uuid.UUID) (*ent.Post, error) {
	post, err := r.client.Post.Get(ctx, id)
	if err != nil {
		return post, err
	}

	err = r.client.Post.DeleteOneID(id).Exec(ctx)
	return post, err
}

func (r *mutationResolver) CreateRequest(ctx context.Context, input ent.CreateRequestInput) (*ent.Request, error) {
	return r.client.Request.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, id uuid.UUID, input ent.UpdateRequestInput) (*ent.Request, error) {
	return r.client.Request.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, id uuid.UUID) (*ent.Request, error) {
	request, err := r.client.Request.Get(ctx, id)
	if err != nil {
		return request, err
	}

	err = r.client.Request.DeleteOneID(id).Exec(ctx)
	return request, err
}

func (r *mutationResolver) CreateSession(ctx context.Context, input ent.CreateSessionInput) (*ent.Session, error) {
	return r.client.Session.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateSession(ctx context.Context, id uuid.UUID, input ent.UpdateSessionInput) (*ent.Session, error) {
	return r.client.Session.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteSession(ctx context.Context, id uuid.UUID) (*ent.Session, error) {
	session, err := r.client.Session.Get(ctx, id)
	if err != nil {
		return session, err
	}

	err = r.client.Session.DeleteOneID(id).Exec(ctx)
	return session, err
}

// Mutation returns travelone.MutationResolver implementation.
func (r *Resolver) Mutation() travelone.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
