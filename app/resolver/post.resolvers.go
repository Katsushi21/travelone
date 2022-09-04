package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *postResolver) ID(ctx context.Context, obj *ent.Post) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) AccountID(ctx context.Context, obj *ent.Post) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) Like(ctx context.Context, obj *ent.Post) ([]*ent.Like, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) Comment(ctx context.Context, obj *ent.Post) ([]*ent.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Post returns travelone.PostResolver implementation.
func (r *Resolver) Post() travelone.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
