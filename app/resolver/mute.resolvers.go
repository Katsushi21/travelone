package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

func (r *muteResolver) AccountID(ctx context.Context, obj *ent.Mute) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *muteResolver) MuteID(ctx context.Context, obj *ent.Mute) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mute returns travelone.MuteResolver implementation.
func (r *Resolver) Mute() travelone.MuteResolver { return &muteResolver{r} }

type muteResolver struct{ *Resolver }
