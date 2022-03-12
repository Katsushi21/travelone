package graph

import "github.com/Katsushi21/traveling_alone/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	posts    []*model.Post
	users    []*model.User
	markers  []*model.Marker
	comments []*model.Comment
	Requests []*model.Request
}
