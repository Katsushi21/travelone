package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *friendResolver) AccountID(ctx context.Context, obj *ent.Friend) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendResolver) FriendID(ctx context.Context, obj *ent.Friend) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Friend returns travelone.FriendResolver implementation.
func (r *Resolver) Friend() travelone.FriendResolver { return &friendResolver{r} }

type friendResolver struct{ *Resolver }
