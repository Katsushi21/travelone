package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *commentResolver) ID(ctx context.Context, obj *ent.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) PostID(ctx context.Context, obj *ent.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) AccountID(ctx context.Context, obj *ent.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comment returns travelone.CommentResolver implementation.
func (r *Resolver) Comment() travelone.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
