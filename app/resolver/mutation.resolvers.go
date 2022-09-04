package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input travelone.AccountInput) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented: CreateAccount - createAccount"))
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, id string, input travelone.AccountInput) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented: UpdateAccount - updateAccount"))
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, id string) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented: DeleteAccount - deleteAccount"))
}

func (r *mutationResolver) LoginQuery(ctx context.Context, input travelone.LoginInput) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented: LoginQuery - loginQuery"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input travelone.CommentInput) (*ent.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (*ent.Comment, error) {
	panic(fmt.Errorf("not implemented: DeleteComment - deleteComment"))
}

func (r *mutationResolver) UploadFile(ctx context.Context, input travelone.UploadFile) (*travelone.File, error) {
	panic(fmt.Errorf("not implemented: UploadFile - uploadFile"))
}

func (r *mutationResolver) CreateFriend(ctx context.Context, input travelone.FriendInput) (*ent.Friend, error) {
	panic(fmt.Errorf("not implemented: CreateFriend - createFriend"))
}

func (r *mutationResolver) DeleteFriend(ctx context.Context, input travelone.FriendInput) (*ent.Friend, error) {
	panic(fmt.Errorf("not implemented: DeleteFriend - deleteFriend"))
}

func (r *mutationResolver) CreateLike(ctx context.Context, input travelone.LikeInput) (*ent.Like, error) {
	panic(fmt.Errorf("not implemented: CreateLike - createLike"))
}

func (r *mutationResolver) DeleteLike(ctx context.Context, input travelone.LikeInput) (*ent.Like, error) {
	panic(fmt.Errorf("not implemented: DeleteLike - deleteLike"))
}

func (r *mutationResolver) CreateMarker(ctx context.Context, input travelone.MarkerInput) (*ent.Marker, error) {
	panic(fmt.Errorf("not implemented: CreateMarker - createMarker"))
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id string, input travelone.MarkerInput) (*ent.Marker, error) {
	panic(fmt.Errorf("not implemented: UpdateMarker - updateMarker"))
}

func (r *mutationResolver) DeleteMarker(ctx context.Context, id string) (*ent.Marker, error) {
	panic(fmt.Errorf("not implemented: DeleteMarker - deleteMarker"))
}

func (r *mutationResolver) CreateMute(ctx context.Context, input travelone.MuteInput) (*ent.Mute, error) {
	panic(fmt.Errorf("not implemented: CreateMute - createMute"))
}

func (r *mutationResolver) DeleteMute(ctx context.Context, input *travelone.MuteInput) (*ent.Mute, error) {
	panic(fmt.Errorf("not implemented: DeleteMute - deleteMute"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input travelone.PostInput) (*ent.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - createPost"))
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input travelone.PostInput) (*ent.Post, error) {
	panic(fmt.Errorf("not implemented: UpdatePost - updatePost"))
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*ent.Post, error) {
	panic(fmt.Errorf("not implemented: DeletePost - deletePost"))
}

func (r *mutationResolver) CreateRequest(ctx context.Context, input travelone.RequestInput) (*ent.Request, error) {
	panic(fmt.Errorf("not implemented: CreateRequest - createRequest"))
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, input travelone.RequestInput) (*ent.Request, error) {
	panic(fmt.Errorf("not implemented: UpdateRequest - updateRequest"))
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, input travelone.RequestInput) (*ent.Request, error) {
	panic(fmt.Errorf("not implemented: DeleteRequest - deleteRequest"))
}

func (r *mutationResolver) CreateSession(ctx context.Context, input travelone.SessionInput) (*ent.Session, error) {
	panic(fmt.Errorf("not implemented: CreateSession - createSession"))
}

func (r *mutationResolver) UpdateSession(ctx context.Context, input travelone.SessionInput) (*ent.Session, error) {
	panic(fmt.Errorf("not implemented: UpdateSession - updateSession"))
}

func (r *mutationResolver) DeleteSession(ctx context.Context, input travelone.SessionInput) (*ent.Session, error) {
	panic(fmt.Errorf("not implemented: DeleteSession - deleteSession"))
}

// Mutation returns travelone.MutationResolver implementation.
func (r *Resolver) Mutation() travelone.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
