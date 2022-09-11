package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *markerResolver) ID(ctx context.Context, obj *ent.Marker) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *markerResolver) PostID(ctx context.Context, obj *ent.Marker) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Marker returns travelone.MarkerResolver implementation.
func (r *Resolver) Marker() travelone.MarkerResolver { return &markerResolver{r} }

type markerResolver struct{ *Resolver }
