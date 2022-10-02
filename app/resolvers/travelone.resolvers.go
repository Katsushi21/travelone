package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

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

func (r *mutationResolver) LoginQuery(ctx context.Context, email string, password string) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateLike(ctx context.Context, input ent.CreateLikeInput) (*ent.Like, error) {
	return r.client.Like.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteLike(ctx context.Context, id uuid.UUID) (*ent.Like, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMarker(ctx context.Context, input ent.CreateMarkerInput) (*ent.Marker, error) {
	return r.client.Marker.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id uuid.UUID, input ent.UpdateMarkerInput) (*ent.Marker, error) {
	return r.client.Marker.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteMarker(ctx context.Context, id uuid.UUID) (*ent.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMute(ctx context.Context, input ent.CreateMuteInput) (*ent.Mute, error) {
	return r.client.Mute.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteMute(ctx context.Context, id uuid.UUID) (*ent.Mute, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput) (*ent.Post, error) {
	return r.client.Post.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id uuid.UUID, input ent.UpdatePostInput) (*ent.Post, error) {
	return r.client.Post.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeletePost(ctx context.Context, id uuid.UUID) (*ent.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRequest(ctx context.Context, input ent.CreateRequestInput) (*ent.Request, error) {
	return r.client.Request.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, id uuid.UUID, input ent.UpdateRequestInput) (*ent.Request, error) {
	return r.client.Request.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, id uuid.UUID) (*ent.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSession(ctx context.Context, input ent.CreateSessionInput) (*ent.Session, error) {
	return r.client.Session.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateSession(ctx context.Context, id uuid.UUID, input ent.UpdateSessionInput) (*ent.Session, error) {
	return r.client.Session.UpdateOneID(*input.AccountID).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteSession(ctx context.Context, id uuid.UUID) (*ent.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns travelone.MutationResolver implementation.
func (r *Resolver) Mutation() travelone.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
