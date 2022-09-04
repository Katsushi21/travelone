package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *requestResolver) AccountID(ctx context.Context, obj *ent.Request) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestResolver) TargetAccountID(ctx context.Context, obj *ent.Request) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestResolver) Status(ctx context.Context, obj *ent.Request) (travelone.RequestStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *requestResolver) TargetAccount(ctx context.Context, obj *ent.Request) (*ent.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

// Request returns travelone.RequestResolver implementation.
func (r *Resolver) Request() travelone.RequestResolver { return &requestResolver{r} }

type requestResolver struct{ *Resolver }
