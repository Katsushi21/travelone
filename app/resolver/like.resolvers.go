package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *likeResolver) PostID(ctx context.Context, obj *ent.Like) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *likeResolver) AccountID(ctx context.Context, obj *ent.Like) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Like returns travelone.LikeResolver implementation.
func (r *Resolver) Like() travelone.LikeResolver { return &likeResolver{r} }

type likeResolver struct{ *Resolver }
