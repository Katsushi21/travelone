package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/traveling_alone/graph/generated"
	"github.com/Katsushi21/traveling_alone/graph/model"
)

func (r *commentResolver) User(ctx context.Context, obj *model.Comment) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostInput) (*model.Post, error) {
	post := &model.Post{
		UID:   input.UID,
		Title: input.Title,
		Body:  input.Body,
		Img:   &input.Img,
	}
	r.posts = append(r.posts, post)
	return post, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	user := &model.User{
		Email:        &input.Email,
		Password:     &input.Password,
		Type:         &input.Type,
		Name:         &input.Name,
		Age:          &input.Age,
		Gender:       &input.Gender,
		Avatar:       &input.Avatar,
		Introduction: &input.Introduction,
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *mutationResolver) CreateMarker(ctx context.Context, input model.MarkerInput) (*model.Marker, error) {
	marker := &model.Marker{
		PostID: input.PostID,
		Title:  &input.Title,
		Lat:    &input.Lat,
		Lng:    &input.Lng,
	}
	r.markers = append(r.markers, marker)
	return marker, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.CommentInput) (*model.Comment, error) {
	comment := &model.Comment{
		UID:    input.UID,
		PostID: input.PostID,
		Body:   &input.Body,
	}
	r.comments = append(r.comments, comment)
	return comment, nil
}

func (r *mutationResolver) CreateRequest(ctx context.Context, input model.RequestInput) (*model.Request, error) {
	request := &model.Request{
		Request:   *input.Request,
		Requested: *input.Requested,
		Status:    input.Status,
	}
	r.requests = append(r.requests, request)
	return request, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input model.PostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateLiked(ctx context.Context, id string, input model.LikedInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSession(ctx context.Context, id string, input model.SessionInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateFriend(ctx context.Context, id string, input model.FriendInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMute(ctx context.Context, id string, input *model.MuteInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMarker(ctx context.Context, id string, input model.MarkerInput) (*model.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input model.CommentInput) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UploadFile(ctx context.Context, input model.UploadFile) (*model.File, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, id string, input model.RequestInput) (*model.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMarker(ctx context.Context, id string) (*model.Marker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRequest(ctx context.Context, id string) (*model.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	return &model.User{ID: obj.UID}, nil
}

func (r *queryResolver) Post(ctx context.Context) ([]*model.Post, error) {
	return r.posts, nil
}

func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *queryResolver) Comment(ctx context.Context) ([]*model.Comment, error) {
	return r.comments, nil
}

func (r *queryResolver) Marker(ctx context.Context) ([]*model.Marker, error) {
	return r.markers, nil
}

func (r *queryResolver) Request(ctx context.Context) ([]*model.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestResolver) Request(ctx context.Context, obj *model.Request) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestResolver) Requested(ctx context.Context, obj *model.Request) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Request returns generated.RequestResolver implementation.
func (r *Resolver) Request() generated.RequestResolver { return &requestResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type requestResolver struct{ *Resolver }
