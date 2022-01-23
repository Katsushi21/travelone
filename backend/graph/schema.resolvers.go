package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi/traveling_alone/graph/generated"
	"github.com/Katsushi/traveling_alone/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMarker(ctx context.Context, input model.MarkerInput) (*model.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.CommentInput) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input *model.PostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateLiked(ctx context.Context, id string, input *model.LikedInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id string, input *model.MarkerInput) (*model.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input *model.CommentInput) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPosts(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUsers(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProfiles(ctx context.Context, id string) (*model.Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetComments(ctx context.Context, id string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
