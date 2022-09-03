package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
)

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input travelone.AccountInput) (*travelone.Account, error) {
	panic(fmt.Errorf("not implemented: CreateAccount - createAccount"))
}

// UpdateAccount is the resolver for the updateAccount field.
func (r *mutationResolver) UpdateAccount(ctx context.Context, id string, input travelone.AccountInput) (*travelone.Account, error) {
	panic(fmt.Errorf("not implemented: UpdateAccount - updateAccount"))
}

// DeleteAccount is the resolver for the deleteAccount field.
func (r *mutationResolver) DeleteAccount(ctx context.Context, id string) (*travelone.Account, error) {
	panic(fmt.Errorf("not implemented: DeleteAccount - deleteAccount"))
}

// LoginQuery is the resolver for the loginQuery field.
func (r *mutationResolver) LoginQuery(ctx context.Context, input travelone.LoginInput) (*travelone.Account, error) {
	panic(fmt.Errorf("not implemented: LoginQuery - loginQuery"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input travelone.CommentInput) (*travelone.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (*travelone.Comment, error) {
	panic(fmt.Errorf("not implemented: DeleteComment - deleteComment"))
}

// UploadFile is the resolver for the uploadFile field.
func (r *mutationResolver) UploadFile(ctx context.Context, input travelone.UploadFile) (*travelone.File, error) {
	panic(fmt.Errorf("not implemented: UploadFile - uploadFile"))
}

// CreateFriend is the resolver for the createFriend field.
func (r *mutationResolver) CreateFriend(ctx context.Context, input travelone.FriendInput) (*travelone.Friend, error) {
	panic(fmt.Errorf("not implemented: CreateFriend - createFriend"))
}

// DeleteFriend is the resolver for the deleteFriend field.
func (r *mutationResolver) DeleteFriend(ctx context.Context, input travelone.FriendInput) (*travelone.Friend, error) {
	panic(fmt.Errorf("not implemented: DeleteFriend - deleteFriend"))
}

// CreateLike is the resolver for the createLike field.
func (r *mutationResolver) CreateLike(ctx context.Context, input travelone.LikeInput) (*travelone.Like, error) {
	panic(fmt.Errorf("not implemented: CreateLike - createLike"))
}

// DeleteLike is the resolver for the deleteLike field.
func (r *mutationResolver) DeleteLike(ctx context.Context, input travelone.LikeInput) (*travelone.Like, error) {
	panic(fmt.Errorf("not implemented: DeleteLike - deleteLike"))
}

// CreateMarker is the resolver for the createMarker field.
func (r *mutationResolver) CreateMarker(ctx context.Context, input travelone.MarkerInput) (*travelone.Marker, error) {
	panic(fmt.Errorf("not implemented: CreateMarker - createMarker"))
}

// UpdateMarker is the resolver for the updateMarker field.
func (r *mutationResolver) UpdateMarker(ctx context.Context, id string, input travelone.MarkerInput) (*travelone.Marker, error) {
	panic(fmt.Errorf("not implemented: UpdateMarker - updateMarker"))
}

// DeleteMarker is the resolver for the deleteMarker field.
func (r *mutationResolver) DeleteMarker(ctx context.Context, id string) (*travelone.Marker, error) {
	panic(fmt.Errorf("not implemented: DeleteMarker - deleteMarker"))
}

// CreateMute is the resolver for the createMute field.
func (r *mutationResolver) CreateMute(ctx context.Context, input travelone.MuteInput) (*travelone.Mute, error) {
	panic(fmt.Errorf("not implemented: CreateMute - createMute"))
}

// DeleteMute is the resolver for the deleteMute field.
func (r *mutationResolver) DeleteMute(ctx context.Context, input *travelone.MuteInput) (*travelone.Mute, error) {
	panic(fmt.Errorf("not implemented: DeleteMute - deleteMute"))
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input travelone.PostInput) (*travelone.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - createPost"))
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input travelone.PostInput) (*travelone.Post, error) {
	panic(fmt.Errorf("not implemented: UpdatePost - updatePost"))
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*travelone.Post, error) {
	panic(fmt.Errorf("not implemented: DeletePost - deletePost"))
}

// CreateRequest is the resolver for the createRequest field.
func (r *mutationResolver) CreateRequest(ctx context.Context, input travelone.RequestInput) (*travelone.Request, error) {
	panic(fmt.Errorf("not implemented: CreateRequest - createRequest"))
}

// UpdateRequest is the resolver for the updateRequest field.
func (r *mutationResolver) UpdateRequest(ctx context.Context, input travelone.RequestInput) (*travelone.Request, error) {
	panic(fmt.Errorf("not implemented: UpdateRequest - updateRequest"))
}

// DeleteRequest is the resolver for the deleteRequest field.
func (r *mutationResolver) DeleteRequest(ctx context.Context, input travelone.RequestInput) (*travelone.Request, error) {
	panic(fmt.Errorf("not implemented: DeleteRequest - deleteRequest"))
}

// CreateSession is the resolver for the createSession field.
func (r *mutationResolver) CreateSession(ctx context.Context, input travelone.SessionInput) (*travelone.Session, error) {
	panic(fmt.Errorf("not implemented: CreateSession - createSession"))
}

// UpdateSession is the resolver for the updateSession field.
func (r *mutationResolver) UpdateSession(ctx context.Context, input travelone.SessionInput) (*travelone.Session, error) {
	panic(fmt.Errorf("not implemented: UpdateSession - updateSession"))
}

// DeleteSession is the resolver for the deleteSession field.
func (r *mutationResolver) DeleteSession(ctx context.Context, input travelone.SessionInput) (*travelone.Session, error) {
	panic(fmt.Errorf("not implemented: DeleteSession - deleteSession"))
}

// Mutation returns travelone.MutationResolver implementation.
func (r *Resolver) Mutation() travelone.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
