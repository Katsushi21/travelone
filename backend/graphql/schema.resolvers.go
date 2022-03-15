package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/models"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input models.PostInput) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMarker(ctx context.Context, input models.MarkerInput) (*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input models.CommentInput) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRequest(ctx context.Context, input models.RequestInput) (*models.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input models.PostInput) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateLiked(ctx context.Context, id string, input models.LikedInput) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UserInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSession(ctx context.Context, id string, input models.SessionInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateFriend(ctx context.Context, id string, input models.FriendInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMute(ctx context.Context, id string, input *models.MuteInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id string, input models.MarkerInput) (*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input models.CommentInput) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UploadFile(ctx context.Context, input models.UploadFile) (*models.File, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, id string, input models.RequestInput) (*models.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMarker(ctx context.Context, id string) (*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, id string) (*models.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Post(ctx context.Context) ([]*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comment(ctx context.Context) ([]*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Marker(ctx context.Context) ([]*models.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Request(ctx context.Context) ([]*models.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
