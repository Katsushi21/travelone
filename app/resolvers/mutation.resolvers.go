package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
	"github.com/Katsushi21/travelone/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.CreateAccountInput) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, id uuid.UUID, input model.UpdateAccountInput) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LoginQuery(ctx context.Context, email string, password string) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.CreateCommentInput) (*ent.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id uuid.UUID, input model.UpdateCommentInput) (*ent.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id uuid.UUID) (*ent.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateFriend(ctx context.Context, input model.CreateFriendInput) (*ent.Friend, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteFriend(ctx context.Context, id uuid.UUID) (*ent.Friend, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateLike(ctx context.Context, input model.CreateLikeInput) (*ent.Like, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteLike(ctx context.Context, id uuid.UUID) (*ent.Like, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMarker(ctx context.Context, input model.CreateMarkerInput) (*ent.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id uuid.UUID, input model.UpdateMarkerInput) (*ent.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMarker(ctx context.Context, id uuid.UUID) (*ent.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMute(ctx context.Context, input model.CreateMuteInput) (*ent.Mute, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMute(ctx context.Context, id uuid.UUID) (*ent.Mute, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (*ent.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id uuid.UUID, input model.UpdatePostInput) (*ent.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePost(ctx context.Context, id uuid.UUID) (*ent.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRequest(ctx context.Context, input model.CreateRequestInput) (*ent.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, id uuid.UUID, input model.UpdateRequestInput) (*ent.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, id uuid.UUID) (*ent.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSession(ctx context.Context, input model.CreateSessionInput) (*ent.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSession(ctx context.Context, id uuid.UUID, input model.UpdateSessionInput) (*ent.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSession(ctx context.Context, id uuid.UUID) (*ent.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns travelone.MutationResolver implementation.
func (r *Resolver) Mutation() travelone.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
