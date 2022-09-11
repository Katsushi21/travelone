package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *sessionResolver) AccountID(ctx context.Context, obj *ent.Session) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Session returns travelone.SessionResolver implementation.
func (r *Resolver) Session() travelone.SessionResolver { return &sessionResolver{r} }

type sessionResolver struct{ *Resolver }
